package user

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database/exception"
)

type CreateUserUseCase struct {
	userRepository entity.UserInterface
}

func NewCreateUserUseCase(userRepository entity.UserInterface) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}

func (c *CreateUserUseCase) Execute(input UserInput) error {
	userFound, err := c.userRepository.FindByEmail(input.Email)
	if err != nil {
		return err
	}
	if userFound != nil {
		return exception.NewEntityInUsed("user already exists")
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return err
	}
	return c.userRepository.Create(user)
}
