package usecase

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RateLimitUseCaseTestSuite struct{
	suite.Suite

}


// func (suite *RateLimitUseCaseTestSuite) TearDownTest() {
//     // Clean up or teardown after tests
// }

func TestSuite(t *testing.T) {
	suite.Run(t, new(RateLimitUseCaseTestSuite))
}

func(suite *RateLimitUseCaseTestSuite ) Hello_UseCase(){
	result := NewHelloUseCase().Hello()
	suite.Assert().Equal(&HelloOuputDTO{
		"hello",
	}, result)
}
