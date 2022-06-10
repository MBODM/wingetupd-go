package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/errs"
)

func init() {

}

func main() {
	exit(5)
	title := fmt.Sprintf("%s %s (by %s %s)", app.Name, app.Version, app.Author, app.Date)
	fmt.Println(title)
	fmt.Println()
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
	var w io.Writer = log.Writer()
	var file = w.(*os.File)
	file.Close()
	os.Exit(exitCode)
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
