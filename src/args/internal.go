package args

import "os"

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
