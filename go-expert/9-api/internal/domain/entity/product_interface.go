package entity

type ProductInterface interface {
	Create(product Product) error
}
