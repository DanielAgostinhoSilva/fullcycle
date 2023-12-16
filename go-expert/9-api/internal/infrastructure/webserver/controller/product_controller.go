package controller

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"strconv"
)

type ProductController struct {
	findAllProductsUseCase   *product.FindAllProductsUseCase
	deleteProductByIdUseCase *product.DeleteProductByIdUseCase
	createProductUseCase     *product.CreateProductUseCase
	findProductByIdUseCase   *product.FindProductByIdUseCase
	updateProductByIdUseCase *product.UpdateProductByIdUseCase
}

func NewProductController(
	findAllProductsUseCase *product.FindAllProductsUseCase,
	deleteProductByIdUseCase *product.DeleteProductByIdUseCase,
	createProductUseCase *product.CreateProductUseCase,
	findProductByIdUseCase *product.FindProductByIdUseCase,
	updateProductByIdUseCase *product.UpdateProductByIdUseCase,
) *ProductController {
	return &ProductController{
		findAllProductsUseCase:   findAllProductsUseCase,
		deleteProductByIdUseCase: deleteProductByIdUseCase,
		createProductUseCase:     createProductUseCase,
		findProductByIdUseCase:   findProductByIdUseCase,
		updateProductByIdUseCase: updateProductByIdUseCase,
	}
}

func (p *ProductController) FindAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	pageProductOutput, err := p.findAllProductsUseCase.Execute(product.ProductPage{Page: page, Size: size})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(pageProductOutput)
	if err != nil {
		panic(err)
	}

}

func (p *ProductController) DeleteById(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	err := p.deleteProductByIdUseCase.Execute(productId)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (p *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	productOutput, err := p.findProductByIdUseCase.Execute(productId)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(productOutput)
	if err != nil {
		panic(err)
	}
}

func (p *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var productInput product.ProductDtoInputDTO
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		panic(err)
	}

	productOutput, err := p.createProductUseCase.Execute(productInput)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(productOutput)
	if err != nil {
		panic(err)
	}

}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	var productInput product.ProductDtoInputDTO
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		panic(err)
	}
	err = p.updateProductByIdUseCase.Execute(product.ProductUpdateInputDto{
		ProductId: productId,
		Name:      productInput.Name,
		Price:     productInput.Price,
	})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p *ProductController) Router(tokenAuth *jwtauth.JWTAuth) func(router chi.Router) {
	return func(router chi.Router) {
		router.Use(jwtauth.Verifier(tokenAuth))
		router.Use(jwtauth.Authenticator(tokenAuth))
		router.Post("/", p.Create)
		router.Get("/{id}", p.FindById)
		router.Get("/", p.FindAll)
		router.Delete("/{id}", p.DeleteById)
		router.Put("/{id}", p.Update)
	}
}

func (p *ProductController) Path() string {
	return "/v1/products"
}
