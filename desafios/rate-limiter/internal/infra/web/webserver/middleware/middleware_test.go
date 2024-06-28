package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/pkg/mock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
)

type MiddlewareRateLimitTestSuite struct {
	suite.Suite
	rateLimitConfig RateLimitConfig
	strategyInput *database.RateLimitInput
	rateLimitMid *RateLimitMiddleware
	strategyMock *mock.StrategyMock
	router          *chi.Mux
	rateLimitInput *database.RateLimitInput
	rateLimitOutput *database.RateLimitOutput

}

func (suite *MiddlewareRateLimitTestSuite) SetupSuite() {
// net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]).String(),
	suite.router = chi.NewRouter()
	httpRespHandler := &handlers.WebResponseHandler{} 
	suite.strategyMock = new(mock.StrategyMock)

	rateLimitConfig := RateLimitConfig{
		MaxReqIP: 5,
		MaxReqToken: 10,
		OperatingWindowMs: 500,
	}

	suite.rateLimitInput = &database.RateLimitInput{Key:"127.0.0.1", Limit:5, Duration:500000000}

	suite.rateLimitOutput = &database.RateLimitOutput{
		Result:    true,
		Limit:     5,
		Total:     4,
		Remaining: 10000,
		ExpiresAt: time.Now(),
	}
	suite.rateLimitMid = NewRateLimitMiddleware(rateLimitConfig, suite.strategyMock, httpRespHandler)
}


func (suite *MiddlewareRateLimitTestSuite) TearDownTest() {
    // Clean up or teardown after tests
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareRateLimitTestSuite))
}

func (suite *MiddlewareRateLimitTestSuite)Test_Rate_Limit(){
	suite.router.Use(suite.rateLimitMid.RateLimit)
	suite.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello World!"}`))
	})
	
	server := httptest.NewServer(suite.router)
	
	suite.Run("Should allow request when rate limiter allows", func() {


		req, err := http.NewRequest("GET", server.URL, nil)
		suite.Assert().NoError(err)

		suite.strategyMock.On("RateLimitStrategy", suite.rateLimitInput).Return(suite.rateLimitOutput, nil).Once()

		resp, err := http.DefaultClient.Do(req)
		suite.Assert().NoError(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
		suite.Equal("5", resp.Header.Get("X-RateLimit-Limit"))
		suite.Equal("4", resp.Header.Get("X-RateLimit-Total"))
	})

	suite.Run("Should return too many requests", func() {
		badReqOutput := *suite.rateLimitOutput
		badReqOutput.Result = false
		badReqOutput.Total=0
		req, err := http.NewRequest("GET", server.URL, nil)
		suite.Assert().NoError(err)

		suite.strategyMock.On("RateLimitStrategy", suite.rateLimitInput).Return(&badReqOutput, nil).Once()

		resp, err := http.DefaultClient.Do(req)
		suite.Assert().NoError(err)

		suite.Equal(http.StatusTooManyRequests, resp.StatusCode)
		suite.Equal("5", resp.Header.Get("X-RateLimit-Limit"))
		suite.Equal("0", resp.Header.Get("X-RateLimit-Total"))
	})
}