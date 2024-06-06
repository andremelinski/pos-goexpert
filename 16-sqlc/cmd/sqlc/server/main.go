package main

import (
	"database/sql"
	"net/http"

	"github.com/andremelinski/pos-goexpert/16-sqlc/cmd/sqlc/routes"
	configs "github.com/andremelinski/pos-goexpert/16-sqlc/config"
	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	config := configs.LoadConfig(".")
	rMux := chi.NewRouter()

	dsn := config.DBUser+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBName+"?charset=utf8mb4&parseTime=True&loc=Local"

	dbConfig, err := sql.Open("mysql", dsn)

	if err !=nil{
		panic(err)
	}
	routes.CategoryRoutesInit(rMux, dbConfig).CategoryRoutes()
	
	http.ListenAndServe(":8000", rMux)
}