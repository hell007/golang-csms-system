package routes

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/cors"
	"go-mvc/web/apis"
)

func Configure(b *bootstrap.Bootstrapper) {

	/* 定义路由 */
	main := b.Party("/", cors.Mycors()).AllowMethods(iris.MethodOptions)
	//main.Use(middleware.ServeAPIS)

	// 首页
	page := main.Party("/")
	page.Get("/", func(ctx iris.Context) {
		ctx.ViewData("path", ctx.Path())
		ctx.View("index.html")
	})

	// api模块路由
	api := main.Party("/api")

	// 测试路由

	test := api.Party("/test", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		test.Get("/v3", apis.V3)
	}

	// 首页路由
	home := api.Party("/home")
	home.Get("/list", apis.Home)

	// 分类路由
	category := api.Party("/category")
	{
		category.Get("/type", apis.CategoryList)
		category.Get("/list", apis.CategoryGoods)
	}

	// 搜索路由
	search := api.Party("/search")
	search.Get("/list", apis.Search)

	// 产品路由
	product := api.Party("/product")
	{
		product.Get("/detail", apis.ProductDetail)
		product.Get("/save", apis.SaveProduct)
	}

	// 用户路由
	user := api.Party("/user")
	{
		user.Post("/register", apis.Register)
		user.Get("/captcha", apis.Captcha)
		user.Post("/login", apis.Login)
		user.Get("/logout", apis.Logout)
		user.Get("/profile", apis.Profile)
		user.Get("/findUser", apis.FindUser)
		user.Get("/city", apis.City)
		user.Get("/userAddress", apis.UserAddress)
		user.Post("/saveAddress", apis.SaveAddress)
		user.Get("/deleteAddress", apis.DeleteAddress)
	}

	// 订单路由
	order := api.Party("/order")
	{
		order.Get("/list", apis.OrderList)
		order.Get("/detail", apis.OrderDetail)
		order.Post("/save", apis.SaveOrder)
	}
}
