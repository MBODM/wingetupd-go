package winget

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

const winGetApp = "winget.exe"

type WinGetResult struct {
	ProcessCall   string
	ConsoleOutput string
	ExitCode      int
}

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
	processCall := getProcessCall(params)
	consoleOutput := string(outputBytes)
	consoleOutput = removeProgressbarChars(consoleOutput)
	return WinGetResult{ProcessCall: processCall, ConsoleOutput: consoleOutput, ExitCode: 0}, nil
}

func createCommand(params string) *exec.Cmd {
	cmdParam := getProcessCall(params)
	return exec.Command("cmd", "/C", cmdParam)
}

func getProcessCall(winGetParams string) string {
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

func removeProgressbarChars(winGetOutput string) string {
	if strings.Contains(winGetOutput, "\b") {
		winGetOutput = strings.NewReplacer(
			"\b|", "",
			"\b/", "",
			"\b-", "",
			"\b\\", "",
			"\b", "",
		).Replace(winGetOutput)
	}
	winGetOutput = strings.Replace(winGetOutput, "\r", "", 1)
	winGetOutput = strings.TrimSpace(winGetOutput)
	return winGetOutput
}
