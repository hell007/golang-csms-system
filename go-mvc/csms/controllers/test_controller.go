package controllers

import (
	"go-mvc/framework/conf"
	models "go-mvc/framework/models/common"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/thumbnail"
	"github.com/kataras/iris/v12"
)

type TestController struct {
	Ctx iris.Context
	Service services.GoodsService
}

func (c *TestController) GetTest() {

	//fmt.Println(d.Master)
	//fmt.Println(d.Master.Host)
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

func (c *TestController) GetLog() {

	id, err := c.Ctx.URLParamInt("id")

	if err == nil {
		c.Ctx.Application().Logger().Errorf("test：[%s]", "333")

		response.Ok(c.Ctx, response.OptionFailur, id)
		return
	}

}