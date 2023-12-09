package user

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database/exception"
)

type FindUserUseCase struct {
	userRepository entity.UserInterface
}

func NewFindUserUseCase(userRepository entity.UserInterface) *FindUserUseCase {
	return &FindUserUseCase{userRepository: userRepository}
}

func (f *FindUserUseCase) Execute(email string) (*UserOutput, error) {
	user, err := f.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, exception.NewEntityNotFound("user not found")
	}

	return &UserOutput{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
