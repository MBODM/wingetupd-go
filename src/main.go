package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/errs"
)

func main() {
	title := fmt.Sprintf("%s %s (by %s %s)", app.Name, app.Version, app.Author, app.Date)
	fmt.Println(title)
	fmt.Println()
	result, err := app.Run()
	if err != nil {
		handleErrors(err)
		os.Exit(1)
	}
	if !result {
		os.Exit(1)
	}
	fmt.Println("Have a nice day.")
	os.Exit(0)
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
