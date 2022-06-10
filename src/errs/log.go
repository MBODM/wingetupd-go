package errs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Inside of log.Fatal() the osExit() func is called.
// If using the "log" package to write some log file,
// it is not clear to me, how/when the file shall be
// closed. Since i could not found any info about it,
// i decided to handle all of the log stuff this way.

var logFileOpened bool

func CreateLog() error {
	// Needs to be idempotent, cause of os.Create() func.
	if !logFileOpened {
		file, err := os.Create(GetLogFile())
		if err != nil {
			return NewExpectedError("Could not create log file", err)
		}
		log.SetOutput(file)
		logFileOpened = true
	}
	return nil
}

func CloseLog() error {
	// Needs to be idempotent, cause of file.Close() func.
	if logFileOpened {
		// Close the file manually, since using a typical "defer" is impossible here.
		// Therefore using a type assertion, since os.File also implements io.Writer.
		writer := log.Writer()
		file := writer.(*os.File)
		err := file.Close()
		if err != nil {
			return NewExpectedError("Could not close log file", err)
		}
		log.SetOutput(os.Stderr)
		logFileOpened = false
	}
	return nil
}

func GetLogFile() string {
	exe, err := os.Executable()
	if err != nil {
		// Since log file is not open yet, the log.Println(err) func,
		// inside the Fatal() func, will print to Stderr. This is ok.
		Fatal(err)
	}
	exe = filepath.Dir(exe)
	tempDir := os.TempDir()
	logFile := filepath.Join(tempDir, "wingetupd.log")
	logFile = filepath.Join(exe, "wingetupdlog.txt")
	return logFile

}

func Fatal(err error) {
	// If log file is already open, log.Println() logs to file.
	// If log file is not open yet, log.Println() logs to Stderr.
	log.Println(err)
	// No checks here, cause CloseLog() is idempotent.
	CloseLog()
	fmt.Println(strings.ToUpper("Fatal unexpected error occurred! Application terminated."))
	fmt.Println()
	fmt.Println("See log file for deatils:")
	fmt.Println(GetLogFile())
	fmt.Println()
	fmt.Println("Please contact the developer:")
	fmt.Println("https://github.com/mbodm")
	fmt.Println()
	fmt.Println("This should never happen. Sorry!")
	os.Exit(23)
}
