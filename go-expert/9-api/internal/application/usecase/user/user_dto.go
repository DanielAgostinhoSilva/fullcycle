package user

type UserDtoOutput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDtoInput struct {
	Name     string
	Email    string
	Password string
}

type UserEmailDtoInput struct {
	Email string
}
