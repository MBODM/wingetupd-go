package app

import "fmt"

type AppError struct {
	Msg string
	Err error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Errorf("%s %w", e.Msg, e.Err).Error()
	}
	return e.Msg
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(msg string, err error) *AppError {
	return &AppError{msg, err}
}

func WrapError(funcName string, innerError error) error {
	return fmt.Errorf("error in "+funcName+"() func: %w", innerError)
}

func EmptyStrArgError(funcName string) error {
	return NewAppError("Empty string argument in "+funcName+"() func", nil)
}
