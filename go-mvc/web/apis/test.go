package apis

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"time"

	"go-mvc/framework/conf"
	"go-mvc/framework/utils/idgen"
	redisClient "go-mvc/framework/cache/redis"
	"go-mvc/framework/utils/response"
)

type Test struct {
	Ctx iris.Context
}

var keys = conf.JWTSecret + time.Now().Format("20060102-15:04:05") //24

func (c *Test) GetV1() {
	err := redisClient.Set("test", "测试", 1 * time.Second).Err()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}
}

func (c *Test) GetV2() {
	val, err := redisClient.Get("test").Result()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}
	fmt.Println("val===", val)
}

// test
func (c *Test) GetTest() {
	var (
		err error
	)

	// 淘宝18位 14位 + 15/16 + 17/18
	// 利用redis生成订单号策略
	// 失效日期需要永久
	// 开始值自定义

	// 15位
	idgen.GenerateOrdersn()

	fmt.Println("====",)

	response.Ok(c.Ctx, response.OptionSuccess, err)
	return
}



