package controllers

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
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
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}
	defer file.Close()

	// 源图
	fileName := thumbnail.ParseName(info.Filename, 3)
	filePath, err1 := files.MakeFilePath(conf.Global.Directory+conf.Global.UploadPicPath[0], fileName)
	if err1 != nil {
		logs.GetLogger().Error(logs.D{"err": err1}, "图片创建目录出错")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	out, err2 := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		logs.GetLogger().Error(logs.D{"err": err2}, "图片生成出错")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	defer out.Close()
	_, _ = io.Copy(out, file)

	// 生成缩略图
	small := thumbnail.ThumbnailSave(fileName, 100, 100, 1)
	medium := thumbnail.ThumbnailSave(fileName, 360, 360, 2)

	// 入库
	gallery := new(models.GoodsGallery)
	gallery.GalleryId = id
	gallery.Small = small
	gallery.Medium = medium
	gallery.Source = fileName
	gallery.GoodsId, _ = strconv.Atoi(gid)

	effect, err3 := services.NewGalleryService().Create(gallery)
	if effect < 0 || err3 != nil {
		logs.GetLogger().Error(logs.D{"err": err3}, "图片添加入库失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, gallery)
	return
}

func (c *GalleryController) PostDelete() {
	var (
		err     error
		effect  int64
		gallery = new(models.GoodsGallery)
	)

	if err = c.Ctx.ReadJSON(&gallery); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err = services.NewGalleryService().Delete(gallery.GalleryId)
	if effect <= 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "删除图片集失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 删除图片
	goodsPath := conf.Global.Directory + conf.Global.UploadPicPath[0]
	files.RemoveFile(goodsPath + gallery.Small)
	files.RemoveFile(goodsPath + gallery.Medium)
	files.RemoveFile(goodsPath + gallery.Source)

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}
