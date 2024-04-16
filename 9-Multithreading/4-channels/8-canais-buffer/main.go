package main

//  Buffer -> qundo vc pode mandar +1 info pra dentro do canal. pra isso, vc precisa ter workers suficientes para conseguir consumir tudo o que eh colocado ali
// Nao recomendado

func main(){
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"
	close(ch)
	ler(ch)
}
// // canal com seta pro lado esquerdo apenas recebe valores
func ler(ch chan string){
	for v := range ch {
		println(v)
	}
}