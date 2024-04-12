package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var views uint64 = 0 


// Thread 1
// func main(){
// 	// caso ocorra um numero mt alto de acesso (ex 10k chamadas e 1000 simultaneamente) nao vai retornar os 10k pq deu erro de concorrencia:
// 	// 2 ou mais chamadas foram realizadas ao mesmo tempo e com isso, tiveram acesso ao msm valor de variavel e com isso gerando um valor diferente, ex: 9989 e nao 10k
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		views++
// 		w.Write([]byte(fmt.Sprintf("visitante de numero %d", views)))
// 	})
// 	http.ListenAndServe(":8080", nil)
	
// }

// utilizando atomic -> utilizando um pacote nativo do golang pra mudar o valor da variavel, utilizando Mutex (lock e unlock da variavel)
func main(){
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// views++
		atomic.AddUint64(&views, 1)
		w.Write([]byte(fmt.Sprintf("visitante de numero %d", views)))
		// m.Unlock()
	})
	http.ListenAndServe(":8080", nil)
	
}