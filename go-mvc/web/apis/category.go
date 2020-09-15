package apis

import (
	"github.com/kataras/iris/v12"
	"strconv"

	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

func CategoryList(ctx iris.Context) {
	var (
		val string
		maps = make(map[string]interface{}, 0)
		p = new(page.Pagination)
	)

	categoryList, err := services.NewCategoryService().List(0)
	if err != nil {
		ctx.Application().Logger().Errorf("Category.CategoryList 查询：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}
	maps["categoryList"] = categoryList

	if len(categoryList) > 0 {
		val =  strconv.Itoa(categoryList[0].Id)
	}

	p.PageNumber = 1
	p.PageSize = 100
	goodsList,_,_ := services.NewGoodsService().GetGoods("category",val, p)
	maps["goodsList"] = goodsList

	response.Ok(ctx, response.OptionSuccess, maps)
	return
}

func CategoryGoods(ctx iris.Context) {
	var (
		err error
		id  int
		p     *page.Pagination
		res   *page.Result
	)

	p, err = page.NewPagination(ctx)
	id, err = ctx.URLParamInt("id")
	if err != nil {
		ctx.Application().Logger().Errorf("Category.CategoryGoods 参数：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	list, total, err := services.NewGoodsService().GetGoods("category", strconv.Itoa(id), p)
	if err != nil {
		ctx.Application().Logger().Errorf("Category.CategoryGoods 查询：[%s]", err)
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