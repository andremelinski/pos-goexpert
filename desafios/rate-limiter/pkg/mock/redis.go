package mock

import (
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/stretchr/testify/mock"
)

// MockRateLimitRepository is a mock implementation of RateLimitRepositoryInterface
type RedisMock struct {
	mock.Mock
}

func (m *RedisMock) Get(key string) (*database.RateLimitOutput, error) {
    args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*database.RateLimitOutput), args.Error(1)
}

// MockCreate mocks the Create method
func (m *RedisMock) Create(input *database.RateLimitInput) {
    m.Called(input )
}

// MockUpdate mocks the Update method
func (m *RedisMock) Update(key string, input *database.RateLimitOutput, result bool) {
    m.Called(key, input, result )
}
