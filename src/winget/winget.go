package winget

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/mbodm/wingetupd-go/eh"
)

const winGetApp = "winget.exe"

func IsInstalled() bool {
	execCommand := createCommand("--version")
	_, err := execCommand.Output()
	return err == nil
}

func Run(params string) (WinGetResult, error) {
	params = strings.TrimSpace(params)
	processCall := createProcessCall(params)
	execCommand := createCommand(params)
	outputBytes, err := execCommand.Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			exitCode := convertExitCode(exitError.ExitCode())
			if exitCode == 1 {
				return WinGetResult{}, eh.NewExpectedError("WinGet app not installed", err)
			} else {
				return WinGetResult{processCall, "", exitCode}, nil
			}
		}
		return WinGetResult{}, eh.WrapError("winget.Run", err)
	}
	consoleOutput := string(outputBytes)
	return WinGetResult{processCall, consoleOutput, 0}, nil
}
