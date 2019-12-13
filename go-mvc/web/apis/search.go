package apis

import (
	"github.com/kataras/iris/v12"

	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type Search struct {
	Ctx iris.Context
}

func (c *Search) GetV1() {
	var (
		err   error
		val   string
		p     *page.Pagination
		res   *page.Result
		list  []models.GoodsDetail
		total int64
	)

	// 分页设置
	val = c.Ctx.URLParam("key")
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Search GetV1 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 热销
	list, total, err = services.NewGoodsService().GetGoods("key", val, p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Search GetV1 查询：[%s]", err)
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
