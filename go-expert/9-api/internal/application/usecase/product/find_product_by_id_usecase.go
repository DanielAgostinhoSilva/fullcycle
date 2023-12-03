package product

import (
	"fmt"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database/exception"
	pkgEntity "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
)

type FindProductByIdUseCase struct {
	repository entity.ProductInterface
}

func NewFindProductByIdUseCase(repository entity.ProductInterface) *FindProductByIdUseCase {
	return &FindProductByIdUseCase{
		repository: repository,
	}
}

func (f *FindProductByIdUseCase) Execute(productId string) (*ProductOutPutDTO, error) {
	id, err := pkgEntity.ParseID(productId)
	if err != nil {
		return nil, err
	}
	productFound, err := f.repository.FindByID(id)
	if err != nil {
		return nil, exception.NewEntityNotFound(fmt.Sprintf("Product not found with id %s", id))
	}
	return &ProductOutPutDTO{
		ID:        productFound.ID.String(),
		Name:      productFound.Name,
		Price:     productFound.Price,
		CreatedAt: productFound.CreatedAt,
	}, nil
}
