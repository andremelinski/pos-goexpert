package strategy

import (
	"errors"
	"testing"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/pkg/mock"
	"github.com/stretchr/testify/suite"
)

type StrategyRateLimitTestSuite struct{
	suite.Suite
	rateLimitStrategy *StrategyRateLimit
	rateLimitInput *database.RateLimitInput
	strategyMock *mock.RedisMock
}

func (suite *StrategyRateLimitTestSuite) SetupSuite() {
	suite.rateLimitInput = &database.RateLimitInput{
		Key      : "key",
		Limit : 10,
		Duration : 500*time.Millisecond,
	}
	// &database.RateLimitOutput{
	// 	Result:    true,
	// 	Limit:     5,
	// 	Total:     5,
	// 	Remaining: 10000,
	// 	ExpiresAt: time.Now(),
	// }
	suite.strategyMock = new(mock.RedisMock)
	suite.strategyMock.On("Get", "key").Return(nil, errors.New("random error"))

	suite.strategyMock.On("Update", "key").Return()

}


func (suite *StrategyRateLimitTestSuite) TearDownTest() {
    // Clean up or teardown after tests
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(StrategyRateLimitTestSuite))
}

func (suite *StrategyRateLimitTestSuite)RateLimitMiddleware_Throw_Error(){

	suite.rateLimitStrategy = NewStrategyRateLimit(suite.strategyMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)

	suite.Empty(output)
	suite.EqualError(err, "random error ")
}