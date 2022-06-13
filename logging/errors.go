package logging

import "fmt"

func wrapError(caller string, msg string, err error) error {
	return fmt.Errorf("[log.%s] %s: %w", caller, msg, err)
}
