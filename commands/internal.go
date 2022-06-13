package commands

import "strings"

func createCommand(command, pkg string) string {
	return command + " --exact --id " + pkg
}

func newSearchResult(pkg string, result *WinGetRunnerResult, valid bool) *SearchResult {
	b := basics{pkg, result.ProcessCall, result.ConsoleOutput, result.ExitCode}
	return &SearchResult{b, valid}
}

func newListResult(pkg string, result *WinGetRunnerResult, installed bool, parserResult *WinGetListParserResult) *ListResult {
	b := basics{pkg, result.ProcessCall, result.ConsoleOutput, result.ExitCode}
	lr := &ListResult{
		basics:      b,
		IsInstalled: installed,
	}
	// ListResult zero values of these fields
	// are false and "" string, which is fine.
	if parserResult != nil {
		lr.IsUpdatable = hasUpdate(result.ConsoleOutput, parserResult.NewVersion)
		lr.InstalledVersion = parserResult.OldVersion
		lr.UpdateVersion = parserResult.NewVersion
	}
	return lr
}

func newUpgradeResult(pkg string, result *WinGetRunnerResult, updated bool) *UpgradeResult {
	b := basics{pkg, result.ProcessCall, result.ConsoleOutput, result.ExitCode}
	return &UpgradeResult{b, updated}
}

func hasUpdate(winGetListOutput, newVersion string) bool {
	hasUpdateText := strings.Contains(winGetListOutput, " Verfügbar ") || strings.Contains(winGetListOutput, " Available ")
	hasNewVersion := newVersion != ""
	return hasUpdateText && hasNewVersion
}
