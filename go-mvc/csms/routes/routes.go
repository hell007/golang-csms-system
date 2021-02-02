package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
	"go-mvc/csms/controllers"
	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/cors"
	"go-mvc/framework/services"
)

// Configure: registers the necessary routes to the app.
// mvc模式
func Configure(b *bootstrap.Bootstrapper) {
	// 定义路由,开启拦截
	main := b.Party("/", cors.Mycors()).AllowMethods(iris.MethodOptions)
	//main.Use(middleware.ServeHTTP)

	//-------智能路由模式测试-------
	// 首页模块
	home := main.Party("/")
	home.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	home.Get("/test", hero.Handler(controllers.HomeIndex()))

	// test resfull api
	test := main.Party("/test")
	{
		test.Get("/getLog", hero.Handler(controllers.TestLog))
	}

	//-------mvc 路由模式------
	// 系统模块
	sys := main.Party("/sys")
	{
		//系统用户子模块
		user := mvc.New(sys.Party("/user"))
		userService := services.NewUserService()
		user.Register(userService)
		user.Handle(new(controllers.UserController))

		//系统角色子模块
		role := mvc.New(sys.Party("/role"))
		roleService := services.NewRoleService()
		role.Register(roleService)
		role.Handle(new(controllers.RoleController))

		//系统casbinrule子模块
		rule := mvc.New(sys.Party("/rule"))
		ruleService := services.NewRuleService()
		rule.Register(ruleService)
		rule.Handle(new(controllers.RuleController))

		//系统casbinrule子模块
		permission := mvc.New(sys.Party("/permission"))
		menuService := services.NewMenuService()
		permission.Register(menuService)
		permission.Handle(new(controllers.PermissionController))

		//系统casbinrule子模块
		access := mvc.New(sys.Party("/access"))
		accessService := services.NewAccessService()
		access.Register(accessService)
		access.Handle(new(controllers.AccessController))
	}

	goods := main.Party("/goods")
	{
		// product
		product := mvc.New(goods.Party("/product"))
		goodsService := services.NewGoodsService()
		product.Register(goodsService)
		product.Handle(new(controllers.ProductController))

		// category
		category := mvc.New(goods.Party("/category"))
		categoryService := services.NewCategoryService()
		category.Register(categoryService)
		category.Handle(new(controllers.CategoryController))

		// sku
		sku := mvc.New(goods.Party("/sku"))
		sku.Handle(new(controllers.SkuController))

		// gallery
		gallery := mvc.New(goods.Party("/gallery"))
		gallery.Handle(new(controllers.GalleryController))
	}

	// order
	order := mvc.New(main.Party("/order"))
	orderService := services.NewOrderService()
	order.Register(orderService)
	order.Handle(new(controllers.OrderController))

	// member
	member := mvc.New(main.Party("/member"))
	memberService := services.NewMemberService()
	member.Register(memberService)
	member.Handle(new(controllers.MemberController))

}
