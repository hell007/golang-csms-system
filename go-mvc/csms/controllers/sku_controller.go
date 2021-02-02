package controllers

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/response"
)

type SkuController struct {
	Ctx iris.Context
}

//save
func (c *SkuController) PostVal() {
	var (
		err error
		skuValList = make([]models.GoodsSkuVal, 0)
	)

	if err = c.Ctx.ReadJSON(&skuValList); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	for _, v := range skuValList {
		if v.Vid > 0 {
			services.NewSkuService().UpdateVal(&v, nil)
		} else {
			services.NewSkuService().CreateVal(&v)
		}
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}
