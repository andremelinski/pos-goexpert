package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
// lock otimista: utliza versionamento para verificar se a info foi atualizada ou nao a medida que a mudanca foi solicitada 
// Ex: manda atualizar o nome e outras coisas, se mandar da 1 select e o nome ainda estiver na v1, vc salva tudo, caso mude a versao
// antes de ter finalizado seu processo, ele deve recomecar. usa lock quando tem poucas concorrencias e precisa de muitas transacoes
// lock pessimista: locka a linha da tabela. Ng consegue mexer enquanto o processo nao for finalizado. Processo mais demorado mas mais seguro
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Product{}, &Category{})

	//lock pessimista: select * from products where id=1 FOR UPDATE
	transaction := db.Begin()
	category := Category{}
	// lock da linha
	err = transaction.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error
	if err != nil{
		panic(err)
	}
	// categories inicial: eletronicos
	category.Name = "update eletronicos"
	// Save updates value in database. If value doesn't contain a matching primary key, value is inserted
	transaction.Debug().Save(&category)
	//  libera a tabela e commita as mudancas
	transaction.Commit()
}