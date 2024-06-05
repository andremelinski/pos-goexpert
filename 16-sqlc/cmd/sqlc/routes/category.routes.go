package routes

import (
	"database/sql"

	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/db"
	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/web/handler"
	"github.com/go-chi/chi"
)

type RoutesInit struct{
	Chi *chi.Mux
	DbConfig *sql.DB
}

func CategoryRoutesInit(chi *chi.Mux, dbConfig *sql.DB)*RoutesInit{
	return &RoutesInit{
		chi,
		dbConfig,
	}
}


func (routes RoutesInit) CategoryRoutes(){

	routes.Chi.Route("/category", func (r chi.Router){
		a := db.New(routes.DbConfig)
		categoryHandler := handler.CategoryHandlerInit(a)
		r.Post("/", categoryHandler.CreateUser)
		// r.Get("/{id}", categoryHandler.GetCategoryById)
	})
}