package controllers

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	//"../../framework/models"
	//"../../framework/services"
)

//type HomeController struct {
//	Ctx iris.Context
//	//Service services.UserService
//}

func HomeList() mvc.Result {
	//datalist := c.Service.GetAll()
	// set the model and render the view template.
	return mvc.View{
		Name: "home/index.html",
		Data: iris.Map{
			"Title": "首页",
		},
		Layout: "layout.html",
	}
}

func HomeTest(ctx iris.Context) {
	fmt.Println("home test")
	golog.Println("fuck fuck fuck")
	ctx.Application().Logger().Error("求你可以写入日志了！")
}
