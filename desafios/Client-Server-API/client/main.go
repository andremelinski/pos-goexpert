package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/Client-Server-API/client/document"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
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

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err = document.DocumentInit(string(bodyBytes)).CreateFile()
			if err != nil {
			panic(err)
		}
	}

	io.Copy(os.Stdout, res.Body)

	defer res.Body.Close()
}
