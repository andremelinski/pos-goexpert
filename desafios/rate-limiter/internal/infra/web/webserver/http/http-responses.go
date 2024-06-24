package http

import (
	"encoding/json"
	"net/http"
)

// aplica o a interface que faz a resposta REST, ou seja, WebResponseHandlerInterface

type WebResponseHandler struct{}

func NewWebResponseHandler() *WebResponseHandler{
	return &WebResponseHandler{}
}

func (hrh *WebResponseHandler) Respond(w http.ResponseWriter, statusCode int, data any){
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(&data)
	}
}

func (hrh *WebResponseHandler) RespondWithError(w http.ResponseWriter, statusCode int, err error){
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
}