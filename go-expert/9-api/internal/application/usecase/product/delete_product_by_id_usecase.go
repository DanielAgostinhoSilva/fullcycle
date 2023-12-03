package product

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	pkgEntity "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
)

type DeleteProductByIdUseCase struct {
	repository entity.ProductInterface
}

func NewDeleteProductByIdUseCase(repository entity.ProductInterface) *DeleteProductByIdUseCase {
	return &DeleteProductByIdUseCase{
		repository: repository,
	}
}

func (f *DeleteProductByIdUseCase) Execute(productId string) error {
	id, err := pkgEntity.ParseID(productId)
	if err != nil {
		return err
	}
	return f.repository.DeleteById(id)
}
