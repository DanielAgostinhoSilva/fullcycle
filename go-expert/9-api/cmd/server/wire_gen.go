// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/product"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/user"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/infrastructure/webserver/controller"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeProductController(db *sql.DB) *controller.ProductController {
	productRepository := database.NewProductRepository(db)
	findAllProductsUseCase := product.NewFindAllProductsUseCase(productRepository)
	deleteProductByIdUseCase := product.NewDeleteProductByIdUseCase(productRepository)
	createProductUseCase := product.NewCreateProductUseCase(productRepository)
	findProductByIdUseCase := product.NewFindProductByIdUseCase(productRepository)
	updateProductByIdUseCase := product.NewUpdateProductByIdUseCase(productRepository)
	productController := controller.NewProductController(findAllProductsUseCase, deleteProductByIdUseCase, createProductUseCase, findProductByIdUseCase, updateProductByIdUseCase)
	return productController
}

func InitializeUserController(db *sql.DB) *controller.UserController {
	userRepository := database.NewUserRepository(db)
	createUserUseCase := user.NewCreateUserUseCase(userRepository)
	findUserUseCase := user.NewFindUserUseCase(userRepository)
	validateUserUseCase := user.NewValidateUserUseCase(userRepository)
	userController := controller.NewUserController(createUserUseCase, findUserUseCase, validateUserUseCase)
	return userController
}

// wire.go:

var setProductRepository = wire.NewSet(database.NewProductRepository, wire.Bind(new(entity.ProductInterface), new(*database.ProductRepository)))

var setUserRepository = wire.NewSet(database.NewUserRepository, wire.Bind(new(entity.UserInterface), new(*database.UserRepository)))