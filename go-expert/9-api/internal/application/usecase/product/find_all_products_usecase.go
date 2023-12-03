package product

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
)

type FindAllProductsUseCase struct {
	repository entity.ProductInterface
}

func NewFindAllProductsUseCase(repository entity.ProductInterface) *FindAllProductsUseCase {
	return &FindAllProductsUseCase{
		repository: repository,
	}
}

func (f *FindAllProductsUseCase) Execute(productPage ProductPage) (*PageProductOutputDTO, error) {
	page := 0
	size := 10
	if productPage.Page > 0 {
		page = productPage.Page
	}
	if productPage.Size > 0 {
		size = productPage.Size
	}
	products, pageable, err := f.repository.FindAll(page, size, "asc")
	if err != nil {
		return nil, err
	}
	var pageProductOutputDTO PageProductOutputDTO
	for _, product := range products {
		productOutput := ProductOutPutDTO{
			ID:        product.ID.String(),
			Name:      product.Name,
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
		}

		pageProductOutputDTO.Content = append(pageProductOutputDTO.Content, productOutput)
	}
	pageProductOutputDTO.Pageable = *pageable
	return &pageProductOutputDTO, nil
}
