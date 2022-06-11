package winget

import (
	"errors"
	"fmt"
	"os/exec"
)

const winGetApp = "winget.exe"

func Exists() bool {
	processCall := createProcessCall("--version")
	execCommand := createCommand(processCall)
	outputBytes, err := execCommand.Output()
	if err != nil {
		// Not differing here, or it ends in complex error handling for user.
		// If some "cmd /C winget.exe --version" not results in an exit code
		// of 0 and in returning some text, it is declared non-existing here.
		return false
	}
	success := execCommand.ProcessState.ExitCode() == 0
	output := string(outputBytes)
	return success && output != ""
}

func Run(params string) (*WinGetResult, error) {
	processCall := createProcessCall(params)
	execCommand := createCommand(processCall)
	outputBytes, err := execCommand.Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			exitCode := convertExitCode(exitError.ExitCode())
			if exitCode == 1 {
				// When using cmd /C workaround: An exit code of 1 means WinGet was not found.
				return nil, fmt.Errorf("[winget.Run] %s not found: %w", winGetApp, err)
			}
			// When using cmd /C workaround: An exit code != 1 is a real WinGet app exit code.
			return &WinGetResult{processCall, "", exitCode}, nil
		}
		// When landing here: This means it is not an ExitError, but some other os/exec error.
		return nil, fmt.Errorf("[winget.Run] %s execution failed: %w", winGetApp, err)
	}
	consoleOutput := string(outputBytes)
	return &WinGetResult{processCall, consoleOutput, 0}, nil
}
