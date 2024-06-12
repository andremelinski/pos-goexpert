package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// ingestao de dependencias. DI -> um facilitador que resolve esse encadeamento de dependencia, fazendo com que possa ser chamado o usecase diretamente ser configurar manualmente todas essas dependencias
func main(){
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		panic(err)
	}

	// // injeta db no repo
	// productRepo := product.NewProductRepository(db)
	// // injeta repo no usecase
	// usecase := product.NewProductUseCase(productRepo)

	usecase := NewUseCase(db)


	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}