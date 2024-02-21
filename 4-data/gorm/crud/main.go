package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
	// gorm controla as acoes do db (cria created_at, updates_at e deleted_at)
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Name: "Test",
		Price: 100.0,
	})

	products := []Product{
	{
		Name: "Notebook",
		Price: 11230.0,
	},{
		Name: "Monitor",
		Price: 1200.99,
	},{
		Name: "Mouse",
		Price: 450.6,
	},
	}
	db.Create(&products)

	secondProduct := Product{}
	db.First(&secondProduct, 2)
	fmt.Println(secondProduct)

	db.First(&secondProduct, "name =?", "Notebook")
	fmt.Println(secondProduct)

	var allProducts []Product
	// acha tudo com paginacao (se remover volta tudo)
	db.Limit(2).Offset(3).Find(&allProducts)
	if len(allProducts)>0{
		fmt.Println(allProducts)
	}


	db.Where("price> ?", 500).Find(&allProducts)
	fmt.Println(allProducts)
	db.Where("name LIKE ?", "%book%").Find(&allProducts)
	fmt.Println(allProducts)

	updateProduct := Product{} 
	db.First(updateProduct, 3)
	updateProduct.Name = "New Monitor"
	db.Save(&updateProduct)
	// log update info
	db.First(updateProduct, 3)
	fmt.Println(updateProduct)

	db.Where("name LIKE ?", "%book%").Delete(&Product{})
	db.Delete(&updateProduct)


}
