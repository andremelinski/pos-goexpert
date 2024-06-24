package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
 Se o limite por IP é de 10 req/s e a de um determinado token é de 100 req/s, o rate limiter deve utilizar as informações do token.
*/
// o usuario vem e bate na req
var (
	limit = 10
	startTime = time.Now()

)

type UserInfoMapper struct {
	username string
	lastLogin time.Time
	call int
}

type UserMapper struct{
	ip  UserInfoMapper
}

func limiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mapperP :=  &map[string]int{}
		arr := []UserInfoMapper{}
		// [
		// 	id1: {
		// 		rate: 1,
		// 		time: time.Now()
		// 	},
		// ]
    go func() {
        // for {
            for i := 0; i < limit; i++ {
					mapper := *mapperP
					if mapper["aqui"] == 0 {
						mapper["aqui"] = len(arr)

						arr = append(arr, UserInfoMapper{
							username: "aqui",
							call: 1,
							lastLogin: time.Now(),
						})
						next.ServeHTTP(w, r)
						fmt.Println("primeira")
					}else{
						duration := time.Now().Sub(arr[mapper["aqui"]].lastLogin)
						arr[mapper["aqui"]].call++
						
						if arr[mapper["aqui"]].call >4 {
							if duration > time.Second {
								arr[mapper["aqui"]].lastLogin = time.Now()
								arr[mapper["aqui"]].call = 0
								next.ServeHTTP(w, r)
								fmt.Println("zerou tudo")
							}else{
								// time.Sleep(2*time.Second)
								fmt.Println("nop")
							}
						}else{
							arr[mapper["aqui"]].call++
							next.ServeHTTP(w, r)
						}
					}
				// }
			}
			log.Print("Executing middlewareOne again")
        // }
    }()
	})
}


func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()

	helloHandler := http.HandlerFunc(hello)
	mux.Handle("/", limiterMiddleware((helloHandler)))

	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}