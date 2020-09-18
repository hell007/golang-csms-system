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
	redisClient "go-mvc/framework/cache/redis"
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
		ctx.Application().Logger().Errorf("Product.ProductList参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 查询
	product, err = services.NewGoodsService().GetProduct(id)
	if err != nil {
		ctx.Application().Logger().Errorf("Product.ProductList查询产品详情错误：[%s]", err)
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
		jsonU string
		form   = new(models.ProductForm)
		user = new(member.LoginUser)
		order = new(morder.Order)
		goods []morder.OrderGoods
	)

	// 参数
	if err = ctx.ReadJSON(&form); err != nil {
		ctx.Application().Logger().Errorf("Product.SaveProduct参数错误：[%s]", err)
		response.Error(ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 解密
	keys = encrypt.AESDecrypt(form.Token, conf.GlobalConfig.JWTSalt)
	jsonU, err = redisClient.Get(conf.GlobalConfig.RedisPrefix + keys).Result()
	if err = json.Unmarshal([]byte(jsonU), &user); err != nil {
		ctx.Application().Logger().Error("Product.SaveProduct解密错误：user不存在")
		response.Failur(ctx, response.OptionFailur, nil)
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
				ctx.Application().Logger().Errorf("Product.SaveProduct提交订单错误：[%s]", err)
				response.Failur(ctx, response.OptionFailur, nil)
				return
			}

		case 2: //预购单

	}

	response.Ok(ctx, response.OptionSuccess, nil)
	return
}

