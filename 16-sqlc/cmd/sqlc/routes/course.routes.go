package routes

import (
	"database/sql"

	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/db"
	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/web/handler"
	"github.com/go-chi/chi"
)


func CourseRoutesInit(chi *chi.Mux, dbConfig *sql.DB)*RoutesInit{
	return &RoutesInit{
		chi,
		dbConfig,
	}
}


func (routes RoutesInit) CourseRoutes(){
	courseDB := db.NewCourseDB(routes.DbConfig)

	courseHandler := handler.CourseHandlerInit(courseDB)

	routes.Chi.Route("/course", func (r chi.Router){
		r.Post("/", courseHandler.CreateCourse)
		r.Get("/", courseHandler.GetCourses)
	})
}