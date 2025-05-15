package usecase_errors

import "fmt"

type UsecaseError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *UsecaseError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
}

func NewUsecaseError(code ErrorCode, msg string, err error) *UsecaseError {
	return &UsecaseError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}
