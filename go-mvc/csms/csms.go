/**
 * name: csms main
 * author: jie
 * note: csms入口
 */

package main

import (
	_ "go-mvc/csms/inits"
	"go-mvc/csms/routes"
	"go-mvc/framework/bootstrap"
	"go-mvc/framework/conf"
	"go-mvc/framework/logs"
	"go-mvc/framework/middleware/identity"
)

const (
	StaticAssets = "./assets/"
	Favicon      = "favicon.ico"
	Uploads      = "../uploads/"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("csms-system", "jie")
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

	// 日志设置
	logs.Start(conf.GlobalConfig.LogsWebPath)

	//监听端口
	app.Listen(":9000")
}
