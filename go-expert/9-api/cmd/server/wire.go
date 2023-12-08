//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/product"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/webserver/controller"
	"github.com/google/wire"
)

var setProductRepository = wire.NewSet(
	database.NewProductRepository,
	wire.Bind(new(entity.ProductInterface), new(*database.ProductRepository)),
)

func InitializeProductController(db *sql.DB) *controller.ProductController {
	wire.Build(
		setProductRepository,
		product.NewCreateProductUseCase,
		product.NewDeleteProductByIdUseCase,
		product.NewFindAllProductsUseCase,
		product.NewFindProductByIdUseCase,
		product.NewUpdateProductByIdUseCase,
		controller.NewProductController,
	)
	return &controller.ProductController{}
}
