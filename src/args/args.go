package args

// I don´t use flag, because i don´t wanna allow -params instead of --params.
// And i wrote this code in half of the time i learn and workaround flag pkg.

import (
	"os"
)

const noConfirmArg = "--no-confirm"
const helpArg = "--help"

func Validate() bool {
	for _, arg := range getArgs() {
		if !isSupportedArg(arg) {
			return false
		}
	}
	return true
}

func NoConfirmExists() bool {
	return argExists(noConfirmArg)
}

func HelpExists() bool {
	return argExists(helpArg)
}

func isSupportedArg(arg string) bool {
	supportedArgs := []string{noConfirmArg, helpArg}
	for _, supportedArg := range supportedArgs {
		if arg == supportedArg {
			return true
		}
	}
	return false
}

func argExists(arg string) bool {
	for _, existingArg := range getArgs() {
		if arg == existingArg {
			return true
		}
	}
	return false
}

func getArgs() []string {
	argsWithoutExe := os.Args[1:]
	return argsWithoutExe
}

// Golang REALLY! is missing some functional and generic stuff.
// It´s rather hard, when you have to write such loops in 2022.
// In other languages we use "any()" or "filter()" since years.
// And that Go community is in some weird way even proud of it.
// But hey, Rust´s ownership/borrowchecker is happytime too. :)
