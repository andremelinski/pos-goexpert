package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


type WebServer struct {
	Router        chi.Router
	Handlers      []RouteHandler
	Middlewares   []Middleware
	WebServerPort int
}

func NewWebServer(
	serverPort int,
	handlers []RouteHandler,
	middlewares []Middleware,
) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      handlers,
		Middlewares:   middlewares,
		WebServerPort: serverPort,
	}
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)

	for _, middleware := range s.Middlewares {
		// intercepta todas as requisicoes e injeta limiter
		s.Router.Use(middleware.Handler)
	}

	for _, handler := range s.Handlers {
		s.Router.MethodFunc(handler.Method, handler.Path, handler.HandlerFunc)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", s.WebServerPort), s.Router)
}