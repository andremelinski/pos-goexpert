package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Product{}, &Category{})

	// categories := []Category{
	// 	{Name: "Eletronicos"},
	// 	{Name: "Cozinha"},
	// 	{Name: "Banheiro"},
	// }
	// db.Create(&categories)

	// // create product
	// products := []Product{
	// {
	// 	Name: "alexa",
	// 	Price: 11230.0,
	// 	Categories: []Category{categories[0], categories[1]},
	// },{
	// 	Name: "Monitor",
	// 	Price: 1200.99,
	// 	Categories: []Category{categories[2], categories[1]},
	// },{
	// 	Name: "Panela",
	// 	Price: 450.6,
	// 	Categories: []Category{categories[1]},
	// },{
	// 	Name: "escova de dente",
	// 	Price: 50.6,
	// 	Categories: []Category{categories[0], categories[2]},
	// },
	// }
	// db.Create(&products)
	
	// get all Products
	var productsInfo []Product
	db.Preload("Category").Find(&productsInfo)

	for _, product := range productsInfo {
		fmt.Println(product.Name, " ", product.Price, ":")
		for _, category := range product.Categories {
			fmt.Println( "categories:",category.Name)
		}
	}
	fmt.Println()
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