package main

import (
	"time"

	"go-mvc/framework/bootstrap"
	"go-mvc/framework/middleware/identity"
	"go-mvc/framework/utils/files"
	"go-mvc/framework/utils/logs"
	"go-mvc/web/routes"
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

	// 日志设置
	logFile := time.Now().Format("2006-01-02") + ".txt"
	log := logs.Logger.GetConf().Logs
	f, _ := files.CreateFile(log.Output + logFile)
	defer f.Close()
	app.Logger().AddOutput(f)
	app.Logger().SetLevel(log.Level)

	//监听端口
	app.Listen(":7000")
}