package main

import "fmt"

var views uint64 = 0 

// Thread 1
func main(){
	// criando um canal com um valor atrelado do tipo string
	canal := make(chan string)

	//  thread 2
	go func ()  {
		canal <- "Ola Mundo" // canal agora tem o valor Ola mundo	
	}()

	// retirando o valor setado na thread 2 na 1
	msg := <-canal // canal agora esta vazio
	fmt.Println(msg)

	//  thread 3
	go func ()  {
		canal <- "Ola Mundo 2" // canal agora tem o valor Ola mundo	
	}()

	msg = <-canal // canal agora esta vazio
	fmt.Println(msg)
	
}