package apis

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	"strconv"

	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

/**
商品分类
*/
func CategoryList(ctx iris.Context) {
	categoryList, err := services.NewCategoryService().List(0)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询商品分类错误")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, categoryList)
	return
}

/**
商品分类商品列表
*/
func CategoryGoods(ctx iris.Context) {
	var (
		err error
		id  int
		p   *page.Pagination
		res *page.Result
	)

	p, err = page.NewPagination(ctx)
	id, err = ctx.URLParamInt("id")
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	list, total, err := services.NewGoodsService().GetGoods("category", strconv.Itoa(id), p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询商品错误")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Total: total,
		Rows:  list,
	}

	response.Ok(ctx, response.OptionSuccess, res)
	return
}
