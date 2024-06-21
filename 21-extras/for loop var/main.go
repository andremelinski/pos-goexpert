package main

import "fmt"

func main() {
	// x := 10
	// for i := range x {
	// 	fmt.Println(i)
	// }

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		// antes: criava um novo v com esse escopo dentro do looop que ta rodando
		// v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		a := <-done
		fmt.Printf("values %v\n", a)
	}
}