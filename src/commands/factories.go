package commands

import (
	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func newSearchResult(basics Basics, valid bool) *SearchResult {
	return &SearchResult{basics, valid}
}

func newListResult(basics Basics, installed bool, parseResult parser.ParseResult) *ListResult {
	return &ListResult{
		Basics:           basics,
		IsInstalled:      installed,
		IsUpdatable:      parseResult.HasUpdate,
		InstalledVersion: parseResult.OldVersion,
		UpdateVersion:    parseResult.NewVersion,
	}
}

func newUpgradeResult(basics Basics, updated bool) *UpgradeResult {
	return &UpgradeResult{basics, updated}
}

func newBasics(pkg string, winGetResult winget.WinGetResult) *Basics {
	return &Basics{pkg, winGetResult.ProcessCall, winGetResult.ConsoleOutput, winGetResult.ExitCode}
}
