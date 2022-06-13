package commands

import "strings"

func guard(caller string, pkg string, runner WinGetRunner) (string, error) {
	// Don´t forget: Argument names for errors are hardcoded here.
	// So make sure the exported functions use the same arg names.
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return "", argIsEmptyStringError(caller, "pkg")
	}
	if runner == nil {
		return "", argIsNilError(caller, "runner")
	}
	return pkg, nil
}

func createCommand(command, pkg string) string {
	return command + " --exact --id " + pkg
}

func newBasics(pkg string, result *WinGetRunnerResult) *basics {
	return &basics{
		Package:             pkg,
		WinGetProcessCall:   result.ProcessCall,
		WinGetConsoleOutput: result.ConsoleOutput,
		WinGetExitCode:      result.ExitCode,
	}
}

func newSearchResult(pkg string, result *WinGetRunnerResult, valid bool) *SearchResult {
	return &SearchResult{*newBasics(pkg, result), valid} // <-- Deref ptr, to promote.
}

func newListResult(pkg string, result *WinGetRunnerResult, installed bool, parserResult *WinGetListParserResult) *ListResult {
	listResult := &ListResult{
		basics:      *newBasics(pkg, result), // <-- Deref ptr, to promote.
		IsInstalled: installed,
	}
	// ListResult zero values of these fields
	// are false and "" string, which is fine.
	if parserResult != nil {
		listResult.IsUpdatable = hasUpdate(result.ConsoleOutput, parserResult.NewVersion)
		listResult.InstalledVersion = parserResult.OldVersion
		listResult.UpdateVersion = parserResult.NewVersion
	}
	return listResult
}

func newUpgradeResult(pkg string, result *WinGetRunnerResult, updated bool) *UpgradeResult {
	return &UpgradeResult{*newBasics(pkg, result), updated} // <-- Deref ptr, to promote.
}

func hasUpdate(winGetListOutput, newVersion string) bool {
	hasUpdateText := strings.Contains(winGetListOutput, " Verfügbar ") || strings.Contains(winGetListOutput, " Available ")
	hasNewVersion := newVersion != ""
	return hasUpdateText && hasNewVersion
}
