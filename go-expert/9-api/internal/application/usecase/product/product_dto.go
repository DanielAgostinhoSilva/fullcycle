package product

import (
	pkgDatabase "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database"
	"time"
)

type ProductDtoInputDTO struct {
	Name  string
	Price float64
}

type ProductOutPutDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductPage struct {
	Page int
	Size int
}

type PageProductOutputDTO struct {
	Content  []ProductOutPutDTO
	Pageable pkgDatabase.Pageable
}
