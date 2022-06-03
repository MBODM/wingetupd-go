package args

// I don´t use flag, because i don´t wanna allow -params instead of --params.
// And i wrote this code in half of the time i learn and workaround flag pkg.

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
