package mock

import (
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/stretchr/testify/mock"
)

type StrategyMock struct {
	mock.Mock
}

// MockGet mocks the Get method
func (m *StrategyMock) RateLimitStrategy(input *database.RateLimitInput) (*database.RateLimitOutput, error) {
	args:= m.Called(input )
    if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*database.RateLimitOutput), args.Error(1)
}