package winget

import (
	"errors"
	"os/exec"
	"strings"
)

func createProcessCall(winGetParams string) string {
	// To get the correct exit code (cause of cmd /C workaround),
	// the complete WinGet call has to be inside a single string.
	// Also keep in mind: WinGet could be started without params.
	winGetParams = strings.TrimSpace(winGetParams)
	processCall := WinGetApp + " " + winGetParams
	return strings.TrimSpace(processCall)
}

func createCommand(processCall string) *exec.Cmd {
	// Since Go is not able to exec Windows store apps,
	// workaround is to use cmd /C as additional layer.
	return exec.Command("cmd", "/C", processCall)
}

func handleExecError(execCaller string, execError error) (int, error) {
	// Since this function has to be used always directly after the exec.Cmd.Output() call,
	// not adding another useless error wrapping layer here. Using the exec caller instead.
	if execError == nil {
		return 0, nil
	}
	var exitError *exec.ExitError
	if errors.As(execError, &exitError) {
		exitCode := convertExitCode(exitError.ExitCode())
		if exitCode == 1 {
			// When using cmd /C workaround: An exit code of 1 means WinGet was not found.
			return 1, wrapError(execCaller, WinGetApp+" not found", execError)
		}
		// When using cmd /C workaround: An exit code != 1 is a real WinGet app exit code.
		return exitCode, nil
	}
	// When landing here: This means it is not an ExitError, but some other os/exec error.
	return 1, wrapError(execCaller, WinGetApp+" execution failed", execError)
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
