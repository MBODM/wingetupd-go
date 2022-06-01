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
	return WinGetResult{
			ProcessCall:   getProcessCall(params),
			ConsoleOutput: string(outputBytes),
			ExitCode:      0},
		nil
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
	fmt.Println(fmt.Sprintf("%U", 8))
	winGetOutput = strings.TrimSpace(winGetOutput)
	winGetOutput = strings.ReplaceAll(winGetOutput, "", fmt.Sprintf("%U", 8))
	winGetOutput = strings.ReplaceAll(winGetOutput, "|", "")
	winGetOutput = strings.ReplaceAll(winGetOutput, "/", "")
	winGetOutput = strings.ReplaceAll(winGetOutput, "-", "")
	winGetOutput = strings.ReplaceAll(winGetOutput, `\`, "")

	winGetOutput = strings.ReplaceAll(winGetOutput, "", "")

	strings.TrimSpace(winGetOutput)
	return winGetOutput
}

// fn remove_progressbar_chars(output_string: &str) -> String {
//     output_string
//         .trim()
//         .replacen(&['\u{0008}', '|', '/', '-', '\\'][..], "", 1)
//         .trim()
//         .to_string()
// }
