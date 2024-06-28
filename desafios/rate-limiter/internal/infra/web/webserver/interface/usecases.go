package interfaces

import (
	"net/http"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/usecase"
)

type WebHandlerInterface interface {
	Hello(w http.ResponseWriter, r *http.Request)
}

type HelloUseCaseInterface interface {
	Hello() *usecase.HelloOuputDTO
}