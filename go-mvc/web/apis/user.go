package apis

import (
	"encoding/json"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	"go-mvc/framework/utils/tool"
	"image/png"
	"reflect"
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
var key = conf.Global.JWTSecret + time.Now().Format(conf.Global.Timeformat)

//5分钟
var captchatime = time.Duration(conf.Global.CaptchaExprire) * time.Second

//2小时
var locktime = 2 * time.Hour

//1小时
var exprire = time.Duration(conf.Global.RedisExprire) * time.Hour

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
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.RegisteFailur, nil)
		return
	}

	//字段验证
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	//验证器注册翻译器
	_ = translations.RegisterDefaultTranslations(validate, trans)
	err = validate.Struct(member)
	if err != nil {
		for _, err2 := range err.(validator.ValidationErrors) {
			logs.GetLogger().Error(nil, "validator验证出错")
			response.Failur(ctx, response.RegisteFailur, err2.Translate(trans))
			return
		}
	}

	// 是否已存在
	user.Mobile = member.Mobile
	has, err = services.NewMemberService().Exist(user)
	if has || err != nil {
		logs.GetLogger().Error(logs.D{"err": user.Mobile}, "用户已存在")
		response.Failur(ctx, response.RegisteExist, nil)
		return
	}

	// 获取用户ip
	ip := ips.ClientIP(ctx.Request())
	member.Ip = ip

	// 数据处理
	member.Password = encrypt.AESEncrypt(member.Password, conf.Global.JWTSalt)
	member.CreateTime = time.Now()
	member.Status = 1

	effect, err = services.NewMemberService().Create(member)
	if effect <= 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "用户注册失败")
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
	err := redisClient.Set(conf.Global.CaptchaKey, code, captchatime).Err()
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "验证码缓存错误")
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
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.LoginFailur, nil)
		return
	}

	// 如果开启验证码校验，进行验证码校对
	if conf.Global.CaptchaVerify {
		code, err = redisClient.Get(conf.Global.CaptchaKey).Result()
		if code != strings.ToUpper(user.Code) || err != nil {
			logs.GetLogger().Error(logs.D{"err": err}, "验证码对比错误")
			response.Failur(ctx, response.CaptchaFailur, nil)
			return
		}
	}

	// 查询用户
	member, has, err = services.NewMemberService().GetMemberByMobile(user.Mobile)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询用户错误")
		response.Failur(ctx, response.LoginFailur, nil)
		return
	}

	// 不存在，或者手机号不正确
	if !has {
		logs.GetLogger().Error(nil, "用户不存在，或者登录的手机号不正确")
		response.Failur(ctx, response.LoginFailur, nil)
		return
	}

	// 存在，验证密码
	equal = encrypt.CheckPWD(user.Password, member.Password, conf.Global.JWTSalt)
	if !equal {
		logs.GetLogger().Error(nil, "用户存在，登录的密码不正确")
		response.Failur(ctx, response.LoginFailur, nil)
		return
	}

	// 存在，不合法会员  { 2 未认证，3 已注销 }
	if member.Status != 1 {
		logs.GetLogger().Error(nil, "用户未认证或者已注销")
		response.Failur(ctx, "您还没有认证通过！请联系客服", nil)
		return
	}

	// 合法会员，登录时间,密码,token更新
	member.LoginTime = time.Now()
	member.Token = encrypt.AESEncrypt(user.Mobile, key)
	member.Password = encrypt.AESEncrypt(user.Password, conf.Global.JWTSalt)
	columns = append(columns, "login_time", "token", "password")
	_, err = services.NewMemberService().Update(member, columns)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "更新出错")
	}

	// 存储合法会员信息
	user.Id = member.Id
	user.Name = member.Name
	user.Gender = member.Gender
	user.Mobile = member.Mobile
	user.Token = member.Token
	jsonU, _ := json.Marshal(user)
	err = redisClient.Set(conf.Global.RedisPrefix+user.Token, jsonU, exprire).Err()
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "缓存会员信息出错")
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
	token := ctx.GetHeader(conf.Global.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// redis 去除
	err = redisClient.Del(conf.Global.RedisPrefix + user.Mobile).Err()
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "删除缓存用户信息出错")
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
	token := ctx.GetHeader(conf.Global.AuthToken)
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
	t, _ = redisClient.Get(conf.Global.RedisPrefix + ip).Result()
	if t != "" {
		times, _ = strconv.Atoi(t)
	}

	if times >= 3 {
		logs.GetLogger().Error(nil, "输入错误次数过多")
		response.Failur(ctx, "输入错误次数过多，请稍后再试！", "error")
		return
	}

	// 参数
	if err = ctx.ReadJSON(&user); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	// 验证码比对 短信验证码最好
	code, err = redisClient.Get(conf.Global.CaptchaKey).Result()
	if code != strings.ToUpper(user.Code) || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "验证码对比错误")
		response.Failur(ctx, response.CaptchaFailur, nil)
		return
	}

	// 查询用户
	member, has, err = services.NewMemberService().GetMemberByMobile(user.Mobile)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询用户错误")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 不存在，或者手机号不正确
	if !has {
		times++
		redisClient.Set(conf.Global.RedisPrefix+ip, times, locktime)
		logs.GetLogger().Error(nil, "用户不存在，或者登录的手机号不正确")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 存在，姓名比对
	if user.Name != member.Name {
		times++
		redisClient.Set(conf.Global.RedisPrefix+ip, times, locktime) //两小时
		logs.GetLogger().Error(nil, "用户不存在，姓名不正确")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 存在，不合法会员  { 2 未认证，3 已注销 }
	if member.Status != 1 {
		logs.GetLogger().Error(nil, "用户未认证或者已注销")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 合法，可以更改
	member.Password = encrypt.AESEncrypt(user.Password, conf.Global.JWTSalt)
	columns = append(columns, "password", "salt")
	_, err = services.NewMemberService().Update(member, columns)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "修改用户错误")
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
	if pid < 0 {
		logs.GetLogger().Error(nil, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	cityList, err := services.NewZoneService().GetCity(pid)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询错误")
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
	token := ctx.GetHeader(conf.Global.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// 会员地址列表
	md, err = services.NewMemberService().Get(user.Id)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询错误")
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
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.Global.AuthToken)
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
		logs.GetLogger().Error(logs.D{"err": err}, "保存地址错误")
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
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err = services.NewAddressService().Delete(id)
	if effect <= 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "删除地址错误")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}
