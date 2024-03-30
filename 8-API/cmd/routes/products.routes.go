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
	
	routes.Chi.Route("/products", func (r chi.Router){
		// chi usa o middleware jwtauth.Verifier pra verificar onde esta esse token e se ele eh valido, com base o secret na chave jwtauth.New("HS256", []byte(cfg.JwTSecret), nil);
		// jwtauth.Authenticator -> pega o token que foi enviado pela requisicao e injeta o JwtAuth no nosso contexto e validaa de fato o jwt, batendo assinatura e expiration
		r.Use(jwtauth.Verifier(routes.JwtAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProductById)
		r.Put("/{id}", productHandler.ProductUpdate)
		r.Delete("/{id}", productHandler.ProductDelete)
	})
}