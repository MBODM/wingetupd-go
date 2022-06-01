package winget

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
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
	if params == "" {
		return WinGetResult{}, errors.New("invalid argument")
	}
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
	return fmt.Sprintf("%s %s", winGetApp, winGetParams)
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
	// Some "double-escape" via backticks AND double-backslash (`\\`) is a must,
	// because the regexp.MustCompile() string argument needs to be escaped too.
	progressBarChars := []string{`\\u0008`, "|", "/", "-", `\\`}
	for _, progressBarChar := range progressBarChars {
		winGetOutput = removeFirstSubStr(winGetOutput, progressBarChar)
	}
	return winGetOutput
}

func removeFirstSubStr(str string, subStr string) string {
	// Above "double-escape" leads to the following string values,
	// that work for regexp.MustCompile(), when looking like that:
	// ^(.*?)\\u0008(.*)$
	// ^(.*?)|(.*)$
	// ^(.*?)/(.*)$
	// ^(.*?)-(.*)$
	// ^(.*?)\\(.*)$
	regexp := regexp.MustCompile("^(.*?)" + subStr + "(.*)$")
	return regexp.ReplaceAllString(str, "${1}$2")
}
