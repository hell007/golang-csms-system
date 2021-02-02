package controllers

import (
	"errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/conf"
	"go-mvc/framework/logs"
	models "go-mvc/framework/models/common"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/thumbnail"
)

type TestController struct {
	Ctx     iris.Context
	Service services.UserService
}

// 图片上传
func (c *TestController) PostUpload2() {
	var (
		err         error
		uploadfiles = make([]models.UploadFile, 0) //接收数组
		path        string
		paths       = make([]string, 0)
	)

	// 读取
	if err = c.Ctx.ReadJSON(&uploadfiles); err != nil {
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, err)
		return
	}

	// 路径
	fileDir := conf.GetUploadFile()

	// 获取返回的上传路径
	for _, v := range uploadfiles {
		path, err = thumbnail.Base64Upload(fileDir, v.FileData)
		paths = append(paths, path)
	}

	response.Ok(c.Ctx, response.OptionSuccess, paths)
}

//本应用写日志的三种方式
func TestLog(ctx iris.Context) {
	err := errors.New("出错了")
	golog.Println("golog日志")
	ctx.Application().Logger().Error("%s", err, "ctx调用golog日志")
	logs.GetLogger().Error(logs.D{"err": err}, "logrus日志")
}