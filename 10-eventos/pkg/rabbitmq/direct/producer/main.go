package main

import (
	"sync"

	rabbitmq "github.com/andremelinski/pos-goexpert/10-eventos/pkg/rabbitmq/direct"
)

type Payload struct{
		body string
		exchangeName string
		keyName string
	}

func main(){
	chn, err := rabbitmq.OpenChannel()
	if err != nil{
		panic(err)
	}

	defer chn.Close()
	wg := sync.WaitGroup{}
	
	arrInfo := []Payload{
		{body: "Hello direct 0", exchangeName: "amq.direct", keyName: ""}, 
		{body: "Hello direct 1", exchangeName: "amq.direct", keyName: ""},
		{body: "Hello direct 2", exchangeName: "amq.direct", keyName: ""},
		{body: "Hello blue", exchangeName: "amq.direct", keyName: "blue"},
		{body: "Hello green", exchangeName: "amq.direct", keyName: "green"},
		{body: "Hello from queue green using black key", exchangeName: "amq.direct", keyName: "black"},
	}
	numberOfWg := len(arrInfo)
	wg.Add(numberOfWg)
	for i := 0; i < numberOfWg; i++ {
		info := arrInfo[i]
		go rabbitmq.Publisher(chn, info.body, info.exchangeName, info.keyName, &wg)
		
	}

	wg.Wait()
}