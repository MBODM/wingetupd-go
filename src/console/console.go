package console

import "fmt"

func ShowUsage(exeFile string, hideError bool) {
	if !hideError {
		fmt.Println("Error: Unknown parameter(s).")
		fmt.Println()
	}
	fmt.Println("Usage: {exe_file} [--no-log] [--no-confirm]")
	fmt.Println()
	fmt.Println("  --no-log      Don´t create log file (useful when running from a folder without write permissions)")
	fmt.Println("  --no-confirm  Don´t ask for update confirmation (useful for script integration)")
	fmt.Println()
	fmt.Println("For more information have a look at the GitHub page (https://github.com/MBODM/wingetupd")
}
