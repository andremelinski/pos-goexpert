package main

import "fmt"

// quando receber o nome, deve alterar o nome original e adc um sobrenome
type Cliente struct {
	nome string
}

// struct vazia -> significa que vc esta criando uma funcao que retorna o endereco da memoria 
// do Cliente que esta sendo criado agora, com o nome vazio ("").
// isso eh usado pq agora em qualquer lugar onde vc passar esse cliente e alterar algo, isso refletira de maneira global
// especie de constructor em OOP que sempre vai retornar o endereco de memoria do struct para ser trabalhado
func newName() *Cliente{
	return &Cliente{nome:""}
}

// para mudar o valor original, deve-se usar o endereco da memoria em que o struct Cliente se encontra
func (c *Cliente) andou() {
	println(c) // endereco da memoria
	// pega o valor do nome endereco da memoria do struct Cliente.
	//Obs: c.nome ou *c.nome eh a msm coisa
	c.nome = "andre melinski"
	fmt.Printf("cliente andou %v andou \n", c.nome)
}

func main() {
	p2 := newName()
	p2.nome = "new name"
	println(newName().nome)
	println(p2.nome)

	p1 := Cliente{"andre"}
	p1.andou()
	fmt.Println(p1.nome)

}