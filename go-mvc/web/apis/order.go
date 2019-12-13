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
	redisClient "go-mvc/framework/utils/redis"
	"go-mvc/framework/utils/response"
)

type Order struct {
	Ctx iris.Context
}

// list
func (c *Order) GetV1() {
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
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Order GetV1 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 解密
	token = c.Ctx.URLParam("token")
	keys = encrypt.AESDecrypt(token, conf.JWTSalt)
	jsonU, err = redisClient.Get(conf.RedisPrefix + keys).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		c.Ctx.Application().Logger().Errorf("Order GetV1 解密：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 查询
	state, _ = c.Ctx.URLParamInt("state")
	list, total, err = services.NewOrderService().GetOrderListByMid(user.Id, state, p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Order GetV1 查询：[%s]", err)
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

// item
func (c *Order) GetItem() {
	var (
		err         error
		ordersn     string
		orderDetail = new(models.OrderDetail)
	)

	// 参数处理
	ordersn = c.Ctx.URLParam("id")
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Order GetItem 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	orderDetail, err = services.NewOrderService().Get(ordersn)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Order GetItem 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, orderDetail)
	return
}

// save
func (c *Order) PostSave() {
	var (
		err    error
		effect int64
		columns []string
		order   = new(models.Order)
	)

	if err = c.Ctx.ReadJSON(&order); err != nil {
		c.Ctx.Application().Logger().Errorf("Order PostSave Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	if order.Id > 0 {
		order.ShipTime = time.Now()
		order.OrderState = 4 //已取消
		columns = append(columns, "order_state")
		effect, err =  services.NewOrderService().Update(order, columns)
	}

	if effect < 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Order PostSave 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}