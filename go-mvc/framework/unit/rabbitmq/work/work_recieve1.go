package main

import "go-mvc/framework/rabbitmq"

func main() {
	r := rabbitmq.NewRabbitMQSimple("test-work")
	r.ConsumerSimple()
}
