package main

import "fmt"

func main(){
	forever := make(chan int)
	go publisher(forever)
	// consumer precisa estar na thread principal para que o publisher seja processado
	consumer(forever)
	
}

func publisher(ch chan int){
	for i := 0; i < 5; i++ {
		println("antes do canal ter o valor")
		ch <- i
		println(i)
		// for soh continua se o canal for liberado/consumido
	}
	//  apos o for loop nada mais entra no canal e com isso, nao da pra publicar mais e dai nao vem deadlock 
	close(ch)
}

func consumer(ch chan int){
	for v := range ch {
		fmt.Printf("REcebeu %v \n", v)
	}
}