package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	interfaces "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/interface"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/middleware/strategy"
	rip "github.com/vikram1565/request-ip"
)

type RateLimiterConfig struct {
	MaxRequestsPerIP int
	MaxRequestsPerToken int
	TimeWindowMilliseconds int
}

type RateLimiterMiddleware struct{
	rateLimiterConfig RateLimiterConfig
	strategy strategy.StrategyInterface
	httpResponse interfaces.WebResponseHandlerInterface
}

func NewRateLimiterMiddleware(rateLimiterConfig RateLimiterConfig, strategy strategy.StrategyInterface, rh interfaces.WebResponseHandlerInterface) *RateLimiterMiddleware{
	return &RateLimiterMiddleware{
rateLimiterConfig,
strategy,
rh,
	}
}

func(rlm *RateLimiterMiddleware)RateLimiter(next http.Handler) http.Handler{
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request){
		apiKey := r.Header.Get("API_KEY")
		userIp := rip.GetClientIP(r)

		result, err := rlm.check(context.Background(), apiKey, userIp)
		if err != nil{
			rlm.httpResponse.RespondWithError(w, http.StatusInternalServerError, errors.Join(errors.New("error RateLimiter normalization: "), err))
			return 
		}

		if !result.Result {
			rlm.httpResponse.RespondWithError(w, http.StatusTooManyRequests, errors.New("rate limit exceeded"))
			return
		}
		next.ServeHTTP(w,r)
	})
}

func(rlm *RateLimiterMiddleware) check(ctx context.Context, apiKey, userIp string) (*strategy.RateLimiterOutput, error) {
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

	strategyInput := &strategy.RateLimitInput{
		Key:      key,
		Limit:    limit,
		Duration: duration,
	}

	result, err := rlm.strategy.RateLimitStrategy(ctx, strategyInput)
	if err != nil {
		return nil, err
	}

	return result, nil
}