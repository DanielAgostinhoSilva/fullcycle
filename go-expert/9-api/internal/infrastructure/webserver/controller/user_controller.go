package controller

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/user"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserController struct {
	createUserCase *user.CreateUserUseCase
	findUserUseCae *user.FindUserUseCase
}

func NewUserController(createUserCase *user.CreateUserUseCase, findUserUseCae *user.FindUserUseCase) *UserController {
	return &UserController{createUserCase: createUserCase, findUserUseCae: findUserUseCae}
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userInput user.UserDtoInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		panic(err)
	}

	err = u.createUserCase.Execute(userInput)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err != nil {
		panic(err)
	}

}

func (u *UserController) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	emailOutput, err := u.findUserUseCae.Execute(email)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(emailOutput)
	if err != nil {
		panic(err)
	}
}

func (u *UserController) Router(router chi.Router) {
	router.Post("/", u.Create)
	router.Get("/{email}", u.FindByEmail)
}

func (u *UserController) Path() string {
	return "/v1/users"
}
