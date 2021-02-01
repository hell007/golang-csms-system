package main

import (
	"go-mvc/framework/bootstrap"
	"go-mvc/framework/cache/redis"
	"go-mvc/framework/conf"
	"go-mvc/framework/logs"
	"go-mvc/framework/middleware/identity"
	"go-mvc/web/routes"
)

const (
	StaticAssets = "./assets/"
	Favicon      = "favicon.ico"
	Uploads      = "../uploads/"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("web-system", "jie")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	// 创建实例
	app := newApp()

	// 静态资源
	app.Favicon(StaticAssets + Favicon)
	app.HandleDir("/assets", StaticAssets)
	app.HandleDir("/uploads", Uploads)

	// logs开启
	logs.Start(conf.GlobalConfig.LogsAppPath)

	// cache开启
	redis.Start()

	//监听端口
	app.Listen(":7000")
}
