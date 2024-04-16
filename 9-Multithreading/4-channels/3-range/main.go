package main

import "fmt"

// sincroniza dados utilizando 2 threads diferentes
func main(){
	forever := make(chan int)
	go publisher(forever)
	// consumer precisa estar na thread principal para que o publisher seja processado
	consumer(forever)
	
}

func publisher(ch chan int){
	for i := 0; i < 5; i++ {
		println("antes do canal ter o valor")
		// canal cheio
		// for soh continua se o canal for liberado/consumido.
		ch <- i
		println(i)
	}
	//  apos o for loop nada mais entra no canal e com isso, nao da pra publicar mais e se nao fechar o canal, da deadlock 
	close(ch)
}

// consome o valor que o canal esta no momento, ou seja, o index do for loop
func consumer(ch chan int){
	for v := range ch {
		fmt.Printf("REcebeu %v \n", v)
	}
}