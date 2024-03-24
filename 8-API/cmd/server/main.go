package main

import (
	"net/http"

	"github.com/andremelinski/pos-goexpert/8-API/configs"
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/db"
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	config, _ := configs.LoadConfig(".")
	println(config.DBDrive)

	dbConfig, err := gorm.Open(sqlite.Open("test.db"))

	if err !=nil{
		panic(err)
	}

	dbConfig.AutoMigrate(&entity.Product{}, &entity.User{})

	// pra utilizar o controller e fazer a juncao das layers controller e DB deve-se injetar productDBInit, o qual poossui os metodos da interface
	userDB := db.UserInitDB(dbConfig)
	userHandler := handlers.UserHandlerInit(userDB, config.TokenAuth, config.JwtExpiresIn)

	productDB := db.ProductInitDB(dbConfig)
	productHandler := handlers.ProductHandlerInit(productDB)

	rMux := chi.NewRouter()
	rMux.Use(middleware.Logger) // intercepta todas as requisicoes e injeta logs 

	rMux.Route("/user", func (r chi.Router){
		r.Post("/", userHandler.CreateUser)
		r.Get("/{email}", userHandler.GetUserByMail)
		r.Post("/generateJwt", userHandler.GetJWT)
	})

	rMux.Route("/product", func (r chi.Router){
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProductById)
		r.Put("/{id}", productHandler.ProductUpdate)
		r.Delete("/{id}", productHandler.ProductDelete)
	})
	rMux.Get("/products", productHandler.GetProducts)

	http.ListenAndServe(":8000", rMux)
}