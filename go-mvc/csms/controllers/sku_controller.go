package controllers

import (
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/response"
	"github.com/kataras/iris/v12"
)

type SkuController struct {
	Ctx iris.Context
}

//save
func (c *SkuController) PostVal() {
	var (
		err    error
		//effect int64
		skuValList = make([]models.GoodsSkuVal, 0)
	)

	if err = c.Ctx.ReadJSON(&skuValList); err != nil {
		c.Ctx.Application().Logger().Errorf("Sku PostVal Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	for _,v := range skuValList {
		if v.Vid > 0 {
			services.NewSkuService().UpdateVal(&v, nil)
		} else {
			services.NewSkuService().CreateVal(&v)
		}
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}
