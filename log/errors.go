package log

import "fmt"

func createError(caller string, msg string) error {
	return fmt.Errorf("[config.%s] %s", caller, msg)
}

func wrapError(caller string, msg string, err error) error {
	return fmt.Errorf("[config.%s] %s: %w", caller, msg, err)
}
