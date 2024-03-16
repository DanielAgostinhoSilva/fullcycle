package exception

type UnauthorizedUserError struct {
	message string
}

func NewUnauthorizedUserError(message string) *UnauthorizedUserError {
	return &UnauthorizedUserError{
		message: message,
	}
}

func (e *UnauthorizedUserError) Error() string {
	return e.message
}
