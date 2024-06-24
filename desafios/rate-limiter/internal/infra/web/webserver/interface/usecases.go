package interfaces

import "net/http"

type HelloWebHandlerInterface interface {
	Hello(w http.ResponseWriter, r *http.Request)
}