package apperror

import "fmt"

type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	httpStatus int    `json:"-"`
}

func NewError(code string, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		httpStatus: 400,
	}
}

func (e *Error) WithHttpStatus(httpStatus int) *Error {
	e.httpStatus = httpStatus
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *Error) HttpStatus() int {
	return e.httpStatus
}

func IsError(err error) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}

	return nil
}
