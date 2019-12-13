package apis

import (
	"github.com/kataras/iris/v12"
	"strconv"

	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type Category struct {
	Ctx iris.Context
}

func (c Category) GetV1() {
	var (
		val string
		maps = make(map[string]interface{}, 0)
		p = new(page.Pagination)
	)

	categoryList, err := services.NewCategoryService().List(0)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Category GetV1 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
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

	response.Ok(c.Ctx, response.OptionSuccess, maps)
	return
}

func (c Category) GetGoods() {
	var (
		err error
		id  int
		p     *page.Pagination
		res   *page.Result
	)

	p, err = page.NewPagination(c.Ctx)
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Category GetGoods 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	list, total, err := services.NewGoodsService().GetGoods("category", strconv.Itoa(id), p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Category GetGoods 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Total: total,
		Rows:  list,
	}

	response.Ok(c.Ctx, response.OptionSuccess, res)
	return
}