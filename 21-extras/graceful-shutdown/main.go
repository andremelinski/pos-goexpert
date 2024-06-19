package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4*time.Second)
		w.Write([]byte("Hello word"))
	})
	
	go func() {
		fmt.Println("Server is running at http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// block o programa e soh esvazia quando receber uma flag do sistema operacional avisando quando deu erro
	stop := make(chan os.Signal, 1)
	// toda vez que ocorrer qualquer uma dessas acoes, canal eh esvaziado e programa vai acabar
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	// timeout deve ser maior que o tempo que a req leva para ser completa, caso contrario, aaprece erro que nao pode realizar o gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Shutting down server")
	if err := server.Shutdown(ctx); err != nil{
		log.Fatalf("could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("server stopped")
}