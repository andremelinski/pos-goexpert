package entity_test

import (
	"testing"

	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/stretchr/testify/assert"
)


func TestNewProduct(t *testing.T){
	product, err := entity.NewProduct("Product1", 1.99)
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Product1")
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Price, 1.99)
}

func TestProduct_ValidateProduct_CorrectPayload(t *testing.T){
	product, err := entity.NewProduct("Product1", 1.99)
	assert.Nil(t, err)
	
	err = product.Validate()
	assert.Nil(t, err)
}
func TestProduct_ValidateProduct_NameIsRequired(t *testing.T){
	product, err := entity.NewProduct("", 1.99)
	assert.Nil(t, product)
	
	assert.Equal(t, entity.ErrRequiredName, err)
}
func TestProduct_ValidateProduct_PriceLowerThanZero(t *testing.T){
	product, err := entity.NewProduct("Product1", -1.99)
	assert.Nil(t, product)
	assert.Equal(t, entity.ErrInvalidPrice,err)
}
func TestProduct_ValidateProduct_WithoutPrice(t *testing.T){
	product, err := entity.NewProduct("Product1", 0)
	assert.Nil(t, product)
	assert.Equal(t, entity.ErrRequiredPrice,err)
}