package routes

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/cors"
	"go-mvc/web/apis"
)

// mvc模式
func Configure(b *bootstrap.Bootstrapper) {

	/* 定义路由 */
	main := b.Party("/", cors.Mycors()).AllowMethods(iris.MethodOptions)
	//main.Use(middleware.ServeAPIS)

	// 首屏渲染
	page :=  main.Party("/")
	page.Get("/", func(ctx iris.Context) {
		ctx.ViewData("path", ctx.Path())
		ctx.View("index.html")
	})

	// api模块路由
	api := main.Party("/api")

	// 测试路由------S
	test := api.Party("/test", func(ctx iris.Context){
		ctx.Next()
	})
	{
		test.Get("/v1", apis.V1)
	}
	// 测试路由------E

	// 首页路由
	home :=  api.Party("/home")
	home.Get("/v1", apis.Home)

	// 分类路由
	category := api.Party("/category")
	{
		category.Get("/v1", apis.CategoryList)
		category.Get("/v2", apis.CategoryGoods)
	}

	// 搜索路由
	search := api.Party("/search")
	search.Get("/v1", apis.Search)

	// 产品路由
	product := api.Party("/product")
	{
		product.Get("/v1", apis.ProductList)
		product.Get("/v2", apis.ProductDo)
	}

	// 用户路由
	user := api.Party("/user")
	{
		user.Get("/register", apis.Register)
		user.Get("/captcha", apis.Captcha)
		user.Post("/login", apis.Login)
		user.Get("/logout", apis.Logout)
		user.Get("/userInfo", apis.UserInfo)
		user.Get("/findUser", apis.FindUser)
		user.Get("/city", apis.City)
		user.Get("/userAddress", apis.UserAddress)
		user.Post("/saveAddress", apis.SaveAddress)
		user.Get("/deleteAddress", apis.DeleteAddress)
	}

	// 订单路由
	order := api.Party("/order")
	{
		order.Get("/orderList", apis.OrderList)
		order.Get("/orderItem", apis.OrderItem)
		order.Post("/saveOrder", apis.SaveOrder)
	}
}
