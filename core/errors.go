package core

import (
	"fmt"
)

type ExpectedError struct {
	Msg string
	Err error
}

// Implement error interface, so ExpectedError becomes a typical error.

func (e *ExpectedError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *ExpectedError) Unwrap() error {
	return e.Err
}

func createError(caller string, msg string) error {
	return fmt.Errorf("[core.%s] %s", caller, msg)
}

func wrapError(caller string, msg string, err error) error {
	return fmt.Errorf("[core.%s] %s: %w", caller, msg, err)
}

func chainError(caller string, err error) error {
	return wrapError(caller, "chained", err)
}

func createArgIsNilError(caller string, arg string) error {
	msg := fmt.Sprintf("argument '%s' is nil", arg)
	return createError(caller, msg)
}

func createArgIsEmptyStringError(caller string, arg string) error {
	msg := fmt.Sprintf("argument '%s' is empty string", arg)
	return createError(caller, msg)
}

func createExpectedError(caller string, msg string) error {
	msg = fmt.Sprintf("[core.%s] %s", caller, msg)
	return &ExpectedError{msg, nil}
}

func wrapIntoExpectedError(caller string, msg string, err error) error {
	msg = fmt.Sprintf("[core.%s] %s", caller, msg)
	return &ExpectedError{msg, err}
}

func unrecoverableError(panicHandler func(error), err error) {
	if panicHandler != nil {
		panicHandler(err)
	}
	panic(err)
}
