package main

import (
	"fmt"
	"log"
	"net/http"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("entrou no handler")

		defer func ()  {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				// debug.PrintStack() // -> aparece a msg de panic
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}	
		}()
		next.ServeHTTP(w, r)
	})
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	log.Println("Listening on", ":3000")
	if err := http.ListenAndServe(":3000", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", ":3000", err)
	}
}