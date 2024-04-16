package main

import (
	"fmt"
	"time"
)

var goFuncWrks = 10

func main(){
	data := make(chan int)
	// cria "goFuncWrks" go routines
	for i := 0; i < goFuncWrks; i++ {
		go worker(i, data)
	}
	
	ler(data)
}

func worker(workerId int, ch <- chan int){
	for v := range ch {
		fmt.Printf("workerId %d received %d\n", workerId, v)
		time.Sleep(2*time.Second)
	}
}

// canal com seta pro lado esquerdo apenas recebe valores
func ler(ch chan<- int){
	for i := 0; i < goFuncWrks; i++ {
		ch <- i
	}
}