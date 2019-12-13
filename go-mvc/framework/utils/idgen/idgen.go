package idgen

import (
	//"github.com/satori/go.uuid"
	//"strings"
	"time"

	redisClient "go-mvc/framework/utils/redis"
)

const ordersnNo = "jie-order-sn"
var  ordersnTime = 365 * 24 * 60 * 60 * time.Second

// 生成uuid 32位
func GenerateUuid() string {
	return "10000000"//strings.ReplaceAll(uuid.Must(uuid.NewV1()).String(), "-", "")
}

// 生成订单号
func GenerateOrdersn() int64{
	// 淘宝18位 14位 + 15/16 + 17/18
	// 利用redis生成订单号策略
	// 失效日期需要永久
	// 开始值自定义

	// 15位
	ordersnStart := "100000000000000"
	_, err  := redisClient.Get(ordersnNo).Result()
	if err != nil {
		redisClient.Set(ordersnNo, ordersnStart, ordersnTime).Err()
	}

	ordersn,_ := redisClient.Incr(ordersnNo).Result()
	redisClient.Set(ordersnNo, ordersn, ordersnTime).Err()
	return ordersn
}