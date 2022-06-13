package app

import (
	"fmt"
)

type ExpectedError struct {
	Msg string
	Err error
}

func (e *ExpectedError) Error() string {
	// Implement error interface the same way golib does.
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *ExpectedError) Unwrap() error {
	return e.Err
}

func createExpectedError(caller string, msg string) error {
	return &ExpectedError{fmt.Sprintf("[app.%s] %s", caller, msg), nil}
}

func wrapErrorIntoExpectedError(caller string, msg string, err error) error {
	// Not using chained text handling here (": "), cause Error() does this.
	e := createExpectedError(caller, msg)
	e.(*ExpectedError).Err = err
	return e
}

func createError(caller string, msg string) error {
	return fmt.Errorf("[app.%s] %s", caller, msg)
}

func wrapError(caller string, msg string, err error) error {
	return fmt.Errorf("[app.%s] %s: %w", caller, msg, err)
}

func chainError(caller string, err error) error {
	return wrapError(caller, "chained", err)
}

func argIsNilError(caller string, arg string) error {
	return createError(caller, fmt.Sprintf("argument '%s' is nil", arg))
}

func notInitializedError(caller string) error {
	return createError(caller, "core is not initialized")
}

// func unrecoverableError(panicHandler func(error), err error) {
// 	if panicHandler != nil {
// 		panicHandler(err)
// 	}
// 	panic(err)
// }
