package main

import (
	"fmt"
	"go-mvc/framework/rabbitmq"
	"strconv"
	"time"
)

func main() {
	r := rabbitmq.NewRabbitMQSimple("test-work")

	for i := 1; i <= 20; i++ {
		r.PublishSimple("Hello Work! " + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}

}
