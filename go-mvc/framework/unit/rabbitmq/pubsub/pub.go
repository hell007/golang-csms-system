package main

import (
	"fmt"
	"go-mvc/framework/rabbitmq"
	"strconv"
	"time"
)

func main() {
	r := rabbitmq.NewRabbitMQPubSub("product")

	for i := 1; i <= 20; i++ {
		r.PublishPub("订阅模式：" + strconv.Itoa(i))
		<-time.Tick(time.Second)
		fmt.Println("订阅模式生产消息：" + strconv.Itoa(i))
	}
}