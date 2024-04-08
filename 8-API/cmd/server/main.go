package main

import (
	"log"
	"net/http"

	"github.com/andremelinski/pos-goexpert/8-API/cmd/routes"
	"github.com/andremelinski/pos-goexpert/8-API/configs"
	_ "github.com/andremelinski/pos-goexpert/8-API/docs"
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Andre Melinski
// @contact.url    http://www.fullcycle.com.br
// @contact.email  andremelinski29@gmail.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main(){
	config, _ := configs.LoadConfig(".")
	dbConfig, err := gorm.Open(sqlite.Open("test.db"))
	// dsn := config.DBUser+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.WebServerPort+")/"+config.DBName+"?charset=utf8mb4&parseTime=True&loc=Local"
	// dbConfig, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err !=nil{
		panic(err)
	}

	dbConfig.AutoMigrate(&entity.Product{}, &entity.User{})
	rMux := chi.NewRouter()
	rMux.Use(middleware.Logger) // intercepta todas as requisicoes e injeta logs 
	rMux.Use(middleware.Recoverer)
	rMux.Use(LogRequest) // meu middleware intercepta todas as requisicoes e injeta logs, mas que poderia ser feito pra voltar sempre application/json 
	// envia pelo contexto da req o metodos 
	rMux.Use(middleware.WithValue("jwt", config.TokenAuth))
	rMux.Use(middleware.WithValue("jwtExpiresIn", config.JwtExpiresIn))

	
	routes.UserRoutesInit(rMux, dbConfig).UserRoutes()
	routes.ProductRoutesInit(rMux, dbConfig, config.TokenAuth).ProductRoutes()

	rMux.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://"+config.DBHost+":"+config.WebServerPort+"/docs/doc.json")))

	http.ListenAndServe(":8000", rMux)
}

func LogRequest(next http.Handler) http.Handler{
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request){
		log.Printf("my request: %s --- response: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w,r)
	})
}