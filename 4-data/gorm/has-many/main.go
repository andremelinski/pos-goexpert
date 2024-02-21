package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	products := []Product{
	{
		Name: "Notebook",
		Price: 11230.0,
		CategoryID: category.ID,
	},{
		Name: "Monitor",
		Price: 1200.99,
		CategoryID: category.ID,
	},{
		Name: "Mouse",
		Price: 450.6,
		CategoryID: category.ID,
	},
	}
	db.Create(&products)
	
	//  get all Products
	var productsInfo []Product
	db.Preload("Category").Find(&productsInfo)

	for _, product := range productsInfo {
		fmt.Println(product.Name, ":", product.Category.Name)
	}

	//  get all Categories and their products
	var categoriesInfo []Category
	db.Preload("Products").Find(&categoriesInfo)

	for _, category := range categoriesInfo {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println( "products:",product.Name," ", product.Price)
		}
	}
}