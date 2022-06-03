package commands

import (
	"strings"

	"github.com/mbodm/wingetupd-go/winget"
)

// Not all possible search or list calls produce the same
// output. So therefore some conditionals are in use here.

func prettifyOutput(winGetResult *winget.WinGetResult) {
	if isSuccessfulSearchOrListOutput(winGetResult) {
		winGetResult.ConsoleOutput = removeProgressbarChars(winGetResult.ConsoleOutput)
		winGetResult.ConsoleOutput = removeLeadingReturn(winGetResult.ConsoleOutput)
		winGetResult.ConsoleOutput = removeGraphs(winGetResult.ConsoleOutput)
	}
}

func isSuccessfulSearchOrListOutput(winGetResult *winget.WinGetResult) bool {
	isSearch := strings.Contains(winGetResult.ProcessCall, "search --exact --id")
	isList := strings.Contains(winGetResult.ProcessCall, "list --exact --id")
	return (isSearch || isList) && winGetResult.ExitCode == 0
}

func removeProgressbarChars(output string) string {
	if strings.Contains(output, "\b") {
		output = strings.NewReplacer(
			"\b|", "",
			"\b/", "",
			"\b-", "",
			"\b\\", "",
			"\b", "",
		).Replace(output)
		output = strings.TrimSpace(output)
	}
	return output
}

func removeLeadingReturn(output string) string {
	runes := []rune(output)
	firstChar := string(runes[0:1])
	if firstChar == "\r" {
		output = strings.Replace(output, "\r", "", 1)
		output = strings.TrimSpace(output)
	}
	return output
}

func removeGraphs(output string) string {
	// A successful search or list output contains "Name " as first text,
	// or sometimes after weird download graphs. Solution: Remove graphs.
	// Don´t confuse that download graphs with the \b progress bar chars.
	namePos := strings.Index(output, "Name ")
	if namePos != -1 && namePos != 1 {
		runes := []rune(output)
		tail := runes[namePos:]
		output = string(tail)
	}
	return output
}
