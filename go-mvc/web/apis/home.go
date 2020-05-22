package apis

import (
	"github.com/kataras/iris/v12"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type Home struct {
	Ctx iris.Context
}

func (c *Home) GetV1() {
	var (
		err       error
		p         *page.Pagination
		hotList   []models.GoodsDetail
		firstList []models.GoodsDetail
		maps      = make(map[string]interface{}, 0)
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Home GetV1 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 热销
	hotList, _, err = services.NewGoodsService().GetGoods("hot", "1", p)
	maps["hot"] = hotList

	// 主推
	firstList, _, err = services.NewGoodsService().GetGoods("first", "1", p)
	maps["first"] = firstList

	response.Ok(c.Ctx, response.OptionSuccess, maps)
	return
}
