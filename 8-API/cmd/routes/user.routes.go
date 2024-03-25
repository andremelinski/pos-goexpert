package routes

import (
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/db"
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type RoutesInit struct{
	Chi *chi.Mux
	DbConfig *gorm.DB
}
func UserRoutesInit(chi *chi.Mux, dbConfig *gorm.DB)*RoutesInit{
	return &RoutesInit{
		chi,
		dbConfig,
	}
}

func (routes RoutesInit) UserRoutes(){
	userDB := db.UserInitDB(routes.DbConfig)
	userHandler := handlers.UserHandlerInit(userDB)
	routes.Chi.Route("/user", func (r chi.Router){
		r.Post("/", userHandler.CreateUser)
		r.Get("/{email}", userHandler.GetUserByMail)
		r.Post("/generateJwt", userHandler.GetJWT)
	})
}