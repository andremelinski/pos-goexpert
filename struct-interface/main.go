package main

import "fmt"

//  interface
// toda struct que tiver o metodo desativar vai implementar a interface Pessoa
// o attachment do struct Client pra interface Pessoa eh automatico
// para implementar interface soh vale metodos e nao propriedades
type Pessoa interface {
	Desativar()
}

//struct: como se fosse o constructor de uma classe
// struct -> descreve as caracteristicas do que voce vai utilizar
type Cidade struct{
	Logradouro string
	Cidade string
	Estado string
}
// struct composta
type Client struct {
	Nome string
	Idade int
	Ativo bool
	Cidade 
}

type Empresa struct {
	Nome string
}

func main() {
	p1 := Client{Nome:"Andre", Idade:23, Ativo:true, Cidade: Cidade{
		Logradouro: "casa", Cidade:"Curitiba", Estado: "PArana",
	}}
		fmt.Println(p1)
	// mudando status diretamente pelo struct
	// p1.Desativar()

	// mudando de status utiliando interface.
	// p1 eh uma pessoa pq ele implementa o metodo desativar
	DesativacaoPorInterface(p1)
	fmt.Println(p1)
	p1.Cidade.Cidade = "new city"

	minhaEmpresa := Empresa{"empresa 1"}
	// minhaEmpresa eh uma pessoa pq ele implementa o metodo desativar
	DesativacaoPorInterface(minhaEmpresa)
}

// metodos em structs
// assim como em OOP, structs podem ter metodos que sao aplicados ao utilizar o struct
// o metodo Desativar faz parte do struct Client. Com isso, toda vez que utilizar o struct, como p1, ele tera esse metodo
// por fazer parte de Client, ele tambem tera acesso as propriedades dentro dele
func(c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("desativando %s", c.Nome)
}

func(e Empresa) Desativar() {
	fmt.Printf("desativando empresa")
}

func DesativacaoPorInterface(p Pessoa){
	p.Desativar()
}