package winget

import (
	"errors"
	"os/exec"
	"strings"
)

const winGetApp = "winget.exe"

func IsInstalled() bool {
	_, err := createCommand("--version").Output()
	return err == nil
}

func Run(params string) (WinGetResult, error) {
	params = strings.TrimSpace(params)
	execCommand := createCommand(params)
	outputBytes, err := execCommand.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode := convertExitCode(exitError.ExitCode())
			if exitCode == 1 {
				err = errors.New("WinGet app not installed")
			}
			return WinGetResult{ExitCode: exitCode}, err
		}
		return WinGetResult{}, err
	}
	processCall := createProcessCall(params)
	consoleOutput := string(outputBytes)
	return WinGetResult{processCall, consoleOutput, 0}, nil
}
