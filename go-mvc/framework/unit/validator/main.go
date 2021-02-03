package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

/**
原理:
将验证规则写在struct对字段tag里，再通过反射（reflect）获取struct的tag，实现数据验证。

验证规则:
required ：必填
email：验证字符串是email格式；例：“email”
url：这将验证字符串值包含有效的网址;例：“url”
max：字符串最大长度；例：“max=20”
min:字符串最小长度；例：“min=6”
excludesall:不能包含特殊字符；例：“excludesall=0x2C”//注意这里用十六进制表示。
len：字符长度必须等于n，或者数组、切片、map的len值为n，即包含的项目数；例：“len=6”
eq：数字等于n，或者或者数组、切片、map的len值为n，即包含的项目数；例：“eq=6”
ne：数字不等于n，或者或者数组、切片、map的len值不等于为n，即包含的项目数不为n，其和eq相反；例：“ne=6”
gt：数字大于n，或者或者数组、切片、map的len值大于n，即包含的项目数大于n；例：“gt=6”
gte：数字大于或等于n，或者或者数组、切片、map的len值大于或等于n，即包含的项目数大于或等于n；例：“gte=6”
lt：数字小于n，或者或者数组、切片、map的len值小于n，即包含的项目数小于n；例：“lt=6”
lte：数字小于或等于n，或者或者数组、切片、map的len值小于或等于n，即包含的项目数小于或等于n；例：“lte=6”

跨字段验证:如想实现比较输入密码和确认密码是否一致等类似场景
eqfield=Field: 必须等于 Field 的值；
nefield=Field: 必须不等于 Field 的值；
gtfield=Field: 必须大于 Field 的值；
gtefield=Field: 必须大于等于 Field 的值；
ltfield=Field: 必须小于 Field 的值；
ltefield=Field: 必须小于等于 Field 的值；
eqcsfield=Other.Field: 必须等于 struct Other 中 Field 的值；
necsfield=Other.Field: 必须不等于 struct Other 中 Field 的值；
gtcsfield=Other.Field: 必须大于 struct Other 中 Field 的值；
gtecsfield=Other.Field: 必须大于等于 struct Other 中 Field 的值；
ltcsfield=Other.Field: 必须小于 struct Other 中 Field 的值；
ltecsfield=Other.Field: 必须小于等于 struct Other 中 Field 的值
*/

type Users struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required,CustomValidationErrors" label:"用户名"` //包含自定义函数
	Phone    string `json:"phone" validate:"required,len=11" label:"手机号"`
	Passwd   string `json:"passwd" validate:"required,max=20,min=6" label:"密码"`
	Repasswd string `json:"repasswd" validate:"required,max=20,min=6,eqfield=Passwd" label:"确认密码"`
	Code     string `json:"code" validate:"required,len=6" label:"验证码"`
}

//验证函数
func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() != "admin"
}

//标记函数
func CustomTranslationError1(ut ut.Translator) error {
	return ut.Add("CustomValidationErrors", "{0}为{1}的用户已存在", true)
}

//标记函数
func CustomTranslationError2(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("CustomValidationErrors", fe.Field(), fmt.Sprintf("%v", fe.Value()))
	return t
}

func main() {

	users := &Users{
		Name:     "admin", //"admin" //自定义验证
		Phone:    "13837462388",
		Passwd:   "123456789",
		Repasswd: "123456789",
		Code:     "708090",
	}

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	//注册自定义函数
	_ = validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)

	//根据提供的标记注册翻译
	_ = validate.RegisterTranslation("CustomValidationErrors", trans, CustomTranslationError1, CustomTranslationError2)

	//验证器注册翻译器
	err := translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}

	err2 := validate.Struct(users)
	if err2 != nil {
		for _, err3 := range err2.(validator.ValidationErrors) {
			//Key: 'Users.Name' Error:Field validation for 'Name' failed on the 'required' tag
			//用户名为必填字段
			fmt.Println(err3.Translate(trans))
			return
		}
	}
	return
}
