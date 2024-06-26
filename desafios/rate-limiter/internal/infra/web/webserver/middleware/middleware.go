package middleware

import (
	"errors"
	"net/http"
	"time"

	interfaces "github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/interface"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/middleware/strategy"
	rip "github.com/vikram1565/request-ip"
)

type RateLimitConfig struct {
	MaxReqIP int
	MaxReqToken int
	OperatingWindowMs int
}

type RateLimitMiddleware struct{
	rateLimitConfig RateLimitConfig
	strategy strategy.StrategyInterface
	httpResponse interfaces.WebResponseHandlerInterface
}

func NewRateLimitMiddleware(rateLimitConfig RateLimitConfig, strategy strategy.StrategyInterface, rh interfaces.WebResponseHandlerInterface) *RateLimitMiddleware{
	return &RateLimitMiddleware{
rateLimitConfig,
strategy,
rh,
	}
}

func(rlm *RateLimitMiddleware)RateLimit(next http.Handler) http.Handler{
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request){
		// TODO vai pra configs
		apiKey := r.Header.Get("API_KEY")
		userIp := rip.GetClientIP(r)

		result, err := rlm.check(apiKey, userIp)
		if err != nil{
			rlm.httpResponse.RespondWithError(w, http.StatusInternalServerError, errors.Join(errors.New("error at RateLimit normalization: "), err))
			return 
		}

		if !result.Result {
			rlm.httpResponse.RespondWithError(w, http.StatusTooManyRequests, errors.New("rate limit exceeded"))
			return
		}
		next.ServeHTTP(w,r)
	})
}

func(rlm *RateLimitMiddleware) check(apiKey, userIp string) (*strategy.RateLimitOutput, error) {
	var key string
	var limit int64
	duration := time.Duration(rlm.rateLimitConfig.OperatingWindowMs) * time.Millisecond

	if apiKey != "" {
		key = apiKey
		limit = int64(rlm.rateLimitConfig.MaxReqToken)
	} else {
		key = userIp
		limit = int64(rlm.rateLimitConfig.MaxReqIP)
	}

	strategyInput := &strategy.RateLimitInput{
		Key:      key,
		Limit:    limit,
		Duration: duration,
	}

	result, err := rlm.strategy.RateLimitStrategy(strategyInput)
	if err != nil {
		return nil, err
	}
	return result, nil
}