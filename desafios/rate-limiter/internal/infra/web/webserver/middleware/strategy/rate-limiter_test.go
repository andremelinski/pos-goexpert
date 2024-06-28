package strategy

import (
	"errors"
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
	redisMock *mock.RedisMock
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
	suite.redisMock = new(mock.RedisMock)
}

// func (suite *StrategyRateLimitTestSuite) TearDownTest() {
//     // Clean up or teardown after tests
// }

func TestSuite(t *testing.T) {
	suite.Run(t, new(StrategyRateLimitTestSuite))
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Throw_Error(){
	suite.redisMock.On("Get", "key").Return(nil, errors.New("random error")).Once()

	suite.rateLimitStrategy = NewStrategyRateLimit(suite.redisMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Empty(output)
	suite.EqualError(err, "random error")
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Create_Limiter(){
	suite.redisMock.On("Get", "key").Return(nil, redis.Nil).Once()
	suite.redisMock.On("Create",  suite.rateLimitInput)
	suite.redisMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	// suite.redisMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.redisMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(suite.rateLimitOutput, output)
	suite.NoError(err)
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Update_Limiter(){
	update := *suite.rateLimitOutput
	update.Total = update.Total -1
	suite.redisMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	suite.redisMock.On("Update",  "key", suite.rateLimitOutput, true).Once()
	suite.redisMock.On("Get", "key").Return(&update, nil).Once()
	// suite.redisMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.redisMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(&update, output)
	suite.NoError(err)
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Create_Limiter_When_Time_Expired(){
	duration, _ := time.ParseDuration("-1.5h")
	
	result := *suite.rateLimitOutput
	result.ExpiresAt = time.Now().Add(duration)
	result.Total = 0
	
	suite.redisMock.On("Get", "key").Return(&result, nil).Once()
	suite.redisMock.On("Create",  suite.rateLimitInput).Once()
	suite.redisMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	// suite.redisMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.redisMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(suite.rateLimitOutput, output)
	suite.NoError(err)
}

func (suite *StrategyRateLimitTestSuite)Test_RateLimitMiddleware_Update_Limiter_When_Blocked(){
	duration, _ := time.ParseDuration("1.5h")
	
	result := *suite.rateLimitOutput
	result.ExpiresAt = time.Now().Add(duration)
	result.Total = 0
	
	suite.redisMock.On("Get", "key").Return(&result, nil).Once()
	suite.redisMock.On("Update",  "key", &result, false).Once()
	suite.redisMock.On("Get", "key").Return(suite.rateLimitOutput, nil).Once()
	// suite.redisMock.On("Update", "key").Return()
	suite.rateLimitStrategy = NewStrategyRateLimit(suite.redisMock)

	output, err := suite.rateLimitStrategy.RateLimitStrategy(suite.rateLimitInput)
	suite.Equal(suite.rateLimitOutput, output)
	suite.NoError(err)
}