package controller

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/product"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ProductController struct {
	repository entity.ProductInterface
}

func NewProductController(repository entity.ProductInterface) *ProductController {
	return &ProductController{
		repository: repository,
	}
}

func (p *ProductController) FindAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	useCase := product.NewFindAllProductsUseCase(p.repository)
	pageProductOutput, err := useCase.Execute(product.ProductPage{Page: page, Size: size})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(pageProductOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (p *ProductController) DeleteById(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	useCase := product.NewDeleteProductByIdUseCase(p.repository)
	err := useCase.Execute(productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (p *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	useCase := product.NewFindProductByIdUseCase(p.repository)
	productOutput, err := useCase.Execute(productId)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		//return
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(productOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var productInput product.ProductDtoInputDTO
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createProductUseCase := product.NewCreateProductUseCase(p.repository)
	productOutput, err := createProductUseCase.Execute(productInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(productOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	var productInput product.ProductDtoInputDTO
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	useCase := product.NewUpdateProductByIdUseCase(p.repository)
	err = useCase.Execute(product.ProductUpdateInputDto{
		ProductId: productId,
		Name:      productInput.Name,
		Price:     productInput.Price,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p *ProductController) Router(router chi.Router) {
	router.Post("/", p.Create)
	router.Get("/{id}", p.FindById)
	router.Get("/", p.FindAll)
	router.Delete("/{id}", p.DeleteById)
	router.Put("/{id}", p.Update)
}

func (p *ProductController) Path() string {
	return "/v1/products"
}
