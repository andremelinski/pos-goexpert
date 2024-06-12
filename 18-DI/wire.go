//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/andremelinski/pos-goexpert/18-DI/product"
	"github.com/google/wire"
)

// Passa todo mundo do usecase que precisa dentro do build e ele vai gerir a ordem de injecao de dependencia
// annotation: uma especie de "tags" que ao rodar o projeto, wire verifica se existe e exedcuta o comando.

// pensando que vc pode ser diversos usecases e os "New" iriam se repetir, vc pode fazer datasets para agrupar essas partes repitidas

var serRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	// toda vez que ver essa interface vc troca pelo product.NewProductRepositor
	wire.Bind(new(product.IProductRepository), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB)*product.ProductUseCase{
	wire.Build(
		// product.NewProductRepository,
		// sybstitui
		serRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}