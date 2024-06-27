package strategy

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/pkg/mock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

type StrategyRateLimitTestSuite struct{
	suite.Suite
	rateLimitStrategy *StrategyRateLimit
	rateLimitInput *database.RateLimitInput
	rateLimitOutput *database.RateLimitOutput
	strategyMock *mock.RedisMock
}

func (suite *StrategyRateLimitTestSuite) SetupSuite() {
	suite.rateLimitInput = &database.RateLimitInput{
		Key      : "key",
		Limit : 10,
		Duration : 500*time.Millisecond,
	}

	suite.rateLimitOutput = &database.RateLimitOutput{
		Result:    true,
		Limit:     5,
		Total:     5,
		Remaining: 10000,
		ExpiresAt: time.Now(),
	}
	suite.strategyMock = new(mock.RedisMock)
}

// func (suite *StrategyRateLimitTestSuite) TearDownTest() {
//     // Clean up or teardown after tests
// }

func TestSuite(t *testing.T) {
	suite.Run(t, new(StrategyRateLimitTestSuite))
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Throw_Error(){
	suite.strategyMock.On("Get", "key").Return(nil, errors.New("random error")).Once()

	suite.rateLimitStrategy = NewStrategyRateLimit(suite.strategyMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	fmt.Println(err)
	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Create_Limiter(){
	suite.strategyMock.On("Get", "key").Return(nil, redis.Nil).Once()
	suite.strategyMock.On("Create",  suite.rateLimitInput)
	suite.strategyMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	// suite.strategyMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.strategyMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(suite.rateLimitOutput, output)
	suite.NoError(err)
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Update_Limiter(){
	update := *suite.rateLimitOutput
	update.Total = update.Total -1
	suite.strategyMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	suite.strategyMock.On("Update",  "key", suite.rateLimitOutput, true).Once()
	suite.strategyMock.On("Get", "key").Return(&update, nil).Once()
	// suite.strategyMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.strategyMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(&update, output)
	suite.NoError(err)
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Create_Limiter_When_Time_Expired(){
	duration, _ := time.ParseDuration("-1.5h")
	
	result := *suite.rateLimitOutput
	result.ExpiresAt = time.Now().Add(duration)
	result.Total = 0
	
	suite.strategyMock.On("Get", "key").Return(&result, nil).Once()
	suite.strategyMock.On("Create",  suite.rateLimitInput).Once()
	suite.strategyMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	// suite.strategyMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.strategyMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(suite.rateLimitOutput, output)
	suite.NoError(err)
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Update_Limiter_When_Blocked(){
	duration, _ := time.ParseDuration("1.5h")
	
	result := *suite.rateLimitOutput
	result.ExpiresAt = time.Now().Add(duration)
	result.Total = 0
	
	suite.strategyMock.On("Get", "key").Return(&result, nil).Once()
	suite.strategyMock.On("Update",  "key", &result, false).Once()
	suite.strategyMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	// suite.strategyMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.strategyMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(suite.rateLimitOutput, output)
	suite.NoError(err)
}