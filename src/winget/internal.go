package winget

import (
	"os/exec"
	"strings"
)

func createProcessCall(winGetParams string) string {
	// To get the correct exit code (cause of cmd /C workaround),
	// the complete WinGet call has to be inside a single string.
	// Also keep in mind: WinGet could be started without params.
	winGetParams = strings.TrimSpace(winGetParams)
	processCall := winGetApp + " " + winGetParams
	return strings.TrimSpace(processCall)
}

func createCommand(processCall string) *exec.Cmd {
	// Since Go is not able to exec Windows store apps,
	// workaround is to use cmd /C as additional layer.
	return exec.Command("cmd", "/C", processCall)
}

func convertExitCode(exitErrorExitCode int) int {
	// Go´s pkg does some weird stuff with the original exit code,
	// to result in a -1 value, as a marker for a stopped process.
	// Building the 2s complement is the opposite of what Go does.
	// So this converts Go´s value back to the original exit code.
	var tmp int32 = int32(exitErrorExitCode)
	var realExitCode int32 = -(^tmp + 1)
	return int(realExitCode)
}
