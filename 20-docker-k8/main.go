package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello word"))
	})
	fmt.Println("hello")
	http.ListenAndServe(":8080", nil)
}