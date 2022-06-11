package prettify

func PrettifyWinGetOutput(winGetOutput string) string {
	winGetOutput = removeProgressBarChars(winGetOutput)
	winGetOutput = removeLeadingReturn(winGetOutput)
	// Todo: Temporary disabled.
	// winGetResult.ConsoleOutput = removeDownloadGraphs(winGetResult.ConsoleOutput)
	return winGetOutput
}
