package apis

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"image/png"
	"strconv"
	"strings"
	"time"

	"go-mvc/framework/conf"
	models "go-mvc/framework/models/member"
	"go-mvc/framework/models/zone"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/captcha"
	"go-mvc/framework/utils/encrypt"
	"go-mvc/framework/utils/ips"
	redisClient "go-mvc/framework/utils/redis"
	"go-mvc/framework/utils/response"
)

type User struct {
	Ctx iris.Context
}

const (
	captchaKey string = "CAPTCHA"
)
var key = conf.JWTSecret + time.Now().Format("20060102-15:04:05") //24
var captchatime = 5 * 60 * time.Second //5分钟
var locktime =  2 * 60 * 60 * time.Second //2小时

// 会员注册
func (c *User) PostRegister() {
	var (
		err     error
		has     bool
		effect  int64
		keys    string
		user = new(models.Member)
		member = new(models.Member)
	)

	// 读取
	if err = c.Ctx.ReadJSON(&member); err != nil {
		c.Ctx.Application().Logger().Errorf("手机号为[%s]的用户，注册失败：[%s]", member.Mobile, err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.RegisteFailur, nil)
		return
	}

	// 是否已存在
	user.Mobile = member.Mobile
	has, err = services.NewMemberService().Exist(user)
	if has || err != nil {
		c.Ctx.Application().Logger().Errorf("手机号为[%s]的用户，已存在：[%s]", user.Mobile, err)
		response.Failur(c.Ctx, response.RegisteExist, nil)
		return
	}

	// 获取用户ip
	ip := ips.ClientIP(c.Ctx.Request())
	member.Ip = ip

	// 数据处理
	keys = key //保证key相同
	member.Salt = encrypt.AESEncrypt(keys, conf.JWTSalt)
	member.Password = encrypt.AESEncrypt(member.Password, keys)
	member.CreateTime = time.Now()
	member.Status = 2 //未认证

	effect, err = services.NewMemberService().Create(member)
	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("手机号为[%s]的用户，注册失败：[%s]", member.Mobile, err)
		response.Failur(c.Ctx, response.RegisteFailur, nil)
		return
	}

	response.Ok_(c.Ctx, response.RegisteSuccess)
	return
}

// 验证码
func (c *User) GetCaptcha() {
	cp := captcha.NewCaptcha(120, 40, 4)
	cp.SetMode(0)
	code, img  := cp.OutPut()

	// 保存code 60s
	err := redisClient.Set(conf.RedisPrefix + captchaKey, code, captchatime).Err()
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User GetCaptcha 验证码：[%s]", err)
	}
	c.Ctx.Header("Content-Type", "image/png; charset=utf-8")
	png.Encode(c.Ctx, img)
}

// 会员登录
func (c *User) PostLogin() {
	var (
		err     error
		has, ckPassword bool
		code, keys string
		columns []string
		user = new(models.LoginUser)
		member = new(models.Member)
	)

	// 参数
	if err = c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("User Login 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.LoginFailur, nil)
		return
	}

	// 验证码比对
	code, err = redisClient.Get(conf.RedisPrefix + captchaKey).Result()
	if code != strings.ToUpper(user.Code)  || err != nil {
		c.Ctx.Application().Logger().Errorf("User Login 验证码对比：[%s]", err)
		response.Failur(c.Ctx, response.CaptchaFailur, nil)
		return
	}

	// 查询用户
	member, has, err = services.NewMemberService().GetMemberByMobile(user.Mobile)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User Login 查询：[%s]", err)
		response.Failur(c.Ctx, response.LoginFailur, nil)
		return
	}

	// 不存在，或者手机号不正确
	if !has {
		c.Ctx.Application().Logger().Error("用户不存在，或者登录的手机号不正确")
		response.Failur(c.Ctx, response.LoginFailur, nil)
		return
	}

	// 存在，验证密码
	salt := encrypt.AESDecrypt(member.Salt, conf.JWTSalt)
	ckPassword = encrypt.CheckPWD(user.Password, member.Password, salt)
	if !ckPassword {
		c.Ctx.Application().Logger().Error("用户存在，登录的密码不正确")
		response.Failur(c.Ctx, response.LoginFailur, nil)
		return
	}

	// 存在，不合法会员  { 2 未认证，3 已注销 }
	if (member.Status != 1) {
		c.Ctx.Application().Logger().Error("用户未认证或者已注销")
		response.Failur(c.Ctx, "您还没有认证通过！请联系客服", nil)
		return
	}

	// 合法会员，登录时间,盐值,密码更新
	keys = key
	member.LoginTime = time.Now()
	member.Salt = encrypt.AESEncrypt(keys, conf.JWTSalt)
	member.Password = encrypt.AESEncrypt(user.Password, keys)
	columns = append(columns, "login_time", "salt", "password")
	_, err = services.NewMemberService().Update(member, columns)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User PostLogin 登录：[%s]", err)
	}

	// 存储合法会员信息
	user.Id = member.Id
	user.Name = member.Name
	user.Gender = member.Gender
	user.Password = ""
	user.Token = encrypt.AESEncrypt(user.Mobile, conf.JWTSalt)
	jsonU, _ := json.Marshal(user)
	err = redisClient.Set(conf.RedisPrefix + user.Mobile, jsonU, conf.RedisExprire).Err()
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User PostLogin 保存会员信息：[%s]", err)
	}

	// 返回user.token 保存
	response.Ok(c.Ctx, response.LoginSuccess, user.Token)
	return
}

// 用户退出
func (c *User) GetLogout() {
	var (
		err     error
		token, keys, jsonU string
		user = new(models.LoginUser)
	)

	// 参数
	token = c.Ctx.URLParam("token")
	if token == ""{
		c.Ctx.Application().Logger().Errorf("User Logout 参数：[%s]", err)
		response.Failur(c.Ctx,  response.OptionFailur, nil)
		return
	}

	// 解密
	keys = encrypt.AESDecrypt(token, conf.JWTSalt)
	jsonU, err = redisClient.Get(conf.RedisPrefix + keys).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		c.Ctx.Application().Logger().Errorf("User Logout 解密：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// redis 去除
	err = redisClient.Del(conf.RedisPrefix + user.Mobile).Err()
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User Logout 去除：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.LoginOutSuccess,nil)
	return
}

// 用户信息
func (c *User) GetV1() {
	var (
		err     error
		token, keys, jsonU string
		user = new(models.LoginUser)
	)

	// 参数
	token = c.Ctx.URLParam("token")
	if token == ""{
		c.Ctx.Application().Logger().Errorf("User V1 参数：[%s]", err)
		response.Failur(c.Ctx,  response.OptionFailur, nil)
		return
	}

	// 解密
	keys = encrypt.AESDecrypt(token, conf.JWTSalt)
	jsonU, err = redisClient.Get(conf.RedisPrefix + keys).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		c.Ctx.Application().Logger().Errorf("User V1 解密：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 处理信息
	user.Password = ""
	response.Ok(c.Ctx, response.OptionSuccess, user)
	return
}

// 找回密码
func (c *User) PostFind() {
	var (
		err error
		has bool
		code, ip, t, keys string
		times = 1
		columns []string
		user = new(models.LoginUser)
		member = new(models.Member)
	)

	// 恶意处理
	ip = ips.ClientIP(c.Ctx.Request())
	t, _ = redisClient.Get(conf.RedisPrefix + ip).Result()
	if t != "" {
		times, _ = strconv.Atoi(t)
	}

	if times >= 3 {
		c.Ctx.Application().Logger().Errorf("User Find 攻击：[%s]", err)
		response.Failur(c.Ctx, "输入错误次数过多，请稍后再试！", "error")
		return
	}

	// 参数
	if err = c.Ctx.ReadJSON(&user); err != nil {
		c.Ctx.Application().Logger().Errorf("User Find 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	// 验证码比对 短信验证码最好
	code, err = redisClient.Get(conf.RedisPrefix + captchaKey).Result()
	if code != strings.ToUpper(user.Code)  || err != nil {
		c.Ctx.Application().Logger().Errorf("User Find 验证码对比：[%s]", err)
		response.Failur(c.Ctx, response.CaptchaFailur, nil)
		return
	}

	// 查询用户
	member, has, err = services.NewMemberService().GetMemberByMobile(user.Mobile)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User Find 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 不存在，或者手机号不正确
	if !has {
		times++
		redisClient.Set(conf.RedisPrefix + ip, times, locktime)
		c.Ctx.Application().Logger().Error("用户不存在，或者登录的手机号不正确")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 存在，姓名比对
	if user.Name != member.Name {
		times++
		redisClient.Set(conf.RedisPrefix + ip, times, locktime) //两小时
		c.Ctx.Application().Logger().Error("用户不存在，姓名不正确")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 存在，不合法会员  { 2 未认证，3 已注销 }
	if (member.Status != 1) {
		c.Ctx.Application().Logger().Error("用户未认证或者已注销")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 合法，可以更改
	keys = key
	member.Salt = encrypt.AESEncrypt(keys, conf.JWTSalt)
	member.Password = encrypt.AESEncrypt(user.Password, keys)
	columns = append(columns, "password", "salt")
	_, err = services.NewMemberService().Update(member, columns)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("User Find 修改：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

// 地址
func (c *User) GetCity() {
	var (
		err error
		jsonP, jsonC, jsonA string
		province, city, area []zone.City
		maps      = make(map[string]interface{}, 0)
	)
	// redis
	jsonP, err = redisClient.Get(conf.RedisPrefix + "province").Result()
	jsonC, err = redisClient.Get(conf.RedisPrefix + "city").Result()
	jsonA, err = redisClient.Get(conf.RedisPrefix + "area").Result()
	err = json.Unmarshal([]byte(jsonP), &province)
	err = json.Unmarshal([]byte(jsonC), &city)
	err = json.Unmarshal([]byte(jsonA), &area)

	// redis不存在
	if err != nil {
		province, err = services.NewZoneService().GetCity(1)
		city, err = services.NewZoneService().GetCity(2)
		area, err = services.NewZoneService().GetCity(3)
		if err != nil {
			c.Ctx.Application().Logger().Errorf("User GetCity 查询地址出错：[%s]", err)
			response.Failur(c.Ctx, response.OptionFailur, nil)
			return
		}

		jsonP, _ := json.Marshal(province)
		err = redisClient.Set(conf.RedisPrefix+"province", jsonP, conf.RedisExprire).Err()
		jsonC, _ := json.Marshal(city)
		err = redisClient.Set(conf.RedisPrefix+"city", jsonC, conf.RedisExprire).Err()
		jsonA, _ := json.Marshal(area)
		err = redisClient.Set(conf.RedisPrefix+"area", jsonA, conf.RedisExprire).Err()
		if err != nil {
			c.Ctx.Application().Logger().Errorf("User GetCity redis：[%s]", err)
		}

		maps["province"] = province
		maps["city"] = city
		maps["area"] = area
		response.Ok(c.Ctx, response.OptionSuccess, maps)
		return
	}

	//redis存在
	maps["province"] = province
	maps["city"] = city
	maps["area"] = area
	response.Ok(c.Ctx, response.OptionSuccess, maps)
	return
}

// 收货地址列表
func (c *User) GetAddress() {
	var (
		err     error
		token, key, jsonU   string
		user   = new(models.LoginUser)
		md = new(models.MemberDetail)
	)

	// 参数
	token = c.Ctx.URLParam("uid")
	if token == ""{
		c.Ctx.Application().Logger().Errorf("User Address 参数：[%s]", err)
		response.Failur(c.Ctx,  response.OptionFailur, nil)
		return
	}

	// 解密
	key = encrypt.AESDecrypt(token, conf.JWTSalt)
	jsonU, err = redisClient.Get(conf.RedisPrefix + key).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		c.Ctx.Application().Logger().Errorf("User Address 解密：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 会员地址
	md, err = services.NewMemberService().Get(user.Id)
	response.Ok(c.Ctx, response.OptionSuccess, md.List)
	return
}

// 地址更新
func (c *User) PostSaveaddr() {
	var (
		err     error
		effect int64
		jsonU string
		user = new(models.LoginUser)
		address = new(models.Address)
		ad = new(models.AddressDetail)
	)

	if err = c.Ctx.ReadJSON(&ad); err != nil {
		c.Ctx.Application().Logger().Errorf("User PostSaveaddr Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	// 解密
	key = encrypt.AESDecrypt(ad.Token, conf.JWTSalt)
	jsonU, err = redisClient.Get(conf.RedisPrefix + key).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		c.Ctx.Application().Logger().Errorf("User PostSaveaddr 解密：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 获取会员 mid
	address = ad.Address
	address.Mid = user.Id

	if address.Id > 0 {
		effect, err = services.NewAddressService().Update(address, nil)
	} else {
		effect, err = services.NewAddressService().Create(address)
	}

	if effect < 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("User PostSaveaddr 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

// 删除地址
func (c *User) GetDeladdr() {
	var (
		err    error
		id     int
		effect int64
	)

	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		c.Ctx.Application().Logger().Errorf("UserController PostDeladdr 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err = services.NewAddressService().Delete(id)
	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("User PostDeladdr 删除：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}