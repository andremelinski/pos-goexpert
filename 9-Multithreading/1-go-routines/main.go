package main

import (
	"fmt"
	"time"
)

func task(name string){
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s: %d ", name, i)
		time.Sleep(1*time.Second)
	}
}

// Thread 1
func main(){
	// sincrono -> executa A e depois B
	// task("A") 
	// task("B")

	// *** colocando o "go" na frente, a funcao vira uma thread
	
	// com o "go" apenas na funcao task("A") e na "B", se faz necessario executar a thread task A para chegar a B, com isso o programa roda.
	// se vc criar uma thread para cada task, nao funciona pq nao tem nada que exiga com que o programa seja "segurado" enqaunto as threads sao executadas, com isso, ele nao loga nada. Para rodar, vc precisa "forcar" um outro processo fora das threads
	// Thread 2
	go task("A") 
	// Thread 3
	go task("B") 

	go func(name string){
		for i := 0; i < 3; i++ {
			fmt.Printf("Task %s: %d is anonymous \n", name, i)
			time.Sleep(1*time.Second)
		}
	}("C")
	time.Sleep(3*time.Second)
}