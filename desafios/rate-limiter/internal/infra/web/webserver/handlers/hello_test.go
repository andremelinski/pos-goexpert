package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/usecase"
	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/pkg/mock"
	"github.com/stretchr/testify/suite"
)

type HelloWebHandlerTestSuite struct{
	suite.Suite
	helloHandler *HelloWebHandler
}
func (suite *HelloWebHandlerTestSuite) SetupSuite() {
	mockUsecase := new(mock.MockUseCase)
	mockUsecase.On("Hello").Return(&usecase.HelloOuputDTO{
		Message: "hello",
	}).Once()

	webResponse := &WebResponseHandler{}
	suite.helloHandler = NewHelloWebHandler(webResponse ,mockUsecase )
}

// func (suite *HelloWebHandlerTestSuite) TearDownTest() {
//     // Clean up or teardown after tests
// }

func TestSuite(t *testing.T) {
	suite.Run(t, new(HelloWebHandlerTestSuite))
}

func  (suite *HelloWebHandlerTestSuite) TestHelloWebHandler_Hello() {
	// Criando uma requisição HTTP falsa
	req, err := http.NewRequest("GET", "/hello", nil)
	suite.Assert().NoError(err)
	// Criando um ResponseRecorder para simular a resposta HTTP
	rr := httptest.NewRecorder()
	
	// Chamando a função Hello do handler
	suite.helloHandler.Hello(rr, req)

	suite.Assert().Equal(http.StatusOK, rr.Code)
	suite.Assert().Equal( "{\"message\":\"hello\"}\n", rr.Body.String())
}
