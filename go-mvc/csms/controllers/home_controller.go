package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	//"../../framework/models"
	//"../../framework/services"
)

//type HomeController struct {
//	Ctx iris.Context
//	//Service services.UserService
//}

func HomeIndex() mvc.Result {
	return mvc.View{
		Name: "home/index.html",
		Data: iris.Map{
			"Title": "首页",
		},
		Layout: "layout.html",
	}
}
