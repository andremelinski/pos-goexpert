package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://stackoverflow.com/questions/73661393/has-one-vs-belongs-to-in-gorm
type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	var products []Product
	fmt.Println(category)
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}