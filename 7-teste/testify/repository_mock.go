package tax

import "github.com/stretchr/testify/mock"

type TaxRepositoryMock struct {
	// simula os metodos que precisa
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax (tax float64) error {
	// valor fica registrado aqui
	args := m.Called(tax)
	// retorna o tipo que eh a response, se fosse int e nao error, seria args.Int
	return args.Error(0)
}