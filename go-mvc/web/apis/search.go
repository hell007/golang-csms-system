package apis

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"

	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

/**
搜索页
*/
func Search(ctx iris.Context) {
	var (
		err   error
		val   string
		p     *page.Pagination
		res   *page.Result
		list  []models.GoodsDetail
		total int64
	)

	// 分页设置
	val = ctx.URLParam("key")
	p, err = page.NewPagination(ctx)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 热销
	list, total, err = services.NewGoodsService().GetGoods("key", val, p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "搜索商品错误")
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
