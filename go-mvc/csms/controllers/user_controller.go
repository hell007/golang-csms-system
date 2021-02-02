package controllers

import (
	"go-mvc/framework/logs"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"

	"go-mvc/framework/conf"
	"go-mvc/framework/middleware/jwt"
	models "go-mvc/framework/models/system"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/encrypt"
	"go-mvc/framework/utils/ips"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

// user/login
// 注意：后台用户登录没有使用salt盐值加密密码，仅仅使用/conf里面的设置值
func (c *UserController) PostLogin() {
	var (
		err        error
		columns    []string
		ckPassword bool
		token      string
		user       = new(models.User)
		ut         = new(models.UserToken) //需要返回的组装user
	)

	if err = c.Ctx.ReadJSON(&user); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	ut.Username = user.Username
	has, err := c.Service.GetUserByName(user.Username, ut)

	if err != nil {
		logs.GetLogger().Error(logs.D{"user": user.Username, "err": err}, "登录失败")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.LoginFailur, nil)
		return
	}

	// 用户名不正确
	if !has {
		logs.GetLogger().Error(nil, "用户名不正确")
		response.Failur(c.Ctx, response.UsernameFailur, nil)
		return
	}

	// 验证密码
	ckPassword = encrypt.CheckPWD(user.Password, ut.Password, conf.Global.JWTSalt)
	if !ckPassword {
		logs.GetLogger().Error(nil, "密码不正确")
		response.Failur(c.Ctx, response.PasswordFailur, nil)
		return
	}

	// 用户存在，生成token
	token, err = jwt.GenerateToken(ut)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "token生成失败")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenCreateFailur, nil)
		return
	}

	// 更新登录时间，盐值, ip
	// 获取用户ip
	ip := ips.ClientIP(c.Ctx.Request())
	user.Ip = ip
	user.LoginTime = time.Now()
	user.Id = ut.Id
	columns = append(columns, "ip", "login_time")
	_, err = c.Service.Update(user, columns)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "更新登录信息出错")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenCreateFailur, nil)
		return
	}

	ut.Token = token
	ut.Password = ""
	response.Ok(c.Ctx, response.LoginSuccess, ut)
}

// user/loginout token过期
func (c *UserController) PostLoginout() {
	response.Ok(c.Ctx, response.LoginOutSuccess, nil)
	return
}

// user/test 刷新token
func (c *UserController) GetTest() {
	var (
		err      error
		oldtoken string
		token    string
	)

	oldtoken, err = jwt.FromAuthHeader(c.Ctx)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.TokenParseFailur)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenParseFailur, nil)
		return
	}

	token, err = jwt.RefreshToken(oldtoken)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.TokenParseFailur)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.TokenRefreshFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.LoginSuccess, token)
	return
}

// user/list?pageNumber=1&pageSize=2&name=曹操
func (c *UserController) GetList() {
	var (
		err      error
		username string
		status   int
		p        *page.Pagination
		res      *page.Result
		list     []models.UserDetail
		total    int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	username = c.Ctx.URLParam("name")
	status, _ = c.Ctx.URLParamInt("status")
	list, total, err = c.Service.List(username, status, p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询出错")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Total: total,
		Rows:  list,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// user/item?id=1
func (c *UserController) GetItem() {
	var (
		err  error
		has  bool
		id   int
		user = new(models.User)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	user, has, err = c.Service.Get(id)
	if !has {
		logs.GetLogger().Error(logs.D{"err": err}, "查询出错")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, user)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// user/save
func (c *UserController) PostSave() {
	var (
		err    error
		effect int64
		user   = models.User{}
	)

	if err = c.Ctx.ReadJSON(&user); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	if user.Id > 0 {
		effect, err = c.Service.Update(&user, nil)
	} else {
		user.CreateTime = time.Now()
		effect, err = c.Service.Create(&user)
	}

	if effect < 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "保存出错")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}

// user/delete?id=1,2
func (c *UserController) GetDelete() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int, 0)
		uid    int
		effect int64
	)

	id = c.Ctx.URLParam("id")
	idList = strings.Split(id, ",")
	if len(idList) == 0 {
		goto FAIL
	}

	for _, v := range idList {
		if v == "" {
			continue
		}

		uid, err = strconv.Atoi(v)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, uid)
	}

	effect, err = c.Service.Delete(ids)
	if effect <= 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "删除出错")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

func (c *UserController) GetClose() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int, 0)
		uid    int
		effect int64
	)

	id = c.Ctx.URLParam("id")
	idList = strings.Split(id, ",")
	if len(idList) == 0 {
		goto FAIL
	}

	for _, v := range idList {
		if v == "" {
			continue
		}

		uid, err = strconv.Atoi(v)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, uid)
	}

	effect, err = c.Service.Close(ids)

	if effect <= 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.OptionFailur)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}
