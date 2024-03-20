package db

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("andre", 123.7)
	assert.NoError(t, err)

	productDB := ProductInitDB(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFinalAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := ProductInitDB(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindByID(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	newProduct, err := entity.NewProduct("Product 100", rand.Float64()*100)
	assert.NoError(t, err)
	db.Create(newProduct)

	productDB := ProductInitDB(db)
	foundProduct, err := productDB.FindByID(newProduct.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 100", foundProduct.Name)
}

func TestUpdate(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	newProduct, err := entity.NewProduct("Product 100", rand.Float64()*100)
	assert.NoError(t, err)
	db.Create(newProduct)

	productDB := ProductInitDB(db)
	newProduct.Price = 5.99
	err = productDB.Update(newProduct)
	assert.NoError(t, err)

	product := entity.Product{}
	err = db.First(&product, "id =?", newProduct.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, product.Price, 5.99)
	assert.Equal(t, product.ID,newProduct.ID)
	assert.Equal(t, "Product 100",newProduct.Name)
}

func TestDelete(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := ProductInitDB(db)

	newProduct, err := entity.NewProduct("Product 100", rand.Float64()*100)
	assert.NoError(t, err)
	db.Create(newProduct)

	err = productDB.Delete(newProduct.ID.String())
	assert.NoError(t, err)

	product := entity.Product{}
	err = db.First(&product, "id =?", newProduct.ID.String()).Error
	assert.Error(t, err)
}