package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// utilizado quando pelo padrao Fan-it:  multiple functions concurrently and then performing some aggregation on the results

type Result struct{
	id int64
	Body string
}

func main(){
	c1 := make(chan Result)
	c2 := make(chan Result)
	// como o i esta no escopo global e utilizado por 2 threads, pode ocorrer de ter concorrencia na hora de somar 1. Para resolver, adicionar de forma atomica
	i :=int64(0)
	// recebe do RabbitMQ
	go func() {
		
		for {
			time.Sleep(time.Second*4)
			msg := Result{i,"Message from RabbitMQ"}
			c1 <- msg
			atomic.AddInt64(&i, 1)
		}
	}()
	// recebe do Kafka
	go func() {
		for {
			time.Sleep(time.Second*4)
			msg := Result{i,"Message from Kafka"}
			c2 <- msg
			atomic.AddInt64(&i, 1)
		}
	}()

	// looping infinito para ficar pegando msg
	for{
		select{
		case msg1 := <- c1:
			fmt.Printf("msg RabbitMQ came first ID: %d - %s \n", msg1.id, msg1.Body)
		case msg2 := <- c2:
			fmt.Printf("msg Kafka came first ID: %d - %s \n",msg2.id, msg2.Body)
		// regra para timeout se os canais demorarem para ser populados
		case <- time.After(time.Second*5):
			println("timeout")
		}
	}

}

