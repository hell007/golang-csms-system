package apis

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"time"

	redisClient "go-mvc/framework/cache/redis"
	"go-mvc/framework/conf"
	"go-mvc/framework/utils/idgen"
	"go-mvc/framework/utils/response"
)

var keys = conf.Global.JWTSecret + time.Now().Format(conf.Global.Timeformat) //24

func Test(ctx iris.Context) {
	response.Ok(ctx, response.OptionSuccess, true)
	return
}

func V1(ctx iris.Context) {
	err := redisClient.Set("test", "测试", 1*time.Second).Err()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}
}

func V2(ctx iris.Context) {
	val, err := redisClient.Get("test").Result()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}
	fmt.Println("val===", val)
}

func V3(ctx iris.Context) {
	var (
		err error
	)

	// 淘宝18位 14位 + 15/16 + 17/18
	// 利用redis生成订单号策略
	// 失效日期需要永久
	// 开始值自定义

	// 15位
	idgen.GenerateOrdersn()

	fmt.Println("====", )

	response.Ok(ctx, response.OptionSuccess, err)
	return
}
