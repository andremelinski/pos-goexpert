package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileNmae := "arquivp.txt"
	f, err := os.Create(fileNmae)
	if err != nil {
		panic(err)
	}
	// usado apenas para escrever string
	tamanho, err := f.WriteString("Hello word")
		fmt.Printf("arquivo criado com string tamanho %d bytes", tamanho)
	if err != nil {
			panic(err)
		}
	// para escrever qualquer coisa, escreve-se bytes no arquivo
	tamanho, err = f.Write([]byte("\n escrevendo bytes algumas informacoes, independente da tipagem: 1234, true") )

	if err != nil {
		panic(err)
	}
	fmt.Printf("adicionado com bytes e arquivo com tamanho %d bytes", tamanho)
	f.Close()

	//  lendo arquivo 
	bytes, err := os.ReadFile(fileNmae)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	// supor que vc tem 100 MB e o arquivo tem 1GB.
	// leitura de pouco em pouco abrindo o arquivo
	bytes2, err := os.Open(fileNmae)
	if err != nil {
	panic(err)
	}
	// criando um buffer a partir do arquivo original
	reader := bufio.NewReader(bytes2)
	buffer := make([]byte, 10)
	for{
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// imprimi o buffer que le, a partir do conteudo que pega do arquivo original
		// n eh a posicao de leitura
		fmt.Println(string(buffer[:n]))
	}

}