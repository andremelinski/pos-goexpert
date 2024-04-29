package main

import (
	"fmt"
	"sync"

	rabbitmq "github.com/andremelinski/pos-goexpert/10-eventos/pkg/rabbitmq/direct"
	"github.com/rabbitmq/amqp091-go"
)

func main(){
	wg := sync.WaitGroup{}
	msgs := make(chan amqp091.Delivery)
	// msgs2 := make(chan amqp091.Delivery)
	// msgs3 := make(chan amqp091.Delivery)
	// chanArr := []chan amqp091.Delivery{msgs, msgs2, msgs3}
	queueNameArr := []string{"queue", "green-q", "blue-q"}
	length := len(queueNameArr)

	chn, err := rabbitmq.OpenChannel()
	if err != nil{
		panic(err)
	}
	wg.Add(length)
	for i := 0; i < length; i++ {
		go rabbitmq.Consume(chn, msgs, queueNameArr[i], &wg)
	}

	for v := range msgs {
		fmt.Println(string(v.Body))
		v.Ack(false)
	}

	wg.Wait()
	println("aqui")
}