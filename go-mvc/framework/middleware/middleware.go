package middleware

import (
	"strings"

	"github.com/kataras/iris/v12/context"

	"go-mvc/framework/conf"
	"go-mvc/framework/middleware/casbin"
	"go-mvc/framework/middleware/jwt"
)

// 后台系统
func ServeHTTP(ctx context.Context) {
	path := ctx.Path()

	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/assets") {
		ctx.Next()
		return
	}

	// jwt token拦截
	if !jwt.Serve(ctx) {
		return
	}

	// 系统菜单不进行权限拦截
	if !strings.Contains(path, "/permission") {
		// casbin权限拦截
		ok := casbin.CheckPermissions(ctx)

		if !ok {
			return
		}
	}

	// Pass to real API
	ctx.Next()
}

// 前台
func ServeAPIS(ctx context.Context) {
	path := ctx.Path()

	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) || strings.Contains(path, "/assets") {
		ctx.Next()
		return
	}

	// jwt token拦截
	if !jwt.Serve(ctx) {
		return
	}

	// Pass to real API
	ctx.Next()
}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
*/
func checkURL(reqPath string) bool {
	for _, v := range conf.GlobalConfig.AuthIgnores {
		if reqPath == v {
			return true
		}
	}
	return false
}
