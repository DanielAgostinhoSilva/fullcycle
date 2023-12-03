package product

import (
	domain "github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
)

type CreateProductUseCase struct {
	Repository domain.ProductInterface
}

func NewCreateProductUseCase(repository domain.ProductInterface) *CreateProductUseCase {
	return &CreateProductUseCase{
		Repository: repository,
	}
}

func (c *CreateProductUseCase) Execute(input ProductDtoInputDTO) (*ProductOutPutDTO, error) {
	product, err := domain.NewProduct(input.Name, input.Price)
	if err != nil {
		return nil, err
	}
	err = c.Repository.Create(product)
	if err != nil {
		return nil, err
	}
	return &ProductOutPutDTO{
		ID:        product.ID.String(),
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
	}, err
}
