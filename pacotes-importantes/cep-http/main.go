package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// curl http://localhost:8080/?cep=82510100
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
	http.HandleFunc("/", buscaCep)
	http.ListenAndServe(":8080", nil)
}

func buscaCep (res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-type", "application/json")

	cepParam := req.URL.Query().Get("cep")
	if cepParam == ""{
		res.Write([]byte("CEP not found"))
		res.WriteHeader(http.StatusNotFound)
		return
	}

	cepInfo, err := buscaCEP(cepParam)
	if err != nil{
		res.Write([]byte(err.Error()))
		res.WriteHeader(http.StatusNotFound)
		return
	}
	// utilizando marshal para voltar
	// b, err := json.Marshal(cepInfo)
	// if err != nil{
	// 	res.Write([]byte(err.Error()))
	// 	res.WriteHeader(http.StatusNotFound)
	// 	return
	// }	
	// res.Write(b)
	// utilizando encoder
	json.NewEncoder(res ).Encode(cepInfo)
	res.WriteHeader(http.StatusOK)
} 

func buscaCEP(cep string) (*CEPInfo,error) {
	c := http.Client{Timeout: 2*time.Second}
	req, err := c.Get("https://viacep.com.br/ws/"+cep+"/json/")
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	
	if err != nil {
		return nil, err
	}

	cepInfo := CEPInfo{}
	json.Unmarshal(res, &cepInfo)
	return &cepInfo, nil
}