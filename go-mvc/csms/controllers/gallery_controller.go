package controllers

import (
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"strconv"

	"go-mvc/framework/conf"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/files"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/thumbnail"
)

type GalleryController struct {
	Ctx iris.Context
}

// 图片表单上传
func (c *GalleryController) PostUpload() {
	file, info, err := c.Ctx.FormFile("file")
	id := c.Ctx.FormValue("id")
	gid := c.Ctx.FormValue("gid")
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Gallery PostUpload 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}
	defer file.Close()

	// 源图
	fileName := thumbnail.ParseName(info.Filename,3)
	filePath, err1 := files.MakeFilePath(conf.GetUploadFile() + conf.GlobalConfig.UploadPicPath[0], fileName)
	if err1 != nil {
		c.Ctx.Application().Logger().Errorf("Gallery PostUpload 目录：[%s]", err1)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	out, err2 := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		c.Ctx.Application().Logger().Errorf("Gallery PostUpload 保存图片：[%s]", err2)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	defer out.Close()
	_, _ = io.Copy(out, file)

	// 生成缩略图
	small := thumbnail.ThumbnailSave(fileName,100, 100, 1)
	medium := thumbnail.ThumbnailSave(fileName,360, 360, 2)

	// 入库
	gallery := new(models.GoodsGallery)
	gallery.GalleryId = id
	gallery.Small = small
	gallery.Medium = medium
	gallery.Source = fileName
	gallery.GoodsId,_ = strconv.Atoi(gid)

	effect,err3 := services.NewGalleryService().Create(gallery)
	if effect < 0 || err3 != nil {
		c.Ctx.Application().Logger().Errorf("Gallery PostUpload 添加：[%s]", err3)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, gallery)
	return
}

func (c *GalleryController) PostDelete() {
	var (
		err    error
		effect int64
		gallery = new(models.GoodsGallery)
	)

	if err = c.Ctx.ReadJSON(&gallery); err != nil {
		c.Ctx.Application().Logger().Errorf("Gallery PostDelete Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	effect, err = services.NewGalleryService().Delete(gallery.GalleryId)
	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Gallery PostDelete 删除：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 删除图片
	goodsPath :=  conf.GetUploadFile() + conf.GlobalConfig.UploadPicPath[0]
	files.RemoveFile(goodsPath + gallery.Small)
	files.RemoveFile(goodsPath + gallery.Medium)
	files.RemoveFile(goodsPath + gallery.Source)

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}
