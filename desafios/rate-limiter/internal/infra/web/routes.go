package web

import (
	"net/http"

	interfaces "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/interface"
)

type RouteHandler struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Middleware struct {
	Name    string
	Handler func(next http.Handler) http.Handler
}

// struct recebe a interface que possui os endpoints desse usecase
type WebRouter struct {
	HelloWebHandler       interfaces.HelloWebHandlerInterface
	// RateLimiterMiddleware middlewares.RateLimiterMiddlewareInterface
}

func NewWebRouter(
	helloWebHandler interfaces.HelloWebHandlerInterface,
	// rateLimiterMiddleware middlewares.RateLimiterMiddlewareInterface,
) *WebRouter {
	return &WebRouter{
		       helloWebHandler,
		// RateLimiterMiddleware: rateLimiterMiddleware,
	}
}


// metodo para cadastrar todas as rotas
func (s *WebRouter) BuildHandlers() []RouteHandler {
	return []RouteHandler{
		{
			Path:        "/",
			Method:      "GET",
			HandlerFunc: s.HelloWebHandler.Hello,
		},
	}
}

// func (s *WebServer) AddMiddleware(middleware Middleware) {
// 	s.Middlewares = append(s.Middlewares, middleware)
// }