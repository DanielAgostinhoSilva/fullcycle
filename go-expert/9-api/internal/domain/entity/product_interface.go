package entity

import "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"

type ProductInterface interface {
	Create(product *Product) error
	FindByID(ID entity.ID) (*Product, error)
	FindAll() ([]Product, error)
	DeleteById(ID entity.ID) error
}
