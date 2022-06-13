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

func createError(caller, msg string) error {
	return fmt.Errorf("[core.%s] %s", caller, msg)
}

func createArgIsNilError(caller, arg string) error {
	return createError(caller, fmt.Sprintf("argument '%s' is nil", arg))
}

func createArgIsEmptyStringError(caller, arg string) error {
	return createError(caller, fmt.Sprintf("argument '%s' is empty string", arg))
}

func wrapError(caller, msg string, err error) error {
	return fmt.Errorf("[core.%s] %s: %w", caller, msg, err)
}

func chainError(caller string, err error) error {
	return wrapError(caller, "chained", err)
}

func createExpectedError(caller, msg string) error {
	return &ExpectedError{fmt.Sprintf("[core.%s] %s", caller, msg), nil}
}

func wrapIntoExpectedError(caller, msg string, err error) error {
	// Not using chained text here, cause Error() does this.
	e := createExpectedError(caller, msg)
	e.(*ExpectedError).Err = err
	return e
}

// func unrecoverableError(panicHandler func(error), err error) {
// 	if panicHandler != nil {
// 		panicHandler(err)
// 	}
// 	panic(err)
// }
