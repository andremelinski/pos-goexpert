package interfaces

import "net/http"

type MiddlewareInterface interface {
	RateLimit(next http.Handler) http.Handler
}