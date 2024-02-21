package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context: permite com que passe as infos para diversa chamadas, podendo ser cancelado para poupar tempo
	ctx := context.Background() // default -> context vazio
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()  // releases resources if slowOperation completes before timeout elapses

}

func bookHotel(ctx context.Context){
	// espereva o resultado chegar (tipo um switch async) e dai faz algo
	select {
	case <- ctx.Done():
		fmt.Println("aconteceu timeout")
		return 
	case <- time.After(5 * time.Second):
		fmt.Println("passou 5 sec. Hotel reservado")
	}
}