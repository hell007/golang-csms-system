package controllers

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/conf"
	models "go-mvc/framework/models/common"
	"go-mvc/framework/models/system"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/thumbnail"
)

type TestController struct {
	Ctx iris.Context
	Service services.UserService
}

func TestList(ctx iris.Context) {
	var (
		err  error
		has  bool
		id   int
		user = new(system.User)
	)

	// 参数处理
	id, err = ctx.URLParamInt("id")
	if err != nil {
		ctx.Application().Logger().Errorf("TestController.Test参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	user, has, err = services.NewUserService().Get(id)
	if !has {
		ctx.Application().Logger().Errorf("TestController.GetItem查询错误：[%s]", err)
		golog.Error("UserController GetItem：" + err.Error())
		response.Failur(ctx,  response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, user)
	return
}

func GetLog(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.Application().Logger().Errorf("test参数错误：[%s]", "test")
		response.Failur(ctx, response.OptionFailur, id)
		return
	}
	response.Ok(ctx, response.OptionSuccess, id)
	return
}

// 图片上传
func (c *TestController) PostUpload2() {
	var (
		err error
		uploadfiles  = make([]models.UploadFile, 0) //接收数组
		path string
		paths = make([]string, 0)
	)

	// 读取
	if err = c.Ctx.ReadJSON(&uploadfiles); err != nil {
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, err)
		return
	}

	// 路径
	fileDir := conf.GetUploadFile()

	// 获取返回的上传路径
	for _,v := range uploadfiles {
		path, err = thumbnail.Base64Upload(fileDir, v.FileData)
		paths = append(paths, path)
	}

	response.Ok(c.Ctx, response.OptionSuccess, paths)
}

