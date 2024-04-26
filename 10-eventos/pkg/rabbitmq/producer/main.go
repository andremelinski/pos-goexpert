package main

import (
	"github.com/andremelinski/pos-goexpert/10-eventos/pkg/rabbitmq"
)

func main(){
	chn, err := rabbitmq.OpenChannel()
	if err != nil{
		panic(err)
	}

	defer chn.Close()
	rabbitmq.Publisher(chn, "Hello world", "amq.direct")
	
}