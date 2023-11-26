package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
)

type ProductRepository struct {
	Db *sql.DB
}

func (p *ProductRepository) Create(product entity.Product) error {
	stmt, err := p.Db.Prepare("INSERT INTO product (id, name, price, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID.String(), product.Name, product.Price, product.CreatedAt)
	return err
}
