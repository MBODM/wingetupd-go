package errs

// Named this "errs" instead of  "errors", cause using it
// everywhere (cross cutting concern) and therefore shall
// be a rather short name. And cause of "great" Go design,
// names like "error" or "errors" are not usable. Yee-Haw!

import "fmt"

func NewExpectedError(msg string, err error) error {
	return &ExpectedError{msg, err}
}

func WrapError(funcName string, err error) error {
	return fmt.Errorf("error in %s() func: %w", funcName, err)
}

func CreateError(funcName string, msg string) error {
	return fmt.Errorf("error in %s() func: %s", funcName, msg)
}

func ArgIsEmptyStringError(funcName string) error {
	return CreateError(funcName, "argument is empty string")
}

func ArgIsNilError(funcName string) error {
	return CreateError(funcName, "argument is nil")
}
