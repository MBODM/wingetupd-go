package winget

import (
	"fmt"
	"os/exec"
	"strings"
)

func createCommand(params string) *exec.Cmd {
	cmdParam := createProcessCall(params)
	return exec.Command("cmd", "/C", cmdParam)
}

func createProcessCall(winGetParams string) string {
	// To get a correct exit code (cause of the cmd /C workaround),
	// the complete WinGet call has to be inside one single string.
	processCall := fmt.Sprintf("%s %s", winGetApp, winGetParams)
	return strings.TrimSpace(processCall)
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
