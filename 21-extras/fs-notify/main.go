package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

// supor que vc tem  rotacionamento de senha e vc nao quer que o sistema fique fora do ar. Ex: mudar senha de banco a cada x tempo sem parar a aplicacao

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

// objetivo: caso mude qualquer coisa dentro do config.json, o programa deve atualizar automatico, sem parar
func main(){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	marshalConfig("config.json")
	fmt.Println(config)

	// segura o for loop. pode ser feito com wait group
	done := make(chan bool)

	go func() {
		for{
			select{
			// dentr odo loop que fica observando as mudancas, caso watcher pegue alguma mudanca 
			case event, ok := <- watcher.Events:
				if !ok {
					return
				}

				fmt.Println("event :", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					marshalConfig("config.json")
					fmt.Println("modified file:", event.Name)
					fmt.Println(config)
				}

			case err, ok := <- watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	// adicionando o arquivo para ficar sendo observado dentro da thread com loop infinito. Com isso, ele observa e pode rodar outras coisas 
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}


	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})


	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", ":3000", err)
	}
	close(done)
	// <- done
}

func marshalConfig(file string){
	data, err := os.ReadFile(file); if err != nil{
		panic(err)
	}

	err = json.Unmarshal(data, &config); if err != nil{
		fmt.Println(err.Error())
		panic(err)
	}
}