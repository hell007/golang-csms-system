package apis

import (
	"github.com/kataras/iris/v12"
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
		ctx.Application().Logger().Errorf("Category.CategoryList查询商品分类错误：[%s]", err)
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
		ctx.Application().Logger().Errorf("Category.CategoryGoods参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	list, total, err := services.NewGoodsService().GetGoods("category", strconv.Itoa(id), p)
	if err != nil {
		ctx.Application().Logger().Errorf("Category.CategoryGoods查询商品错误：[%s]", err)
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
