package main

import "go-mvc/framework/rabbitmq"

func main() {
	r := rabbitmq.NewRabbitMQRouting("test-exchange", "test-routingKey-1")
	r.RecieveRouting()
}