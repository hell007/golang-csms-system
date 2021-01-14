package yaag

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/bootstrap"
)

/**
下载YAAG中间件
go get github.com/betacraft/yaag/...

导入依赖包
import github.com/betacraft/yaag/yaag Import github.com/betacraft/yaag/irisyaag

初始化yaag
yaag.Init(&yaag.Config(On: true, DocTile: "Iris", DocPath: "apidoc.html"))

注册yaag中间件
app.Use(irisyaag.New()) irisyaag记录响应主体并向apidoc提供所有必要的信息
*/

type myXML struct {
	Result string `xml:"result"`
}

func Configure(b *bootstrap.Bootstrapper) {

	//初始化中间件
	yaag.Init(&yaag.Config{
		On:       true, //是否开启自动生成API文档功能
		DocTitle: "Iris",
		DocPath:  "apidoc.html", //生成API文档名称存放路径
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})

	//注册中间件
	b.Use(irisyaag.New())
	b.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"result": "Hello World!"})
	})
	b.Get("/plain", func(ctx iris.Context) {
		ctx.Text("Hello World!")
	})
	b.Get("/xml", func(ctx iris.Context) {
		ctx.XML(myXML{Result: "Hello World!"})
	})
	b.Get("/complex", func(ctx iris.Context) {
		value := ctx.URLParam("key")
		ctx.JSON(iris.Map{"value": value})
	})
}
