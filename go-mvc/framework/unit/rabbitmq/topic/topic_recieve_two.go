package main

import "go-mvc/framework/rabbitmq"

func main() {
	r := rabbitmq.NewRabbitMQTopic("exchange-topic", "test.*.two")
	r.RecieveTopic()
}