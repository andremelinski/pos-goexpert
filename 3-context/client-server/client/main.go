package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

// exemplo: client tenta fazer uma req que nao pode passar de 5sec se nao ele cancela pelo context
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// ao final de tudo, vc fecha a req
	defer cancel()
	//  preparando a request com o context. Context -> se a req passar de 5 sec, ela eh cancelada
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, res.Body)

	defer res.Body.Close()

}