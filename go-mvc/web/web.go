package main

import (
	"io"
	"os"
	"time"

	"go-mvc/framework/bootstrap"
	"go-mvc/framework/conf"
	"go-mvc/framework/middleware/identity"
	"go-mvc/framework/utils/files"
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

	// 日志设置 同时写文件日志与控制台日志
	logFile := time.Now().Format(conf.GlobalConfig.TimeformatShort) + ".log"
	f, _ := files.CreateFile(conf.GlobalConfig.LogsOutput + logFile)
	defer f.Close()
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	app.Logger().SetLevel(conf.GlobalConfig.LogsLevel)

	//监听端口
	app.Listen(":7000")
}
