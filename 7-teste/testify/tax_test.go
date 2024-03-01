package tax_test

import (
	"errors"
	"testing"

	tax "github.com/andremelinski/pos-goexpert/7-test/testify"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectResult := 5.0

	result, err := tax.CalculateTax(amount)
	assert.Error(t, err, "deu erro")
	assert.Contains(t, err.Error(), "erro")
	assert.Equal(t,  expectResult, result)
}

func TestCalculateAndSave(t *testing.T){
	repository := &tax.TaxRepositoryMock{}
	// quando o saveTax receber 10.0 ele retorna nil como error
	repository.On("SaveTax", 10.00).Return(nil) //.Once or Twice() chama 1 ou 2 vezes a funcao
	// mockando o erro que iria vir do db
	repository.On("SaveTax", 0.0).Return(errors.New("saving error"))

	err := tax.CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = tax.CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "saving error")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
	
}