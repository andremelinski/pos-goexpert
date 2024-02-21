package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	db.Ping()
	
	if err != nil {
		panic(err)
	}

	p1 := NewProduct("fone", 50)
	err = InsertProductDB(db, p1)
	
	p1.Price = 19.9
	UpdateProductDB(db, p1)

	if err != nil {
		panic(err)
	}

	foundProduct, err := getProductByID(db, p1.ID)

	fmt.Println(foundProduct)
	if err != nil {
		panic(err)
	}

	foundProducts, err := getProducts(db)

	fmt.Println(foundProducts)
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
// 
func InsertProductDB(db *sql.DB, product *Product) error {
	statement, err := db.Prepare("INSERT INTO products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(product.ID, product.Name, product.Price)

	defer statement.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductDB(db *sql.DB, product *Product) error {
	statement, err := db.Prepare("UPDATE products set name =?, price =? where id=?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(product.Name, product.Price, product.ID)

	defer statement.Close()
	if err != nil {
		return err
	}
	return nil
}

func getProductByID(db *sql.DB, id string) (*Product, error) {
	statement, err := db.Prepare("SELECT * from products where id=?")
	if err != nil {
		return nil, err
	}

	product := Product{}
	// inserindo na memoria em que o object vazio product esta
	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	defer statement.Close()
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func getProducts(db *sql.DB) ([]Product, error) {
	productRows, err := db.Query("SELECT * from products")
	if err != nil {
		return nil, err
	}

	products := []Product{}
	for productRows.Next() {
		product := Product{}
		err = productRows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}

	// inserindo na memoria em que o object vazio product esta

	defer productRows.Close()

	if err != nil {
		return nil, err
	}
	return products, nil
}

func DeleteProductByID(db *sql.DB, id string) error {
	statement, err := db.Prepare("delete from products where id=?")
	if err != nil {
		return nil
	}

	_, err = statement.Exec(id)

	defer statement.Close()
	if err != nil {
		return err
	}
	return nil
}