package middleware

import (
	"testing"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/web/webserver/http"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/pkg/mock"
	"github.com/stretchr/testify/suite"
)

type MiddlewareRateLimitTestSuite struct {
	suite.Suite
	rateLimitConfig RateLimitConfig
	createTestData *database.RateLimitInput
	rateLimitMid *RateLimitMiddleware
}

func (suite *MiddlewareRateLimitTestSuite) SetupSuite() {

	suite.createTestData = &database.RateLimitInput{
		Key: "key",
        Limit: 10,
    	Duration: 5000*time.Millisecond,
	}

	strategyMock := new(mock.StrategyMock)
	responseHandler := &http.WebResponseHandler{}

	rateLimitConfig := RateLimitConfig{
		MaxReqIP: 5,
		MaxReqToken: 10,
		OperatingWindowMs: 500,
	}
	suite.rateLimitMid = NewRateLimitMiddleware(rateLimitConfig, strategyMock, responseHandler)
}


func (suite *MiddlewareRateLimitTestSuite) TearDownTest() {
    // Clean up or teardown after tests
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareRateLimitTestSuite))
}

func (suite *MiddlewareRateLimitTestSuite)Test_Rate_Limit(){
}