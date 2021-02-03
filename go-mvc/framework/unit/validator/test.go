package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required,max=20,min=2" label:"用户名"` //包含自定义函数
	Phone    string `json:"phone" validate:"required,len=11" label:"手机号"`
	Passwd   string `json:"passwd" validate:"required,max=20,min=6" label:"密码"`
	Repasswd string `json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd" label:"确认密码"`
	Code     string `json:"code" validate:"required,len=6" label:"验证码"`
}

type validationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}

func WrapValidationErrors(errs validator.ValidationErrors) []validationError {
	validationErrors := make([]validationError, 0, len(errs))
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, validationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}

	return validationErrors
}

func main() {

	user := &User{
		Name:     "",
		Phone:    "13837462388",
		Passwd:   "123456789",
		Repasswd: "123456789",
		Code:     "708090",
	}

	validate := validator.New()
	err := validate.Struct(user)

	if errs, ok := err.(validator.ValidationErrors); ok {
		validationErrors := WrapValidationErrors(errs)
		fmt.Println("===>", validationErrors)
	}
}
