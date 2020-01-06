package main

import (
	"time"

	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/identity"
	"go-mvc/framework/utils/files"
	"go-mvc/web/routes"
	"go-mvc/framework/conf"
)

const (
	StaticAssets = "./assets/"
	Favicon      = "favicon.ico"
	Uploads = "../uploads/"
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

	// 日志设置
	logFile := time.Now().Format(conf.GlobalConfig.TimeformatShort) + ".log"
	f, _ := files.CreateFile(conf.GlobalConfig.LogsOutput + logFile)
	defer f.Close()
	app.Logger().AddOutput(f)
	app.Logger().SetLevel(conf.GlobalConfig.LogsLevel)

	//监听端口
	app.Listen(":7000")
}