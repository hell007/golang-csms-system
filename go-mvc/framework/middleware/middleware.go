package middleware

import (
	"github.com/kataras/iris/v12/context"
	"go-mvc/framework/conf"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/tool"
	"strings"

	"go-mvc/framework/middleware/casbin"
	"go-mvc/framework/middleware/jwt"
)

// 后台系统拦截
func ServeHTTP(ctx context.Context) {
	path := ctx.Path()

	// 静态资源、login接口、首页等...不需要验证
	if checkURL(path, conf.Global.AuthIgnoresWeb) {
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

	ctx.Next()
	return
}

// api路由拦截器
func ServeAPIS(ctx context.Context) {
	path := ctx.Path()

	// 静态资源、login接口、首页等...不需要验证
	if checkURL(path, conf.Global.AuthIgnoresApp) {
		ctx.Next()
		return
	}

	// token验证
	if VerifyUrl(path, conf.Global.AuthVerifyApp) {

		token := ctx.GetHeader(conf.Global.AuthToken)
		if tool.IsEmpty(token) {
			response.Unauthorized(ctx, response.TokenParseFailurAndEmpty, nil)
			return
		}

		_, err := tool.GetUserByToken(token)
		if err != nil {
			response.Unauthorized(ctx, response.TokenParseFailurAndInvalid, nil)
			return
		}
	}

	ctx.Next()
	return
}
