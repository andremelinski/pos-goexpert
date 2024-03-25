package routes

import (
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/db"
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"
)

type ProductRoutes struct{
	Chi *chi.Mux
	DbConfig *gorm.DB
	JwtAuth *jwtauth.JWTAuth
}

func ProductRoutesInit(chi *chi.Mux, dbConfig *gorm.DB, jwtAuth *jwtauth.JWTAuth)*ProductRoutes{
	return &ProductRoutes{
		chi,
		dbConfig,
		jwtAuth,
	}
}

func (routes ProductRoutes) ProductRoutes(){
	productDB := db.ProductInitDB(routes.DbConfig)
	productHandler := handlers.ProductHandlerInit(productDB)
	
	routes.Chi.Route("/product", func (r chi.Router){
		r.Use(jwtauth.Verifier(routes.JwtAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProductById)
		r.Put("/{id}", productHandler.ProductUpdate)
		r.Delete("/{id}", productHandler.ProductDelete)
	})
	routes.Chi.Get("/products", productHandler.GetProducts)
}