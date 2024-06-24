package handlers

import (
	"net/http"

	interfaces "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/interface"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/usecase"
)

type HelloWebHandler struct {
	ResponseHandler interfaces.WebResponseHandlerInterface
}

func NewHelloWebHandler(rh interfaces.WebResponseHandlerInterface) *HelloWebHandler {
	return &HelloWebHandler{
		ResponseHandler: rh,
	}
}

func (h *HelloWebHandler) Hello(w http.ResponseWriter, r *http.Request) {
	dto := usecase.NewHelloUseCase().Hello()
	h.ResponseHandler.Respond(w, http.StatusOK, dto)
}