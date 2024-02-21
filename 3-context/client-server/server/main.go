package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	log.Println("Rquest ")
	ctx := r.Context()
	defer log.Println("Rquest finalizada")
	select{
	case <- time.After(5*time.Second):
		log.Println("request passou de 5 segundos")
	case <- ctx.Done():
		log.Println("request cancelada pelo client")
	}
}