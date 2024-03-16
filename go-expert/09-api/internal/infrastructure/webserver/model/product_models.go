package model

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
	"time"
)

type ProductModel struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductInput struct {
	Name  string
	Price float64
}
