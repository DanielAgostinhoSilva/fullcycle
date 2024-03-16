package mock

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database"
	pkgEntity "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (p *ProductRepositoryMock) Create(product *entity.Product) error {
	args := p.Called(product)
	return args.Error(0)
}

func (p *ProductRepositoryMock) FindByID(ID pkgEntity.ID) (*entity.Product, error) {
	args := p.Called(ID)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (p *ProductRepositoryMock) FindAll(page, limit int, sort string) ([]entity.Product, *database.Pageable, error) {
	args := p.Called(page, limit, sort)
	return args.Get(0).([]entity.Product), args.Get(1).(*database.Pageable), args.Error(2)
}

func (p *ProductRepositoryMock) DeleteById(ID pkgEntity.ID) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepositoryMock) Update(product *entity.Product) error {
	//TODO implement me
	panic("implement me")
}
