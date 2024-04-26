package main

import (
	"fmt"

	"github.com/andremelinski/pos-goexpert/10-eventos/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func main(){
	msgs := make(chan amqp091.Delivery)
	chn, err := rabbitmq.OpenChannel()
	if err != nil{
		panic(err)
	}
	go rabbitmq.Consume(chn, msgs)
	for v := range msgs {
		fmt.Println(string(v.Body))
		v.Ack(false)
	}
}