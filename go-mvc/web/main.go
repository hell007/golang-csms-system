package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	_ "go-mvc/docs"

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

// swagger middleware for Iris
// swagger embed files

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8085
// @BasePath /vue
func main() {
	// 创建实例
	app := newApp()

	// 静态资源
	app.Favicon(StaticAssets + Favicon)
	app.HandleDir("/assets", StaticAssets)
	app.HandleDir("/uploads", Uploads)

	// logs开启
	logs.Start(conf.Global.Directory + conf.Global.LogsAppPath)

	// cache开启
	redis.Start()

	// swagger配置方法一：其他文档
	config := swagger.Config{
		// 指向swagger init生成文档的路径
		URL:          "http://www.xxx.com/swagger/doc.json",
		DeepLinking:  true,

	}
	app.Get("/swagger/*any", swagger.CustomWrapHandler(&config,swaggerFiles.Handler))
	// swagger配置方法二：默认文档
	// app.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	//监听端口
	app.Listen(":7000")
}
