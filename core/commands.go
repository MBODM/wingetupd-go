package core

import (
	"strings"

	"example.com/mbodm/wingetupd/parse"
	"example.com/mbodm/wingetupd/prettify"
	"example.com/mbodm/wingetupd/winget"
)

func search(pkg string) (*SearchResult, error) {
	result, err := winget.Run("search --exact --id " + pkg)
	if err != nil {
		return nil, chainError("search", err)
	}
	result.ConsoleOutput = prettify.PrettifyWinGetOutput(result.ConsoleOutput)
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return newSearchResult(pkg, result, valid), nil
}

func list(pkg string) (*ListResult, error) {
	result, err := winget.Run("list --exact --id " + pkg)
	if err != nil {
		return nil, chainError("list", err)
	}
	result.ConsoleOutput = prettify.PrettifyWinGetOutput(result.ConsoleOutput)
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parseResult, err := parse.ParseListOutput(result.ConsoleOutput)
		if err != nil {
			return nil, chainError("list", err)
		}
		return newListResult(pkg, result, installed, parseResult), nil
	}
	return newListResult(pkg, result, false, nil), nil
}

func upgrade(pkg string) (*UpgradeResult, error) {
	result, err := winget.Run("upgrade --exact --id " + pkg)
	if err != nil {
		return nil, chainError("upgrade", err)
	}
	updated := result.ExitCode == 0
	return newUpgradeResult(pkg, result, updated), nil
}

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

// Using golang´s "promoted fields" feature here.

type basics struct {
	Package             string
	WinGetProcessCall   string
	WinGetConsoleOutput string
	WinGetExitCode      int
}

type SearchResult struct {
	basics
	IsValid bool
}

type ListResult struct {
	basics
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

type UpgradeResult struct {
	basics
	SuccessfullyUpdated bool
}
