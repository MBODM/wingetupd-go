package core

import (
	"example.com/mbodm/wingetupd/parse"
	"example.com/mbodm/wingetupd/winget"
)

func newSearchResult(pkg string, winGetResult *winget.WinGetResult, valid bool) *SearchResult {
	b := basics{pkg, winGetResult.ProcessCall, winGetResult.ConsoleOutput, winGetResult.ExitCode}
	return &SearchResult{b, valid}
}

func newListResult(pkg string, winGetResult *winget.WinGetResult, installed bool, parseResult *parse.ParseResult) *ListResult {
	b := basics{pkg, winGetResult.ProcessCall, winGetResult.ConsoleOutput, winGetResult.ExitCode}
	lr := &ListResult{
		basics:      b,
		IsInstalled: installed,
	}
	// ListResult zero values of these fields
	// are false and "" string, which is fine.
	if parseResult != nil {
		lr.IsUpdatable = parseResult.HasUpdate
		lr.InstalledVersion = parseResult.OldVersion
		lr.UpdateVersion = parseResult.NewVersion
	}
	return lr
}

func newUpgradeResult(pkg string, winGetResult *winget.WinGetResult, updated bool) *UpgradeResult {
	b := basics{pkg, winGetResult.ProcessCall, winGetResult.ConsoleOutput, winGetResult.ExitCode}
	return &UpgradeResult{b, updated}
}
