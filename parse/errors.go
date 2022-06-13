package parse

import (
	"fmt"
)

func createError(caller, msg string) error {
	return fmt.Errorf("[parse.%s] %s", caller, msg)
}
