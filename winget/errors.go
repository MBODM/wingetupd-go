package winget

import (
	"fmt"
)

func createError(caller string, msg string) error {
	return fmt.Errorf("[winget.%s] %s", caller, msg)
}

func wrapError(caller string, msg string, err error) error {
	return fmt.Errorf("[winget.%s] %s: %w", caller, msg, err)
}
