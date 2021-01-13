package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Id    string `json:"id"`
	Name     string `json:"name" validate:"required,CustomValidationErrors"` //包含自定义函数
	Phone    string `json:"phone" validate:"required"`
	Passwd   string `json:"passwd" validate:"required,max=20,min=6"`
	Repasswd string `json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd"`
	Code     string `json:"code" validate:"required,len=6"`
}

func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() != "admin"
}

func main() {

	users := &Users{
		Name:   "test",//"admin" //自定义验证
		Phone:  "13837462388",
		Passwd: "123456789",
		Repasswd: "123456789x",
		Code:   "708090",
	}
	validate := validator.New()

	//注册自定义函数
	_ = validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)

	err := validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			return
		}
	}
	return
}
