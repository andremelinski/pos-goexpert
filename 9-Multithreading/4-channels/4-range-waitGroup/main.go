package main

import (
	"fmt"
	"sync"
)

func main(){
	wg := sync.WaitGroup{}
	forever := make(chan int)
	go recebe(forever)
	consumer(forever, &wg)
}

func recebe(ch chan int){
	for i := 0; i < 5; i++ {
		println("antes do canal ter o valor")
		ch <- i
		println(i)
		// for soh continua se o canal for liberado/consumido
	}
	//  apos o for loop nada mais entra no canal e com isso, nao da pra publicar mais e dai nao vem deadlock 
	// fecha o canal pq vc adicionou um valor a ele pra usar depois, mas nao o colocou em uma variavel
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup){
	for v := range ch {
		fmt.Printf("REcebeu %v \n", v)
	}
}