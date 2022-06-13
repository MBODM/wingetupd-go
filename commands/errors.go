package commands

import "fmt"

func createError(caller, msg string) error {
	return fmt.Errorf("[commands.%s] %s", caller, msg)
}

func wrapError(caller, msg string, err error) error {
	return fmt.Errorf("[commands.%s] %s: %w", caller, msg, err)
}

func argIsEmptyStringError(caller, arg string) error {
	return createError(caller, fmt.Sprintf("argument '%s' is empty string", arg))
}

func argIsNilError(caller, arg string) error {
	return createError(caller, fmt.Sprintf("argument '%s' is nil", arg))
}

func runnerError(caller string, err error) error {
	return wrapError(caller, "given WinGetRunner returned error", err)
}

func parserError(caller string, err error) error {
	return wrapError(caller, "given WinGetListParser returned error", err)
}
