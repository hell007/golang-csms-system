package apis

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	"go-mvc/framework/utils/tool"
	"time"

	"go-mvc/framework/conf"
	members "go-mvc/framework/models/member"
	models "go-mvc/framework/models/order"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

/**
订单列表
*/
func OrderList(ctx iris.Context) {
	var (
		err   error
		state int
		p     *page.Pagination
		res   *page.Result
		list  []models.Orders
		user  = new(members.LoginUser)
		total int64
	)

	// 分页设置
	p, err = page.NewPagination(ctx)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	//通过token获取redis保存的用户
	token := ctx.GetHeader(conf.Global.AuthToken)
	user, _ = tool.GetUserByToken(token)

	// 查询
	state, _ = ctx.URLParamInt("state")
	list, total, err = services.NewOrderService().GetOrderListByMid(user.Id, state, p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询用户订单错误")
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

/**
订单详情
*/
func OrderDetail(ctx iris.Context) {
	var (
		err         error
		ordersn     string
		orderDetail = new(models.OrderDetail)
	)

	// 参数处理
	ordersn = ctx.URLParam("id")
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	orderDetail, err = services.NewOrderService().Get(ordersn)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询订单详情错误")
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, orderDetail)
	return
}

/**
订单保存
*/
func SaveOrder(ctx iris.Context) {
	var (
		err     error
		effect  int64
		columns []string
		order   = new(models.Order)
	)

	if err = ctx.ReadJSON(&order); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	if order.Id > 0 {
		order.ShipTime = time.Now()
		order.OrderState = 4 //已取消
		columns = append(columns, "order_state")
		effect, err = services.NewOrderService().Update(order, columns)
		if effect < 0 || err != nil {
			logs.GetLogger().Error(logs.D{"err": err}, "更新订单状态错误")
			response.Failur(ctx, response.OptionFailur, nil)
			return
		}
	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}
