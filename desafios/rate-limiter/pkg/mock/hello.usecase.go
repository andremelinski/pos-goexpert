package mock

import (
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/usecase"
	"github.com/stretchr/testify/mock"
)
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) Hello() *usecase.HelloOuputDTO {
    args := m.Called()
	return args.Get(0).(*usecase.HelloOuputDTO)
}