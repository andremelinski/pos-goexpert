package interfaces

import "net/http"

type WebResponseHandlerInterface interface {
	Respond(w http.ResponseWriter, statusCode int, data any)
	RespondWithError(w http.ResponseWriter, statusCode int, err error)
}