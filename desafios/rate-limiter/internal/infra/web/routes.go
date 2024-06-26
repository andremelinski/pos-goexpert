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

// struct recebe a interface que possui os endpoints desse usecase + middlewares
type WebRouter struct {
	HelloWebHandler       interfaces.HelloWebHandlerInterface
	Middlewares interfaces.MiddlewareInterface
}

func NewWebRouter(
	helloWebHandler interfaces.HelloWebHandlerInterface,
	middlewares interfaces.MiddlewareInterface,
) *WebRouter {
	return &WebRouter{
		helloWebHandler,
		middlewares,
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

func (s *WebRouter) BuilMiddlewares() []Middleware{
		return []Middleware{
		{
			Name: "rateLimit",
			Handler: s.Middlewares.RateLimit,
		},
	}
}