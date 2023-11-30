package entity

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
)

type ProductInterface interface {
	Create(product *Product) error
	FindByID(ID entity.ID) (*Product, error)
	FindAll(page, limit int, sort string) ([]Product, *database.Pageable, error)
	DeleteById(ID entity.ID) error
	Update(product *Product) error
}
