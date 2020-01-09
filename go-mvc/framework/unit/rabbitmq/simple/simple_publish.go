package main

import (
	"go-mvc/framework/rabbitmq"
	"log"
)

func main() {
	r := rabbitmq.NewRabbitMQSimple("test-simple")
	r.PublishSimple("Hello Simple！！！")
	log.Println("发送成功！")
}
