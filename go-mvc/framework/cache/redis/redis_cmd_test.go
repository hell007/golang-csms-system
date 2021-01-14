package redis

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	models "go-mvc/framework/models/system"
	"go-mvc/framework/services"
)

// 参考文档
// https://github.com/go-redis/redis.git
// https://www.jianshu.com/p/7485a8693a52

// $ go test -v -cover redis_cmd_test.go redis_cmd.go redis_client.go

func TestSet(t *testing.T) {
	// 字符串
	err := Set("test", "redis", 0).Err()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}

	// 储存 struct转为json字符串
	list := services.NewUserService().GetAll()
	jsonData, _ := json.Marshal(list)

	err2 := Set("test2", jsonData, 100*time.Second).Err()
	if err2 != nil {
		fmt.Println("client.Set fail", err2)
	}
}

func TestGet(t *testing.T) {
	// 字符串
	val, err := Get("test").Result()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}
	fmt.Println("val===", val)

	// 获取 json字符串转struct
	var (
		list []models.User
	)

	jsonData, err := Get("test2").Result()
	if err != nil {
		fmt.Println("client.Set fail", err)
	}

	if err2 := json.Unmarshal([]byte(jsonData), &list); err2 != nil {
		fmt.Println("son.Unmarshal fail", err2)
	}

	for _, v := range list {
		fmt.Println("user ===", v)
		fmt.Println("user.username ===", v.Username)
	}
}

func Testmain(t *testing.T) {
	//订阅
	sub := Subscribe("channel")
	//新协程
	go func() {
		for {
			//发布
			pub := Publish("channel", "message")
			if pub.Err() != nil {
				fmt.Println("Publish err", "message")
			} else {
				fmt.Println("Publish msg", "message")
			}
		}
	}()

	//从订阅获取信息，获取一次则程序结束
	msg, err := sub.ReceiveMessage()
	if err != nil {
		fmt.Println("Subscribe err", err)
	} else {
		fmt.Println("Subscribe msg", msg.String())
	}
}
