package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ViaCEPInfo struct {
	Api string 
	Cep string `json:"cep"` 
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	IBGE string `json:"ibge"`
	Gia string `json:"gia"`
	DDD string `json:"ddd"`
	Siafi string `json:"siafi"`
}

type BrasilInfo struct {
	Api string 
	Cep string `json:"cep"` 
	Logradouro string `json:"street"`
	Bairro string `json:"neighborhood"`
	Localidade string `json:"city"`
	UF string `json:"state"`
}

type ExternalApis struct{
	Cep string
}


func ExternalApisInit(cep string) *ExternalApis{
	return &ExternalApis{
		cep,
	}
}

func (e *ExternalApis) BrasilApi(c1 chan *BrasilInfo){
	// time.Sleep(20*time.Second)
	url := "https://brasilapi.com.br/api/cep/v1/"+e.Cep
	resp, err := http.DefaultClient.Get(url)
	
	if err != nil{
		fmt.Println(err)
		return
	}

	body := resp.Body
	res, err := io.ReadAll(body)
	
	if err != nil{
		fmt.Println(err)
		return 
	}

	cepInfo := BrasilInfo{}
	json.Unmarshal(res, &cepInfo)

	if cepInfo.Bairro == "" {
		fmt.Println(errors.New(string(res)))
		return
	}
	
	defer resp.Body.Close()

	cepInfo.Api = "brasilApi"
	c1 <- &cepInfo

}

func (e *ExternalApis) ViaCepAPI(c2 chan *ViaCEPInfo){
	
	url := "https://viacep.com.br/ws/"+e.Cep+"/json/"
	resp, err := http.DefaultClient.Get(url)
	body := resp.Body
	
	data, _ := io.ReadAll(body)
	
	defer resp.Body.Close()
		
	if err != nil{
		fmt.Println(err)
		return 
	}
	cepInfo := ViaCEPInfo{}
	json.Unmarshal(data, &cepInfo)

	if cepInfo.Bairro == "" {
		fmt.Println(errors.New(string(data)))
		return 
	}

	cepInfo.Api = "ViaCep"
	c2 <- &cepInfo
}