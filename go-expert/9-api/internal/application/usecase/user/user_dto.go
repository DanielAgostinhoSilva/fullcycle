package user

type UserOutput struct {
	ID    string `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Name     string
	Email    string
	Password string
}

type UserEmailInput struct {
	Email string
}

type UserLoginInput struct {
	Email    string
	Password string
}
