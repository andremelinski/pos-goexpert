package interfaces

import "net/http"

type MiddlewareInterface interface {
	RateLimiter(next http.Handler) http.Handler
}