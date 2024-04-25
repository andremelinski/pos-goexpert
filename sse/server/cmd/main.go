// db config
package main

import (
	"net/http"

	"github.com/andremelinski/pos-goexpert/sse/server/cmd/routes"
	"github.com/andremelinski/pos-goexpert/sse/server/config"
)

func main(){
	dbConfig := config.GetDynamodbClient()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	//  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     w.Write([]byte("Hello, this is the homepage!"))
    // })

	routes.NewNotificationRoutes(mux, dbConfig).NotificationRoutesHandler()
	http.ListenAndServe(":8000", mux)
}

	