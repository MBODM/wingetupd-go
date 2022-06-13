package main

import (
	"errors"
	"fmt"
	"os"

	"example.com/mbodm/wingetupd/app"
	"example.com/mbodm/wingetupd/logging"
)

func main() {
	fmt.Println()
	fmt.Printf("%s %s (by %s %s)", app.Name, app.Version, app.Author, app.Date)
	fmt.Println()
	fmt.Println()
	if err := logging.CreateLog(); err != nil {
		fmt.Println("Error: Could not create log file.")
		os.Exit(1)
	}
	result, err := app.Run()
	if err != nil {
		handleErrors(err)
		exit(1)
	}
	if !result {
		exit(1)
	}
	fmt.Println("Have a nice day.")
	exit(0)
}

func exit(exitCode int) {
	logging.CloseLog()
	os.Exit(1)
}

func handleErrors(err error) {
	var expectedError *app.ExpectedError
	if errors.As(err, &expectedError) {
		// Need this, in case of STRG+C was pressed in update confirmation.
		if expectedError.Msg == "STRG+C" {
			// Do nothing and just quit without msg, if STRG+C was pressed.
		} else {
			fmt.Println("Error: " + expectedError.Msg + ".")
		}
	} else {
		fmt.Println("Unexpected error(s) occurred:", err)
	}
}
