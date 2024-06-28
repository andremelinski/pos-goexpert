package main

import (
	"strconv"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/configs"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/middleware"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/middleware/strategy"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/usecase"
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
	httpResp := handlers.NewWebResponseHandler()

	// --- MIDDLEWARE ---
	rateLimitConfig := middleware.RateLimitConfig{
		MaxReqIP: configs.IPMaxRequests,
		MaxReqToken: configs.TokenMaxRequests,
		OperatingWindowMs: configs.OperatingWindowMs,
	}

	// layer com conexao com o banco e regra do rate limit
	rateLimitRepo := database.NewRateLimitRepository(client)
	strategy := strategy.NewStrategyRateLimit(rateLimitRepo)
	mid := middleware.NewRateLimitMiddleware(rateLimitConfig, strategy, httpResp)

	// --- HANDLER ---
	helloUseCase := usecase.NewHelloUseCase()
	helloWebHandler := handlers.NewHelloWebHandler(httpResp, helloUseCase)

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