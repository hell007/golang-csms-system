package apis

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/utils/tool"
	"image/png"
	"strconv"
	"strings"
	"time"

	redisClient "go-mvc/framework/cache/redis"
	"go-mvc/framework/conf"
	models "go-mvc/framework/models/member"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/captcha"
	"go-mvc/framework/utils/encrypt"
	"go-mvc/framework/utils/ips"
	"go-mvc/framework/utils/response"
)

//24
var key = conf.GlobalConfig.JWTSecret + time.Now().Format(conf.GlobalConfig.Timeformat)

//5分钟
var captchatime = time.Duration(conf.GlobalConfig.CaptchaExprire) * time.Second

//2小时
var locktime = 2 * 60 * 60 * time.Second

//1小时
var exprire = time.Duration(conf.GlobalConfig.RedisExprire) * time.Second

/**
会员注册
*/
func Register(ctx iris.Context) {
	var (
		err    error
		has    bool
		effect int64
		user   = new(models.Member)
		member = new(models.Member)
	)

	// 读取
	if err = ctx.ReadJSON(&member); err != nil {
		ctx.Application().Logger().Errorf("user.Register：手机号为[%s]的用户，注册失败，原因：[%s]", member.Mobile, err)
		response.Error(ctx, iris.StatusBadRequest, response.RegisteFailur, nil)
		return
	}

	// 是否已存在
	user.Mobile = member.Mobile
	has, err = services.NewMemberService().Exist(user)
	if has || err != nil {
		ctx.Application().Logger().Errorf("user.Register：手机号为[%s]的用户，已存在：[%s]", user.Mobile, err)
		response.Failur(ctx, response.RegisteExist, nil)
		return
	}

	// 获取用户ip
	ip := ips.ClientIP(ctx.Request())
	member.Ip = ip

	// 数据处理
	member.Password = encrypt.AESEncrypt(member.Password, conf.GlobalConfig.JWTSalt)
	member.CreateTime = time.Now()
	member.Status = 1

	effect, err = services.NewMemberService().Create(member)
	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("user.Register：手机号为[%s]的用户，注册失败：[%s]", member.Mobile, err)
		response.Failur(ctx, response.RegisteFailur, nil)
		return
	}

	response.Ok_(ctx, response.RegisteSuccess)
	return
}

/**
验证码
*/
func Captcha(ctx iris.Context) {
	cp := captcha.NewCaptcha(120, 40, 4)
	cp.SetMode(0)
	code, img := cp.OutPut()

	// 保存code 60s
	err := redisClient.Set(conf.GlobalConfig.CaptchaKey, code, captchatime).Err()
	if err != nil {
		ctx.Application().Logger().Errorf("User.Captcha验证码保存错误：[%s]", err)
	}
	ctx.Header("Content-Type", "image/png; charset=utf-8")
	_ = png.Encode(ctx, img)
}

/**
会员登录
postman post 设置
raw: {"mobile": "","password": ""}
*/
func Login(ctx iris.Context) {
	var (
		err        error
		has, equal bool
		code       string
		columns    []string
		user       = new(models.LoginUser)
		member     = new(models.Member)
	)

	// 参数
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("User.Login参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.LoginFailur, nil)
		return
	}

	// 如果开启验证码校验，进行验证码校对
	if conf.GlobalConfig.CaptchaVerify {
		code, err = redisClient.Get(conf.GlobalConfig.CaptchaKey).Result()
		if code != strings.ToUpper(user.Code) || err != nil {
			ctx.Application().Logger().Errorf("User.Login验证码对比错误：[%s]", err)
			response.Failur(ctx, response.CaptchaFailur, nil)
			return
		}
	}

	// 查询用户
	member, has, err = services.NewMemberService().GetMemberByMobile(user.Mobile)
	if err != nil {
		ctx.Application().Logger().Errorf("User.Login查询用户错误：[%s]", err)
		response.Failur(ctx, response.LoginFailur, nil)
		return
	}

	// 不存在，或者手机号不正确
	if !has {
		ctx.Application().Logger().Error("User.Login错误：用户不存在，或者登录的手机号不正确")
		response.Failur(ctx, response.LoginFailur, nil)
		return
	}

	// 存在，验证密码
	equal = encrypt.CheckPWD(user.Password, member.Password, conf.GlobalConfig.JWTSalt)
	if !equal {
		ctx.Application().Logger().Error("User.Login错误：用户存在，登录的密码不正确")
		response.Failur(ctx, response.LoginFailur, nil)
		return
	}

	// 存在，不合法会员  { 2 未认证，3 已注销 }
	if member.Status != 1 {
		ctx.Application().Logger().Error("User.Login错误：用户未认证或者已注销")
		response.Failur(ctx, "您还没有认证通过！请联系客服", nil)
		return
	}

	// 合法会员，登录时间,密码,token更新
	member.LoginTime = time.Now()
	member.Token = encrypt.AESEncrypt(user.Mobile, key)
	member.Password = encrypt.AESEncrypt(user.Password, conf.GlobalConfig.JWTSalt)
	columns = append(columns, "login_time", "token", "password")
	_, err = services.NewMemberService().Update(member, columns)
	if err != nil {
		ctx.Application().Logger().Errorf("User.Login更新错误：[%s]", err)
	}

	// 存储合法会员信息
	user.Id = member.Id
	user.Name = member.Name
	user.Gender = member.Gender
	user.Mobile = member.Mobile
	user.Token = member.Token
	jsonU, _ := json.Marshal(user)
	err = redisClient.Set(conf.GlobalConfig.RedisPrefix+user.Token, jsonU, exprire).Err()
	if err != nil {
		ctx.Application().Logger().Errorf("User.Login保存会员信息错误：[%s]", err)
	}

	// 返回user.token 保存
	response.Ok(ctx, response.LoginSuccess, user.Token)
	return
}

/**
用户退出
*/
func Logout(ctx iris.Context) {
	var (
		err  error
		user = new(models.LoginUser)
	)

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.GlobalConfig.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// redis 去除
	err = redisClient.Del(conf.GlobalConfig.RedisPrefix + user.Mobile).Err()
	if err != nil {
		ctx.Application().Logger().Errorf("User.Logout去除错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.LoginOutSuccess, nil)
	return
}

/**
用户信息
*/
func Profile(ctx iris.Context) {
	var (
		user = new(models.LoginUser)
	)

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.GlobalConfig.AuthToken)
	user, _ = tool.GetUserByToken(token)
	response.Ok(ctx, response.OptionSuccess, user)
	return
}

/**
找回密码
*/
func FindUser(ctx iris.Context) {
	var (
		err         error
		has         bool
		code, ip, t string
		times       = 1
		columns     []string
		user        = new(models.LoginUser)
		member      = new(models.Member)
	)

	// 恶意处理
	ip = ips.ClientIP(ctx.Request())
	t, _ = redisClient.Get(conf.GlobalConfig.RedisPrefix + ip).Result()
	if t != "" {
		times, _ = strconv.Atoi(t)
	}

	if times >= 3 {
		ctx.Application().Logger().Errorf("User.FindUser攻击：[%s]", err)
		response.Failur(ctx, "输入错误次数过多，请稍后再试！", "error")
		return
	}

	// 参数
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.Application().Logger().Errorf("User.FindUser参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	// 验证码比对 短信验证码最好
	code, err = redisClient.Get(conf.GlobalConfig.CaptchaKey).Result()
	if code != strings.ToUpper(user.Code) || err != nil {
		ctx.Application().Logger().Errorf("User.FindUser验证码对比错误：[%s]", err)
		response.Failur(ctx, response.CaptchaFailur, nil)
		return
	}

	// 查询用户
	member, has, err = services.NewMemberService().GetMemberByMobile(user.Mobile)
	if err != nil {
		ctx.Application().Logger().Errorf("User.FindUser查询用户错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 不存在，或者手机号不正确
	if !has {
		times++
		redisClient.Set(conf.GlobalConfig.RedisPrefix+ip, times, locktime)
		ctx.Application().Logger().Error("User.FindUser错误：用户不存在，或者登录的手机号不正确")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 存在，姓名比对
	if user.Name != member.Name {
		times++
		redisClient.Set(conf.GlobalConfig.RedisPrefix+ip, times, locktime) //两小时
		ctx.Application().Logger().Error("User.FindUser错误：用户不存在，姓名不正确")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 存在，不合法会员  { 2 未认证，3 已注销 }
	if member.Status != 1 {
		ctx.Application().Logger().Error("User.FindUser错误：用户未认证或者已注销")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 合法，可以更改
	member.Password = encrypt.AESEncrypt(user.Password, conf.GlobalConfig.JWTSalt)
	columns = append(columns, "password", "salt")
	_, err = services.NewMemberService().Update(member, columns)
	if err != nil {
		ctx.Application().Logger().Errorf("User.FindUser修改用户错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}

// 全国省市县
func City(ctx iris.Context) {
	var (
		err error
	)

	// 根据pid查询
	pid, _ := ctx.URLParamInt("aid")
	cityList, err := services.NewZoneService().GetCity(pid)
	if err != nil {
		ctx.Application().Logger().Errorf("User.City查询错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, cityList)
	return
}

// 收货地址列表
func UserAddress(ctx iris.Context) {
	var (
		err  error
		user = new(models.LoginUser)
		md   = new(models.MemberDetail)
	)

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.GlobalConfig.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// 会员地址列表
	md, err = services.NewMemberService().Get(user.Id)
	if err != nil {
		ctx.Application().Logger().Errorf("User.UserAddress查询错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, md.List)
	return
}

/**
地址更新
*/
func SaveAddress(ctx iris.Context) {
	var (
		err     error
		effect  int64
		user    = new(models.LoginUser)
		address = new(models.Address)
		ad      = new(models.AddressDetail)
	)

	if err = ctx.ReadJSON(&ad); err != nil {
		ctx.Application().Logger().Errorf("User.SaveAddress读取地址错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.GlobalConfig.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// 获取会员 mid
	address = ad.Address
	address.Mid = user.Id

	if address.Id > 0 {
		effect, err = services.NewAddressService().Update(address, nil)
	} else {
		effect, err = services.NewAddressService().Create(address)
	}

	if effect < 0 || err != nil {
		ctx.Application().Logger().Errorf("User.SaveAddress保存地址错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}

/**
删除地址
*/
func DeleteAddress(ctx iris.Context) {
	var (
		err    error
		id     int
		effect int64
	)

	id, err = ctx.URLParamInt("id")
	if err != nil {
		ctx.Application().Logger().Errorf("User.DeleteAddress参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err = services.NewAddressService().Delete(id)
	if effect <= 0 || err != nil {
		ctx.Application().Logger().Errorf("User.DeleteAddress删除地址错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}
