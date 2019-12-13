package conf

import (
	"runtime"
	"time"
)

/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
年　 06,2006
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST
您看出规律了么！哦是的，你发现了，这里面没有一个是重复的，所有的值表示都唯一对应一个时间部分。
并且涵盖了很多格式组合。
*/

// 时间格式化字符串
const (
	SysTimeform      string = "2006-01-02 15:04:05"
	SysTimeformShort string = "2006-01-02"
)

// jwt
const (
	JWTSecret  string = "SECRET_" //加盐
	JWTSalt string = "ARRSWdczx13213EDDSWQ!!@W"
	JWTTimeout int    = 3600         //单位秒：设置1小时
)

// redis
const (
	RedisPrefix string = "JIE_REDIES_KEY_"
)
var RedisExprire = 1 * 30 * 60 * time.Second //单位秒：设置1小时

// auth
var AuthIgnores = []string{
	"/",
	"/sys/user/login",
	"/sys/user/loginout",
	"/sys/rule/menus",
	"/sys/test/upload",
	"/api/test/index",
}

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// upload
const (
	BaseUrl      string = "http://127.0.0.1:9000"
	UploadBase   string = BaseUrl + "/uploads"
	UploadSall   string = "100x100"
	UploadMedium string = "360x360"
	UploadSource string = "800x800"
	UploadNormal string = ""
	UploadGoods  string = "/goods/"
	UploadArticle  string = "/article/"
	UploadAds  string = "/ads/"
	UploadGoodsEditor  string = "/images/goods/"
	UploadArticleEditor  string = "/images/article/"
)
var (
	UploadFile string
	ConfigPath string
)
// 返回上传文件路径
func GetUploadFile() string {
	sysType := runtime.GOOS
	if sysType == "darwin" {
		UploadFile = "/Users/wzh/Development/git/go/golang-csms-system/golang-mvc/uploads"
	} else if sysType == "windows" {
		UploadFile = "D:/Dev/cygwin/work/golang/golang-csms-system/golang-mvc/uploads"
	}
	return UploadFile
}

func GetConfigPath() string {
	sysType := runtime.GOOS
	if sysType == "darwin" {
		ConfigPath = "/Users/wzh/Development/git/go/golang-csms-system/golang-mvc/framework/conf/"
	} else if sysType == "windows" {
		ConfigPath = "D:/Dev/cygwin/work/golang/golang-csms-system/golang-mvc/framework/conf/"
	}
	return ConfigPath
}

