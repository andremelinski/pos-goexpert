package main

import (
	"log"
	"net/http"

	"github.com/andremelinski/pos-goexpert/8-API/cmd/routes"
	"github.com/andremelinski/pos-goexpert/8-API/configs"
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	config, _ := configs.LoadConfig(".")
	dbConfig, err := gorm.Open(sqlite.Open("test.db"))

	if err !=nil{
		panic(err)
	}

	dbConfig.AutoMigrate(&entity.Product{}, &entity.User{})
	rMux := chi.NewRouter()
	rMux.Use(middleware.Logger) // intercepta todas as requisicoes e injeta logs 
	rMux.Use(middleware.Recoverer)
	rMux.Use(LogRequest)
	// envia pelo contexto da req o metodos 
	rMux.Use(middleware.WithValue("jwt", config.TokenAuth))
	rMux.Use(middleware.WithValue("jwtExpiresIn", config.JwtExpiresIn))

	
	routes.UserRoutesInit(rMux, dbConfig).UserRoutes()
	routes.ProductRoutesInit(rMux, dbConfig, config.TokenAuth).ProductRoutes()

	http.ListenAndServe(":8000", rMux)
}

func LogRequest(next http.Handler) http.Handler{
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request){
		log.Printf("my request: %s --- response: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w,r)
	})
}