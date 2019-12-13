package apis

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"

	"go-mvc/framework/conf"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/models/member"
	morder "go-mvc/framework/models/order"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/encrypt"
	"go-mvc/framework/utils/idgen"
	redisClient "go-mvc/framework/utils/redis"
	"go-mvc/framework/utils/response"
)

type Product struct {
	Ctx iris.Context
}

func (c *Product) GetV1() {
	var (
		err     error
		id      int
		product = new(models.Product)
	)

	// 参数
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Product GetV1 参数：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	product, err = services.NewGoodsService().GetProduct(id)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Product GetV1 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, product)
	return
}

func (c *Product) PostDo() {
	var (
		err    error
		effect int64
		jsonU string
		form   = new(models.ProductForm)
		user = new(member.LoginUser)
		order = new(morder.Order)
		goods []morder.OrderGoods
	)

	// 参数
	if err = c.Ctx.ReadJSON(&form); err != nil {
		c.Ctx.Application().Logger().Errorf("Product PostDo Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 解密
	keys = encrypt.AESDecrypt(form.Token, conf.JWTSalt)
	jsonU, err = redisClient.Get(conf.RedisPrefix + keys).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		c.Ctx.Application().Logger().Error("Product PostDo user不存在")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 数据
	goods = form.Goods
	order = form.Order
	order.Mid = user.Id

	switch form.Code { //区分订单还是购物车
		case 1: //订单
			order.OrderState = 0
			order.Ordersn = strconv.FormatInt(idgen.GenerateOrdersn(),10)
			order.CreateTime = time.Now()
			order.TotalPrice = "0"
			order.ShipPrice = "0"
			order.OrderPrice = "0"
			order.PayPrice = "0"
			order.ReturnPrice = "0"

			effect, err = services.NewOrderService().Create(order, goods)
			if effect <= 0 && err != nil {
				c.Ctx.Application().Logger().Errorf("Product PostDo 添加：[%s]", err)
				response.Failur(c.Ctx, response.OptionFailur, nil)
				return
			}

		case 2: //预购单

	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

