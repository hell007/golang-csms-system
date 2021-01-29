package middleware

import "strings"

/**
url校验
true:则跳过不需验证，如登录接口等...
false:需要进一步验证,比如token验证
*/
func checkURL(p string, list []string) bool {
	for _, v := range list {
		if p == v {
			return true
		}
	}
	return false
}

/**
token url 验证
Auth.verify 需要验证，排除 “/user/login”
*/
func VerifyUrl(p string, list []string) bool {
	for _, v := range list {
		if strings.Contains(p, v) {
			return true
		}
	}
	return false
}
