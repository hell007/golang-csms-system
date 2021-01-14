package idgen

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/satori/go.uuid"
	"strings"
	"time"

	redisClient "go-mvc/framework/cache/redis"
)

const (
	orderSnKey   = "jie-order-sn"
	orderSnStart = "100000000000000"                // 15位
	orderSnTime  = 365 * 24 * 60 * 60 * time.Second // 1年
)

// uuid 测试
func Generate() {
	u1 := uuid.NewV1()
	u2 := uuid.NewV2(0x00)
	u3 := uuid.NewV3(u1, "uid")
	u4 := uuid.NewV4()
	u5 := uuid.NewV5(u1, "u1")

	fmt.Println("u1==", u1)
	fmt.Println("u2==", u2)
	fmt.Println("u3==", u3)
	fmt.Println("u4==", u4)
	fmt.Println("u5==", u5)
}

// 生成uuid 32位
func GenerateUuid() string {
	var err error
	return strings.ReplaceAll(uuid.Must(uuid.NewV1(), err).String(), "-", "")
}

// 生成订单号
func GenerateOrdersn() int64 {
	// 淘宝18位 14位 + 15/16 + 17/18
	// 利用redis生成订单号策略
	// 失效日期需要永久
	// 开始值自定义

	// 15位
	var (
		err     error
		ordersn int64
	)

	//redisClient.Del(orderSnKey)

	//stringTime  := time.Now().Format("01-02")
	//fmt.Println(stringTime)

	_, err = redisClient.Get(orderSnKey).Result()
	if err != nil {
		redisClient.Set(orderSnKey, orderSnStart, orderSnTime).Err()
		golog.Infof("GenerateOrdersn Get %s", err)
	}

	ordersn, err = redisClient.Incr(orderSnKey).Result()
	if err != nil {
		ordersn = 0
		golog.Infof("GenerateOrdersn Incr %s", err)
	}

	redisClient.Set(orderSnKey, ordersn, orderSnTime).Err()
	return ordersn
}
