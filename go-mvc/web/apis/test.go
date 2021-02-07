package apis

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-mvc/framework/services"
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
	member, _, _ := services.NewMemberService().GetMemberByMobile("138888888888")
	response.Ok(ctx, response.OptionSuccess, member)
	return
}
