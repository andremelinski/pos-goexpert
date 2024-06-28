package handlers

import (
	"net/http"

	interfaces "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/interface"
)

type HelloWebHandler struct {
	ResponseHandler interfaces.WebResponseHandlerInterface
	HelloUseCase interfaces.HelloUseCaseInterface
}

func NewHelloWebHandler(rh interfaces.WebResponseHandlerInterface, us interfaces.HelloUseCaseInterface) *HelloWebHandler {
	return &HelloWebHandler{
		ResponseHandler: rh,
		HelloUseCase: us,
	}
}

func (h *HelloWebHandler) Hello(w http.ResponseWriter, r *http.Request) {
	dto := h.HelloUseCase.Hello()
	h.ResponseHandler.Respond(w, http.StatusOK, dto)
}