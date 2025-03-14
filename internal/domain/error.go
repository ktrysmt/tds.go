package domain

type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	InvalidInput ErrorType = "INVALID_INPUT"
	Internal     ErrorType = "INTERNAL"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
	}
}
