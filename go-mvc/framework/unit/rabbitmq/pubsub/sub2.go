package main

import "go-mvc/framework/rabbitmq"

func main() {
	r := rabbitmq.NewRabbitMQPubSub("product")
	r.RecieveSub()
}