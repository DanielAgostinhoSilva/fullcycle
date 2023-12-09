package exception

type EntityInUsedError struct {
	message string
}

func NewEntityInUsed(message string) *EntityInUsedError {
	return &EntityInUsedError{
		message: message,
	}
}

func (e *EntityInUsedError) Error() string {
	return e.message
}
