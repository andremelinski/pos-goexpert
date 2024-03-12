package db

import (
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}


func ProductInit(db *gorm.DB) *ProductDB{
	return &ProductDB{
		db,
	}
}

func (p *ProductDB) Create(product *entity.Product) error{
	newProduct, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		return err
	}

	return p.DB.Create(newProduct).Error
}
func (p *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error){
	products := []entity.Product{}
	var err error
	if sort != "" && sort != "asc" && sort != "desc"{
		sort = "asc"
	}
	if page >0 && limit >0 {
		err = p.DB.Limit(limit).Offset((page-1)*page).Order("created_at "+sort).Find(&products).Error
	}else{
		err = p.DB.Limit(100).Order("created_at "+sort).Find(&products).Error
	}

	return products, err
}
func (p *ProductDB) FindByID(id string) (*entity.Product, error){
	product := &entity.Product{}

	err := p.DB.Where("ID =?", id).First(product).Error
	if product.Name != "" {
		return nil, err
	}

	return product, nil
}
func (p *ProductDB) Update(product *entity.Product) error{
	_, err := p.FindByID(product.ID.String()); if err != nil {
		return err 
	}

	return p.DB.Save(product).Error
}
func (p *ProductDB) Delete(id string) error{
	_, err := p.FindByID(id); if err != nil {
		return err 
	}

	return p.DB.Delete(id).Error
}