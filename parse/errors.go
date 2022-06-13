package parse

import (
	"fmt"
)

func createError(caller string, msg string) error {
	return fmt.Errorf("[parse.%s] %s", caller, msg)
}
