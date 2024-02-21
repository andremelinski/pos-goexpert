package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1:1 relation

type Category struct {
	ID    int `gorm:"primarykey"`
	Name string
}

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
	CategoryID int
	Category
	SerialNumber
	// gorm controla as acoes do db (cria created_at, updates_at e deleted_at)
	gorm.Model
}

type SerialNumber struct {
	gorm.Model
	ID    int `gorm:"primarykey"`
	Number string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	// create category
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

	// Create Serial Number

	db.Create(&SerialNumber{
		Number: "123456",
		ProductID: products[0].ID,
	})

	//  get all
	var productsInfo []Product
	db.Preload("Category").Preload("SerialNumber").Find(&productsInfo)

	for _, product := range productsInfo {
		fmt.Println(product)
	}
}