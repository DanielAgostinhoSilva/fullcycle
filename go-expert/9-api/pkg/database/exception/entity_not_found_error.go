package exception

type EntityNotFoundError struct {
	message string
}

func NewEntityNotFound(message string) *EntityNotFoundError {
	return &EntityNotFoundError{
		message: message,
	}
}

func (e *EntityNotFoundError) Error() string {
	return e.message
}
