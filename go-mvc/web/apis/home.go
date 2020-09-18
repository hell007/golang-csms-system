package apis

import (
	"github.com/kataras/iris/v12"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

/**
首页
 */
func Home(ctx iris.Context) {
	var (
		err       error
		p         *page.Pagination
		hotList   []models.GoodsDetail
		firstList []models.GoodsDetail
		maps      = make(map[string]interface{}, 0)
	)

	// 分页设置
	p, err = page.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("Home.Home参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 热销
	hotList, _, err = services.NewGoodsService().GetGoods("hot", "1", p)
	if err != nil {
		ctx.Application().Logger().Errorf("Home.Home查询热销商品错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}
	maps["hot"] = hotList

	// 主推
	firstList, _, err = services.NewGoodsService().GetGoods("first", "1", p)
	if err != nil {
		ctx.Application().Logger().Errorf("Home.Home查询热销商品错误：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}
	maps["first"] = firstList

	response.Ok(ctx, response.OptionSuccess, maps)
	return
}
