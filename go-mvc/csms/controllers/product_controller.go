package controllers

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/conf"
	"go-mvc/framework/logs"
	"go-mvc/framework/models/common"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/files"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/thumbnail"
	"io"
	"os"
	"strconv"
	"strings"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.GoodsService
}

// list
func (c *ProductController) GetList() {
	var (
		err      error
		name     string
		category int
		isOnSale int
		isFirst  int
		isHot    int
		p        *page.Pagination
		res      *page.Result
		list     []models.GoodsDetail
		total    int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	name = c.Ctx.URLParam("name")
	category, _ = c.Ctx.URLParamInt("categoryId")
	isOnSale, _ = c.Ctx.URLParamInt("isOnSale")
	isFirst, _ = c.Ctx.URLParamInt("isFirst")
	isHot, _ = c.Ctx.URLParamInt("isHot")

	list, total, err = c.Service.List(name, category, isOnSale, isFirst, isHot, p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询")
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

// item
func (c *ProductController) GetItem() {
	var (
		err         error
		id          int
		goods       = new(models.Goods)
		galleryList = make([]models.GoodsGallery, 0)
		skuValList  = make([]models.GoodsSkuVal, 0)
		maps        = make(map[string]interface{}, 0)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	goods, galleryList, _, skuValList, err = c.Service.Get(id)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	maps["goods"] = goods
	maps["galleryList"] = galleryList
	maps["skuValList"] = skuValList

	response.Ok(c.Ctx, response.OptionSuccess, maps)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// save
func (c *ProductController) PostSave() {
	var (
		err    error
		effect int64
		goods  = new(models.Goods)
	)

	if err = c.Ctx.ReadJSON(&goods); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	if goods.Id > 0 {
		//goods.UpdateTime = time.Now()
		effect, err = c.Service.Update(goods, nil)
	} else {
		//goods.CreateTime = time.Now()
		effect, err = c.Service.Create(goods)
	}

	if effect < 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "保存失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

// kindeditor图片上传
func (c *ProductController) PostUpload() {
	var (
		res = new(common.Kindeditor)
	)

	// 图片表单上传
	file, info, err := c.Ctx.FormFile("imgFile")
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}
	defer file.Close()

	// 源图
	fileName := thumbnail.ParseName(info.Filename, 0)
	filePath, err1 := files.MakeFilePath(conf.GetUploadFile()+conf.GlobalConfig.UploadEditor[0], fileName)
	if err1 != nil {
		logs.GetLogger().Error(logs.D{"err": err1}, "创建目录出错")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	out, err2 := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		logs.GetLogger().Error(logs.D{"err": err2}, "生成图片出错")
		res.Error = 1
		res.Message = err2.Error()
		_, _ = c.Ctx.JSON(res)
		return
	}

	defer out.Close()
	_, _ = io.Copy(out, file)

	res.Error = 0
	res.Url = conf.GlobalConfig.UploadUrl + conf.GlobalConfig.UploadEditor[0] + fileName
	_, _ = c.Ctx.JSON(res)
	return
}

//  delete?id=1,2
func (c *ProductController) GetDelete() {
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
		logs.GetLogger().Error(logs.D{"err": err}, "删除失败")
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
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

// close?id=1,2
func (c *ProductController) GetClose() {
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
		logs.GetLogger().Error(logs.D{"err": err}, "close失败")
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
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
