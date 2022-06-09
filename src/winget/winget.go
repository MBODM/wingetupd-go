package winget

import (
	"errors"
	"os/exec"

	"github.com/mbodm/wingetupd-go/errs"
)

const winGetApp = "winget.exe"

func IsInstalled() bool {
	processCall := createProcessCall("--version")
	execCommand := createCommand(processCall)
	_, err := execCommand.Output()
	return err == nil
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
				return &WinGetResult{}, errs.NewExpectedError("WinGet not installed", err)
			} else {
				return &WinGetResult{processCall, "", exitCode}, nil
			}
		}
		return &WinGetResult{}, errs.WrapError("winget.Run", err)
	}
	consoleOutput := string(outputBytes)
	return &WinGetResult{processCall, consoleOutput, 0}, nil
}
