package user

import (
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/database/exception"
)

var (
	userNotFoundErr     = exception.NewEntityNotFound("user not found")
	unauthorizedUserErr = exception.NewUnauthorizedUserError("Invalid user or password")
)

type ValidateUserUseCase struct {
	userRepository entity.UserInterface
}

func NewValidateUserUseCase(userRepository entity.UserInterface) *ValidateUserUseCase {
	return &ValidateUserUseCase{userRepository: userRepository}
}

func (v *ValidateUserUseCase) Execute(input UserLoginInput) (*UserOutput, error) {
	email, err := v.userRepository.FindByEmail(input.Email)
	if err != nil || email == nil {
		return nil, userNotFoundErr
	}
	if !email.ValidatePassword(input.Password) {
		return nil, unauthorizedUserErr
	}
	return &UserOutput{
		ID:    email.Id.String(),
		Name:  email.Name,
		Email: email.Email,
	}, nil
}
