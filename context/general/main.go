package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	// context: permite com que passe as infos para diversa chamadas, podendo ser cancelado para poupar tempo
	ctx := context.Background() // default -> context vazio
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()  // releases resources if slowOperation completes before timeout elapses
	// funciona como um timeOut. obedece o context -> tem 1 segundo pra finalizar a req, se nao quebra e no final, 
	// se der tudo certo, ele da um cancel pra liberar a req
	req, err := http.NewRequestWithContext(ctx, "GET", "http:/google.com", nil)

	// executando a request montada pelo NewRequestWithContext
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		panic(err)
	}
	
	defer resp.Body.Close()
}