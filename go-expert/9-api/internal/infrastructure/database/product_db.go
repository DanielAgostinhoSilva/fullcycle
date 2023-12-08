package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database"
	extenal "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/entity"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (p *ProductRepository) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID)
	if err != nil {
		return err
	}
	stmt, err := p.Db.Prepare("UPDATE product SET name = ?, price = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Name, product.Price)
	return err
}

func (p *ProductRepository) Create(product *entity.Product) error {
	stmt, err := p.Db.Prepare("INSERT INTO product (id, name, price, created_at) VALUES (?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID.String(), product.Name, product.Price, product.CreatedAt)
	return err
}

func (p *ProductRepository) FindByID(id extenal.ID) (*entity.Product, error) {
	var product entity.Product
	stmt, err := p.Db.Prepare("SELECT * FROM product WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id.String()).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]entity.Product, *database.Pageable, error) {
	var products []entity.Product
	rows, pageable, err := database.QueryWithPagination(p.Db, "SELECT * FROM product ORDER BY name "+sort, page, limit)
	if err != nil {
		return nil, nil, err
	}
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, nil, err
		}
		products = append(products, product)
	}
	return products, pageable, nil
}

func (p *ProductRepository) DeleteById(ID extenal.ID) error {
	stmt, err := p.Db.Prepare("DELETE FROM product WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ID)
	return err
}
