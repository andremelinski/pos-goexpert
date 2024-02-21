package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

type ContaTag struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"` // json:"-" desconsidera a chave
}

func main() {
	conta1 := Conta{12, 500}
	res, err := json.Marshal(conta1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// enconder -> pega o arquivo e transforma ele sem precisar criar a variavel res
	//  json.NewEncoder(os.Stdout) => fala onde deve ser enviado o arquivo modificado pelo encoder, no caso vai pro terminal
	enconder := json.NewEncoder(os.Stdout)
	err = enconder.Encode(conta1)
	if err != nil {
		panic(err)
	}

	//***  caminho inverso *** json -> struct
	jsonConta2 := []byte(`{"Numero":2, "Saldo": 200}`)
	conta2 := Conta{}
	/* como vc comeca conta2 com o endereco dela vazio e vc quer popular ela,
	necessario passar o endereco da memoria que conta2 esta, por isso usa ponteiro.
	Caso o unmarshal nao consiga fazer o enderecamento chave da res pro struct da consta2 (Conta),
	ele volta como 0 no campo, ja que o bind da resposta do unmarshal e struct nao foi feito corretamente.
	Se todo o bind estiver errado, ele volta como 0 
	*/
	err = json.Unmarshal(jsonConta2, &conta2)
	if err != nil {
		panic(err)
	}

	fmt.Println(conta2)


	// *** TAGS *** e se eu quisesse user "n" como chave para Numero e "s" para chave de Saldo?
	c1Tag := ContaTag{500, 20}
	// ao realizar o marshal, as chaves Numero e Saldo viram "n" e "s", respectivamente, devido as tags
	bytesC1, err := json.Marshal(c1Tag)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytesC1))

	jsonContaTag := []byte(`{"n":2000, "s": 50}`)
	contaTag := ContaTag{}

	err = json.Unmarshal(jsonContaTag, &contaTag)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaTag)

}