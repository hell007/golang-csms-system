package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/cors"
	"go-mvc/web/apis"
)

// mvc模式
func Configure(b *bootstrap.Bootstrapper) {

	/* 定义路由 */
	main := b.Party("/", cors.Mycors()).AllowMethods(iris.MethodOptions)
	//main.Use(middleware.ServeAPIS)

	// 首页模块
	page :=  main.Party("/")
	page.Get("/", func(ctx iris.Context) {
		ctx.ViewData("path", ctx.Path())
		ctx.View("index.html")
	})

	// api模块
	api := main.Party("/api")
	{
		test := mvc.New(api.Party("/test"))
		test.Handle(new(apis.Test))

		//api子模块
		home := mvc.New(api.Party("/home"))
		home.Handle(new(apis.Home))

		category := mvc.New(api.Party("/category"))
		category.Handle(new(apis.Category))

		search := mvc.New(api.Party("/search"))
		search.Handle(new(apis.Search))

		product := mvc.New(api.Party("/product"))
		product.Handle(new(apis.Product))

		user := mvc.New(api.Party("/user"))
		user.Handle(new(apis.User))

		order := mvc.New(api.Party("/order"))
		order.Handle(new(apis.Order))
	}
}
