package entity

import (
	"errors"
	"time"

	"github.com/andremelinski/pos-goexpert/8-API/pkg/entity"
)

// erros que a aplicacao pode ter ao longo do desenvolvimento dos endpoints:
var (
	// buscar pelo id
	ErrRequiredID = errors.New("id is required")
	// id nao eh um uuid
	ErrInvalidID = errors.New("id is invalid")
	// preenchimento de campos obrigtorios
	ErrRequiredName = errors.New("name is required")
	ErrRequiredPrice = errors.New("price is required")
	// erro se user colocar um price <= 0
	ErrInvalidPrice = errors.New("price is invalid")
)
// struct utilizada para salvar no BD
type Product struct{
	ID entity.ID `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Created_at time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error){
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		Created_at: time.Now(),
	}

	err := product.Validate()
	if err != nil{
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	stringID := p.ID.String()
	if stringID =="" {
		return ErrRequiredID
	}
	if _, err := entity.ParseID(stringID); err !=nil {
		return ErrInvalidID
	}
	if p.Price == 0{
		return ErrRequiredPrice
	}
	if p.Price <0{
		return ErrInvalidPrice
	}
	if p.Name == ""{
		return ErrRequiredName
	}

	return nil
}