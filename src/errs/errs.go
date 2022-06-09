package errs

// Named this "errs" instead of  "errors", cause using it
// everywhere (cross cutting concern) and therefore shall
// be a rather short name. And cause of "great" Go design,
// names like "error" or "errors" are not usable. Yee-Haw!

import "fmt"

type ExpectedError struct {
	Msg string
	Err error
}

func (e *ExpectedError) Error() string {
	if e.Err != nil {
		return fmt.Errorf("%s %w", e.Msg, e.Err).Error()
	}
	return e.Msg
}

func (e *ExpectedError) Unwrap() error {
	return e.Err
}

func NewExpectedError(msg string, err error) *ExpectedError {
	return &ExpectedError{msg, err}
}

func WrapError(funcName string, innerError error) error {
	return fmt.Errorf("error in %s() func: %w", funcName, innerError)
}

func ArgIsEmptyStringError(funcName string) error {
	return fmt.Errorf("empty string argument in " + funcName + "() func")
}

func ArgIsNilError(funcName string) error {
	return fmt.Errorf("nil argument in " + funcName + "() func")
}
