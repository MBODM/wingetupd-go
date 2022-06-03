package commands

import (
	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func newSearchResult(b basics, valid bool) *SearchResult {
	return &SearchResult{b, valid}
}

func newListResult(b basics, installed bool, parseResult parser.ParseResult) *ListResult {
	return &ListResult{
		basics:           b,
		IsInstalled:      installed,
		IsUpdatable:      parseResult.HasUpdate,
		InstalledVersion: parseResult.OldVersion,
		UpdateVersion:    parseResult.NewVersion,
	}
}

func newUpgradeResult(b basics, updated bool) *UpgradeResult {
	return &UpgradeResult{b, updated}
}

func newBasics(pkg string, winGetResult winget.WinGetResult) *basics {
	return &basics{pkg, winGetResult.ProcessCall, winGetResult.ConsoleOutput, winGetResult.ExitCode}
}
