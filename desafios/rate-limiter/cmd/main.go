package main

import (
	"fmt"
	"strconv"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/configs"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/handlers"
	http_response "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/http"
)

func main(){

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	
	dBName, err := strconv.Atoi(configs.DBName)
    if err != nil {
        panic(err)
    }
	client := database.NewRedisClient(configs.DBHost, configs.DBPort, configs.DBPassword, dBName).Client()
	fmt.Println(client.Ping())

	httpResp := http_response.NewWebResponseHandler()

	// meu handler que aplica WebResponseHandlerInterface -> Faz toda logica Rest API e fala se retorna Respond or RespondWithEror 
	// msm coisa do TS para erro customizavel e quando quer pegar a interface Error no catch
	helloWebHandler := handlers.NewHelloWebHandler(httpResp)

	// injeto todos os handler que a interface HelloWebHandlerInterface usa para a montagem de metodos (post, get, etc)
	webRouter := web.NewWebRouter(helloWebHandler)


	webServer := web.NewWebServer(
		configs.WebServerPort,
		webRouter.BuildHandlers(),
		 []web.Middleware{},
		// webRouter.BuildMiddlewares(),
	)

	webServer.Start()

}