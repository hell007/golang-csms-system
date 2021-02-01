package apis

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	"go-mvc/framework/utils/tool"
	"strconv"
	"time"

	"go-mvc/framework/conf"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/models/member"
	morder "go-mvc/framework/models/order"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/idgen"
	"go-mvc/framework/utils/response"
)

/**
产品详情
*/
func ProductDetail(ctx iris.Context) {
	var (
		err     error
		id      int
		product = new(models.Product)
	)

	// 参数
	id, err = ctx.URLParamInt("id")
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	product, err = services.NewGoodsService().GetProduct(id)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询产品详情错误")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, product)
	return
}

func SaveProduct(ctx iris.Context) {
	var (
		err    error
		effect int64
		form   = new(models.ProductForm)
		user   = new(member.LoginUser)
		order  = new(morder.Order)
		goods  []morder.OrderGoods
	)

	// 参数
	if err = ctx.ReadJSON(&form); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.GlobalConfig.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// 数据
	goods = form.Goods
	order = form.Order
	order.Mid = user.Id

	switch form.Code { //区分订单还是购物车
	case 1: //订单
		order.OrderState = 0
		order.Ordersn = strconv.FormatInt(idgen.GenerateOrdersn(), 10)
		order.CreateTime = time.Now()
		order.TotalPrice = "0"
		order.ShipPrice = "0"
		order.OrderPrice = "0"
		order.PayPrice = "0"
		order.ReturnPrice = "0"

		effect, err = services.NewOrderService().Create(order, goods)
		if effect <= 0 && err != nil {
			logs.GetLogger().Error(logs.D{"err": err}, "提交订单错误")
			response.Failur(ctx, response.OptionFailur, nil)
			return
		}

	case 2: //预购单

	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}
