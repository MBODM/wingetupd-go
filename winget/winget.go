package winget

// A few words about exec.Cmd.Output() exit code handling in Go:
// If there was no error at all, this means the exit code was 0.
// If exit code was != 0, there will be an exec.ExitError error.
// If exit code was == 1, the "cmd /C" workaround not found exe.
// If something else went wrong, there will be an os/exec error.
// These facts lead to the somewhat special error handling here.

const WinGetApp = "winget.exe"

func Exists() error {
	processCall := createProcessCall("--version")
	execCommand := createCommand(processCall)
	outputBytes, err := execCommand.Output()
	exitCode, err := handleExecError("Exists", err)
	if err != nil {
		return err
	}
	if exitCode != 0 {
		return createError("Exists", WinGetApp+" exit code was not 0")
	}
	output := string(outputBytes)
	if output == "" {
		return createError("Exists", WinGetApp+" output was empty")
	}
	return nil
}

func Run(params string) (*WinGetResult, error) {
	processCall := createProcessCall(params)
	execCommand := createCommand(processCall)
	outputBytes, err := execCommand.Output()
	exitCode, err := handleExecError("Run", err)
	if err != nil {
		return nil, err
	}
	if exitCode != 0 {
		return &WinGetResult{processCall, "", exitCode}, nil
	}
	consoleOutput := string(outputBytes)
	return &WinGetResult{processCall, consoleOutput, exitCode}, nil
}
