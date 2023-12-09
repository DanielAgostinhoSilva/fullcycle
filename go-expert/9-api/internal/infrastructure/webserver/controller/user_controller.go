package controller

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/application/usecase/user"
	webserver "github.com/DanielAgostinhoSilva/fullcycle/9-api/pkg/webserver/exception"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"time"
)

var (
	invalidUserOrPasswordErr = webserver.NewBadRequestError("Invalid user or password")
)

type UserController struct {
	createUserCase      *user.CreateUserUseCase
	findUserUseCae      *user.FindUserUseCase
	validateUserUseCase *user.ValidateUserUseCase
}

func NewUserController(
	createUserCase *user.CreateUserUseCase,
	findUserUseCae *user.FindUserUseCase,
	validateUserUseCase *user.ValidateUserUseCase,
) *UserController {
	return &UserController{
		createUserCase:      createUserCase,
		findUserUseCae:      findUserUseCae,
		validateUserUseCase: validateUserUseCase,
	}
}

func (u *UserController) GetJwtToken(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwExpiresIn").(int)
	var userLoginInput user.UserLoginInput
	err := json.NewDecoder(r.Body).Decode(&userLoginInput)
	if err != nil {
		panic(invalidUserOrPasswordErr)
	}
	userOutput, err := u.validateUserUseCase.Execute(userLoginInput)
	if err != nil {
		panic(err)
	}

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": userOutput.ID,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(accessToken)
	if err != nil {
		panic(err)
	}

}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userInput user.UserInput
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
	router.Get("/generate_token", u.GetJwtToken)
}

func (u *UserController) Path() string {
	return "/v1/users"
}
