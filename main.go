package main

import (
	"errors"
	"fmt"
	"os"

	"example.com/mbodm/wingetupd/app"
	"example.com/mbodm/wingetupd/errs"
	"example.com/mbodm/wingetupd/winget"
)

func main() {
	if winget.Exists() {
		result, err := winget.Run("search --exact --id Mozilla.Firefox")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result.ConsoleOutput)
	}
	os.Exit(1)
	fmt.Println()
	fmt.Printf("%s %s (by %s %s)", app.Name, app.Version, app.Author, app.Date)
	fmt.Println()
	fmt.Println()
	intro()
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

func intro() {
	err := errs.CreateLog()
	if err != nil {
		handleErrors(err)
		os.Exit(1)
	}
}

func exit(exitCode int) {
	errs.CloseLog()
	os.Exit(1)
}

func handleErrors(err error) {
	var expectedError *errs.ExpectedError
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
