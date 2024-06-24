package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	interfaces "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/interface"
	"github.com/go-redis/redis"
	rip "github.com/vikram1565/request-ip"
)

type RateLimiterConfig struct {
	Redis *redis.Client
	MaxRequestsPerIP int
	MaxRequestsPerToken int
	TimeWindowMilliseconds int
}

type RateLimiterMiddleware struct{
	rateLimiterConfig RateLimiterConfig
	httpResponse interfaces.WebResponseHandlerInterface
}

func NewRateLimiterMiddleware(rateLimiterConfig RateLimiterConfig, rh interfaces.WebResponseHandlerInterface) *RateLimiterMiddleware{
	return &RateLimiterMiddleware{
rateLimiterConfig,
rh,
	}
}

func(rlm *RateLimiterMiddleware)RateLimiter(next http.Handler) http.Handler{
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request){
		apiKey := r.Header.Get("API_KEY")
		userIp := rip.GetClientIP(r)

		_, err := rlm.check(context.Background(), apiKey, userIp)

		if err != nil{
			rlm.httpResponse.RespondWithError(w, http.StatusInternalServerError, errors.Join(errors.New("error RateLimiter normalization: "), err))
			return 
		}
		next.ServeHTTP(w,r)
	})
}

type RateLimiterInput struct {
	Key            string
	Limit    int64
	Duration time.Duration
}

type RateLimiterResult struct {
	Result    int
	Limit     int64
	Total     int64
	Remaining int64
	ExpiresAt time.Time
}

func(rlm *RateLimiterMiddleware) check(ctx context.Context, apiKey, userIp string) (*RateLimiterResult, error) {
	var key string
	var limit int64
	duration := time.Duration(rlm.rateLimiterConfig.TimeWindowMilliseconds) * time.Millisecond

	if apiKey != "" {
		key = apiKey
		limit = int64(rlm.rateLimiterConfig.MaxRequestsPerToken)
	} else {
		key = userIp
		limit = int64(rlm.rateLimiterConfig.MaxRequestsPerIP)
	}

	req := &RateLimiterInput{
		Key:      key,
		Limit:    limit,
		Duration: duration,
	}
	fmt.Println(req)

	result, err := rlm.Strategy.Check(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}