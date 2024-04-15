package main

func main(){
	forever := make(chan string)
	go recebe(forever)
	ler(forever)
}

// canal com seta pro lado direito apenas recebe valores
func recebe(ch chan <- string){
		ch <- "hello"
}

// canal com seta pro lado esquerdo apenas recebe valores
func ler(ch <-chan string){
	println(<-ch)
}