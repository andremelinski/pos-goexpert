package main

import (
	"fmt"
	"sync"
)

var lenghThread = 10 
var lenghThreadAnonymous = 5 

func task(name string, wg *sync.WaitGroup){
	for i := 0; i < lenghThread; i++ {
		fmt.Printf("Task %s: %d ", name, i)
		wg.Done()
	}
}

// Thread 1
func main(){
	// ja que a thread 1 nao ve as outras, o main soh via o time.Sleep. Para esperar as outras threads usa o wait group
	waitGroup := sync.WaitGroup{}
	totalLength := lenghThread + lenghThread + lenghThreadAnonymous
	waitGroup.Add(totalLength)
	// Thread 2
	go task("A", &waitGroup) 
	// Thread 3
	go task("B", &waitGroup) 

	go func(name string){
		for i := 0; i < lenghThreadAnonymous; i++ {
			fmt.Printf("Task %s: %d is anonymous \n", name, i)
			waitGroup.Done()
			
		}
	}("C")
	
	waitGroup.Wait()
}

/**
wg.Add() -> quantos eventos o thread main vai ter que esperar antes de finalizar
wg.Wait() -> espera finalizar todas as threads antes de seguir  
wg.Done() -> subtrai a contagem a cada execucao realizada 
se errar na contagem de rotinas e clolocar a mais do que existe: fatal error: all goroutines are asleep - deadlock!
**/