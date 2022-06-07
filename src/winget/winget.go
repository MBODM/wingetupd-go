package winget

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/mbodm/wingetupd-go/app"
)

const winGetApp = "winge.exe"

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
				return WinGetResult{}, app.NewAppError("WinGet app not installed", err)
			} else {
				return WinGetResult{processCall, "", exitCode}, nil
			}
		}
		return WinGetResult{}, app.WrapError("winget.Run", err)
	}
	consoleOutput := string(outputBytes)
	return WinGetResult{processCall, consoleOutput, 0}, nil
}
