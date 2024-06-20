package main

import "fmt"

func main(){
	forever := make(chan bool)

	// go func(){
	// 	for i := 0; i < 20; i++ {
	// 		println(i)
	// 	}
	// 	// nao da dead lock se vc enche o canal com alguma coisa.
	// 	// valor precisa estar dentro de outra go routine para nao dar erro de dead lock
	// 	forever <- true
	// }()
		
	// <- forever
	i :=0

		go func(){
		for{
			i++
			println(i)
			if i == 20 {
				forever <- true
				return
			}
		}
		// nao da dead lock se vc enche o canal com alguma coisa.
		// valor precisa estar dentro de outra go routine para nao dar erro de dead lock
		
	}()
		
	<- forever
	fmt.Println("acabou")
}