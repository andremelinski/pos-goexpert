package main

import (
	"net/http"
)

type blogName struct {
	name string
}

func main() {
	// creating a multiplexing server ()
	// usefull when you need to use 2 or more servers at same time
	// or pass default parameters through
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func1)
	mux.Handle("/", blogName{"andre"})
	http.ListenAndServe(":8080", mux)

	// adding second mux server
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func2)
	http.ListenAndServe(":6000", mux2)
}

func func1(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello word"))
	res.WriteHeader(http.StatusOK)
}

func func2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello word"))
	res.WriteHeader(http.StatusOK)
}

func (b blogName) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello " + b.name))
	res.WriteHeader(http.StatusOK)
}
