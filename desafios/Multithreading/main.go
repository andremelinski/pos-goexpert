package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/Multithreading/apis"
)

//  go run main.go -cep=82540091
// testes: colocar time.Sleep(time.Second*5) em 1 e ver se retorna outro
// para simular que um CEP nao foi emcontrado em um mas em outro: setar o CEP em uma url e mandar o comando com um CEO errado no outro
func main(){
	ch1 := make(chan *apis.BrasilInfo)
	ch2 := make(chan *apis.ViaCEPInfo)
	cep := strings.Split(os.Args[1], "=")[1]
	cepInfo := apis.ExternalApisInit(cep)
	
	go cepInfo.BrasilApi(ch1)
	go cepInfo.ViaCepAPI(ch2)

	select{
	case msg1 := <-ch1:
		json.NewEncoder(os.Stdout).Encode(msg1)
	case msg2 := <-ch2:
		json.NewEncoder(os.Stdout).Encode(msg2)
	case <- time.After(time.Second*1):
		json.NewEncoder(os.Stdout).Encode("timeout")
	}

}