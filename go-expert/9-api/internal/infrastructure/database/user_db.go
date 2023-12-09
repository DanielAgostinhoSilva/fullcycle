package database

import (
	"database/sql"
	"errors"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (p *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	stmt, err := p.Db.Prepare("SELECT * FROM user WHERE email = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &user, err
}

func (p *UserRepository) Create(user *entity.User) error {
	stmt, err := p.Db.Prepare("INSERT INTO user (id, name, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Id.String(), user.Name, user.Email, user.Password)
	return err
}
