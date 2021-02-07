package middleware

import (
	"fmt"
	"github.com/juju/ratelimit"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go-mvc/framework/conf"
	"go-mvc/framework/utils/ratelimiter"
	"go-mvc/framework/utils/response"
	"go-mvc/framework/utils/tool"
	"strings"
	"time"

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

/**
令牌桶限流
*/
func RateLimit(fillInterval time.Duration, cap int64) func(ctx context.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(ctx context.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) < 1 {
			response.Error(ctx, iris.StatusBadRequest, response.RateLimiterFailur, nil)
			return
		}
		ctx.Next()
	}
}

// 此方法有问题
func Limit(ctx context.Context) {
	limiter := ratelimiter.TokenBucket{}
	limiter.Set(100, 10)
	fmt.Println("b==",limiter.Allow())
	if !limiter.Allow() {
		response.Error(ctx, iris.StatusBadRequest, response.RateLimiterFailur, nil)
		return
	}
	ctx.Next()
	return
}
