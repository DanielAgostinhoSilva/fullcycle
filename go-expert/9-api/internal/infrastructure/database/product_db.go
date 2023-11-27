package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	extenal "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	stmt, err := p.Db.Prepare("INSERT INTO product (id, name, price, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID.String(), product.Name, product.Price, product.CreatedAt)
	return err
}

func (p *ProductRepository) FindByID(id extenal.ID) (*entity.Product, error) {
	var product entity.Product
	stmt, err := p.Db.Prepare("SELECT * FROM product WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id.String()).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
	return &product, err
}

func (p *ProductRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	stmt, err := p.Db.Prepare("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepository) DeleteById(ID extenal.ID) error {
	stmt, err := p.Db.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ID)
	return err
}
