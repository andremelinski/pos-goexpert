package main

import "fmt"

// ponteiros funcionam como caixa do correio
//  ao inves de trabalhar copiando o valor que esta no a, usamos o endereco de memoria

func soma(a,b *int) int{
	*a =50
	return *a + *b
}

func main(){
	a := 10

	// toda vez que usamos o "*" estamos falando do enderecamento que esta na memoria que nesse momento tem o valor de 10
	var ponteiro *int = &a
	println(ponteiro) // retorna o endereco da memoria que o valor de a esta
	*ponteiro = 20
	println(a) // retorna 20 pq esta mudando o valor que esta nesse endereco de memoria
	b := &a // se "b" representa o endereco de memoria que "a". "b" eh um ponteiro de "a"
	println(b) // retorna o endereco da memoria de "a"
	// para retornar o valor do endereco de memoria, usar o dereferencing 
	fmt.Printf("valor que esta na memoria de 'b' %d",*b) // o "*" "pergunta" qual valor que esta guardado nesse valor de memoria
	*b = 5
	println(a) // retorna 5 ja que "b" faz referencia ao endereco de mem que a variavel "a"
	var1 := 2
	var2 := 50
	println(soma(&var1,&var2))
	fmt.Printf("como mudou o valor em mem na func soma, valor de var1 muda: %d",var1)
}