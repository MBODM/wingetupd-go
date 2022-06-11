package core

import (
	"strings"

	"example.com/mbodm/wingetupd/errs"
	"example.com/mbodm/wingetupd/parse"
	"example.com/mbodm/wingetupd/prettify"
	"example.com/mbodm/wingetupd/winget"
)

func search(pkg string) (*SearchResult, error) {
	result, err := winget.Run("search --exact --id " + pkg)
	if err != nil {
		return nil, errs.WrapError("core.search", err)
	}
	err = prettify.PrettifyWinGetOutput(result)
	if err != nil {
		return nil, errs.WrapError("core.search", err)
	}
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return newSearchResult(pkg, result, valid), nil
}

func list(pkg string) (*ListResult, error) {
	result, err := winget.Run("list --exact --id " + pkg)
	if err != nil {
		return nil, errs.WrapError("core.list", err)
	}

	err = prettify.PrettifyWinGetOutput(result)
	if err != nil {

	}

	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parseResult, err := parse.ParseListOutput(result.ConsoleOutput)
		if err != nil {
			return nil, errs.WrapError("core.list", err)
		}
		return newListResult(pkg, result, installed, parseResult), nil
	}
	return newListResult(pkg, result, false, nil), nil
}

func upgrade(pkg string) (*UpgradeResult, error) {
	result, err := runWinGetCommand("upgrade --exact --id ", pkg)
	if err != nil {
		return &UpgradeResult{}, errs.WrapError("core.upgrade", err)
	}
	updated := result.ExitCode == 0
	return newUpgradeResult(pkg, result, updated), nil
}

func runWinGetCommand(pkg string, cmd string) (*winget.WinGetResult, error) {
	result, err := winget.Run(cmd + pkg)
	if err != nil {
		return nil, errs.WrapError("core.upgrade", err)
	}
	return result, nil
}
