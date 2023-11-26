package entity

type UserInterface interface {
	Create(user User) error
	FindByEmail(email string) (*User, error)
}
