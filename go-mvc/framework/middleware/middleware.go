package middleware

import (
	"github.com/juju/ratelimit"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go-mvc/framework/conf"
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
令牌桶算法（Token Bucket）是网络流量整形（Traffic Shaping）和速率限制（Rate Limiting）中最常使用的一种算法。典型情况下，令牌桶算法用来控制发送到网络上的数据的数目，并允许突发数据的发送。

我们有一个固定的桶，桶里存放着令牌（token）。一开始桶是空的，系统按固定的时间（rate）往桶里添加令牌，直到桶里的令牌数满，多余的请求会被丢弃。当请求来的时候，从桶里移除一个令牌，如果桶是空的则拒绝请求或者阻塞。

令牌桶有以下特点：
令牌按固定的速率被放入令牌桶中
桶中最多存放 B 个令牌，当桶满时，新添加的令牌被丢弃或拒绝
如果桶中的令牌不足 N 个，则不会删除令牌，且请求将被限流（丢弃或阻塞等待）
令牌桶限制的是平均流入速率（允许突发请求，只要有令牌就可以处理，支持一次拿3个令牌，4个令牌...），并允许一定程度突发流量。

扩展：
如果我们要针对每一个公众号做单独限流，那么其实也很简单：
在真正处理数据之前判断用来区分公众号的key是否存在在这个map中。
如果没有那么初始化一个令牌桶并放入map；
如果已存在则直接使用即可。
需要注意的是因为对全局map做操作，故而要在读写之前做安全锁：
var tokenBuckets map[string]*ratelimit.Bucket
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
