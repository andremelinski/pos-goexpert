package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEPInfo struct {
	Cep string `json:cep` 
	Logradouro string `json:logradouro`
	Complemento string `json:complemento`
	Bairro string `json:bairro`
	Localidade string `json:localidade`
	UF string `json:uf`
	IBGE string `json:ibge`
	Gia string `json:gia`
	DDD string `json:ddd`
	Siafi string `json:siafi`
}

func main() {
	fmt.Println(os.Args )
	for _, cep := range os.Args[1:] {
		fmt.Println(cep)
		req, err := http.Get("https://viacep.com.br/ws/"+cep+"/json/")
		if err != nil {
			panic(err)
		}

		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		
		if err != nil {
			panic(err)
		}
		data := CEPInfo{}
		err = json.Unmarshal(res, &data)
		if err != nil {
			panic(err)
		}
		fileName := "cep.txt"
		// filePointer, err := os.Create(fileName)
		filePointer, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)

		if err != nil {
			panic(err)
		}
		filePointer.WriteString(fmt.Sprintf("CEP: %s, Localidade %s, UF: %s \n", data.Cep, data.Localidade, data.UF))

		defer filePointer.Close()

		fmt.Println(data)
	}
}