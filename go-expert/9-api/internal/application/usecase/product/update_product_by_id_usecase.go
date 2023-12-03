package product

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	pkgEntity "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
)

type ProductUpdateInputDto struct {
	ProductId string
	Name      string
	Price     float64
}

type UpdateProductByIdUseCase struct {
	repository entity.ProductInterface
}

func NewUpdateProductByIdUseCase(repository entity.ProductInterface) *UpdateProductByIdUseCase {
	return &UpdateProductByIdUseCase{
		repository: repository,
	}
}

func (u *UpdateProductByIdUseCase) Execute(productUpdateDto ProductUpdateInputDto) error {
	id, err := pkgEntity.ParseID(productUpdateDto.ProductId)
	if err != nil {
		return err
	}
	productFound, err := u.repository.FindByID(id)
	if err != nil {
		return err
	}
	productFound.Name = productUpdateDto.Name
	productFound.Price = productUpdateDto.Price
	return u.repository.Update(productFound)
}
