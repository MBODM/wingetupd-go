package commands

import "fmt"

func createError(caller, msg string) error {
	return fmt.Errorf("[commands.%s] %s", caller, msg)
}

func createArgIsNilError(caller, arg string) error {
	return createError(caller, fmt.Sprintf("argument '%s' is nil", arg))
}

func createRunnerError(caller string) error {
	return createError(caller, "given WinGetRunner returned nil")
}

func createParserError(caller string) error {
	return createError(caller, "given WinGetListParser returned nil")
}
