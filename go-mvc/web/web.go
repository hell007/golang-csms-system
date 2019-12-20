package main

import (
	"time"

	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/identity"
	"go-mvc/framework/utils/files"
	"go-mvc/framework/utils/logs"
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

	conf.Init()

	// 静态资源
	app.Favicon(StaticAssets + Favicon)
	app.HandleDir("/assets", StaticAssets)
	app.HandleDir("/uploads", Uploads)

	// 日志设置
	logFile := time.Now().Format("2006-01-02") + ".log"
	log := logs.Logger.GetConf().Logs
	f, _ := files.CreateFile(log.Output + logFile)
	defer f.Close()
	app.Logger().AddOutput(f)
	app.Logger().SetLevel(log.Level)

	//监听端口
	app.Listen(":7000")
}