package main

import (
	"strconv"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/configs"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/handlers"
	http_response "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/http"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/middleware"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/middleware/strategy"
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

	// aplica WebResponseHandlerInterface -> expoe uma resposta http retornando Respond or RespondWithEror 
	// msm coisa do TS para erro customizavel e quando quer pegar a interface Error
	httpResp := http_response.NewWebResponseHandler()

	// --- MIDDLEWARE ---
	rateLimitConfig := middleware.RateLimiterConfig{
		MaxRequestsPerIP: configs.IPMaxRequests,
		MaxRequestsPerToken: configs.TokenMaxRequests,
		TimeWindowMilliseconds: configs.TimeWindowMilliseconds,
	}

	// layer com conexao com o banco e regra do rate limit
	strategy := strategy.NewStrategyRateLimit(client)
	mid := middleware.NewRateLimiterMiddleware(rateLimitConfig, strategy, httpResp)

	// --- HANDLER ---
	helloWebHandler := handlers.NewHelloWebHandler(httpResp)

	// --- WEB SERVER ---
	// injeto todos os handler que a interface HelloWebHandlerInterface usa para a montagem de metodos (post, get, etc)
	webRouter := web.NewWebRouter(helloWebHandler, mid)

	webServer := web.NewWebServer(
		configs.WebServerPort,
		webRouter.BuildHandlers(),
		 webRouter.BuilMiddlewares(),
	)

	webServer.Start()

}