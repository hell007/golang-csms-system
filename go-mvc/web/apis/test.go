package apis

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/utils/uuids"
	"time"

	redisClient "go-mvc/framework/cache/redis"
	"go-mvc/framework/conf"
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
	//110108012 + 13888888888
	var ordersn string

	ordersn, _ = uuids.SnowId()

	response.Ok(ctx, response.OptionSuccess, ordersn)
	return
}
