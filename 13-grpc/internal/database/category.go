package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
    db *sql.DB
    ID string
    Name string
    Description string
}

func NewCategory(db *sql.DB) *Category {
    return &Category{db: db}
}
func (c *Category) Create(name, description string) (Category, error){
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) LoadAll() ([]Category, error){
	categoryRows, err := c.db.Query("SELECT * from categories;")
	if err != nil {
		return nil, err
	}

	catgories := []Category{}
	for categoryRows.Next() {
		product := Category{}
		err = categoryRows.Scan(&product.ID, &product.Name, &product.Description)
		if err != nil{
			return nil, err
		}
		catgories = append(catgories, product)
	}

	// inserindo na memoria em que o object vazio product esta

	defer categoryRows.Close()

	return catgories, nil
}

func (c *Category) FindCategoryByCourseId(coursrId string) (Category, error){
	category := Category{}
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", coursrId).Scan(&category.ID, &category.Name, &category.Description);

	if err != nil {
		return Category{}, err
	}

	return category, nil

}