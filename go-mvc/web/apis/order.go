package apis

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"time"

	"go-mvc/framework/conf"
	members "go-mvc/framework/models/member"
	models "go-mvc/framework/models/order"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/encrypt"
	"go-mvc/framework/utils/page"
	redisClient "go-mvc/framework/cache/redis"
	"go-mvc/framework/utils/response"
)

// list
func OrderList(ctx iris.Context) {
	var (
		err      error
		state    int
		token, keys, jsonU string
		p        *page.Pagination
		res      *page.Result
		list     []models.Orders
		user = new(members.LoginUser)
		total    int64
	)

	// 分页设置
	p, err = page.NewPagination(ctx)
	if err != nil {
		ctx.Application().Logger().Errorf("Order.OrderList 参数：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 解密
	token = ctx.URLParam("token")
	keys = encrypt.AESDecrypt(token, conf.GlobalConfig.JWTSalt)
	jsonU, err = redisClient.Get(conf.GlobalConfig.RedisPrefix + keys).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		ctx.Application().Logger().Errorf("Order.OrderList 解密：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	// 查询
	state, _ = ctx.URLParamInt("state")
	list, total, err = services.NewOrderService().GetOrderListByMid(user.Id, state, p)
	if err != nil {
		ctx.Application().Logger().Errorf("Order.OrderList 查询：[%s]", err)
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

// item
func OrderItem(ctx iris.Context) {
	var (
		err         error
		ordersn     string
		orderDetail = new(models.OrderDetail)
	)

	// 参数处理
	ordersn = ctx.URLParam("id")
	if err != nil {
		ctx.Application().Logger().Errorf("Order.OrderItem 参数：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	orderDetail, err = services.NewOrderService().Get(ordersn)
	if err != nil {
		ctx.Application().Logger().Errorf("Order.OrderItem 查询：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, orderDetail)
	return
}

// save
func SaveOrder(ctx iris.Context) {
	var (
		err    error
		effect int64
		columns []string
		order   = new(models.Order)
	)

	if err = ctx.ReadJSON(&order); err != nil {
		ctx.Application().Logger().Errorf("Order.SaveOrder Json：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	if order.Id > 0 {
		order.ShipTime = time.Now()
		order.OrderState = 4 //已取消
		columns = append(columns, "order_state")
		effect, err =  services.NewOrderService().Update(order, columns)
	}

	if effect < 0 || err != nil {
		ctx.Application().Logger().Errorf("Order.SaveOrder 操作：[%s]", err)
		response.Failur(ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}