package commands

import (
	"strings"

	"github.com/mbodm/wingetupd-go/errs"
	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func Search(pkg string) (*SearchResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return &SearchResult{}, errs.ArgIsEmptyStringError("commands.Search")
	}
	result, err := winget.Run("search --exact --id " + pkg)
	if err != nil {
		return &SearchResult{}, errs.WrapError("commands.Search", err)
	}
	prettifyOutput(result)
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return newSearchResult(pkg, result, valid), nil
}

func List(pkg string) (*ListResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return &ListResult{}, errs.ArgIsEmptyStringError("commands.List")
	}
	result, err := winget.Run("list --exact --id " + pkg)
	if err != nil {
		return &ListResult{}, errs.WrapError("commands.List", err)
	}
	prettifyOutput(result)
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parseResult, err := parser.ParseListOutput(result.ConsoleOutput)
		if err != nil {
			return &ListResult{}, errs.WrapError("commands.List", err)
		}
		return newListResult(pkg, result, installed, parseResult), nil
	}
	return newListResult(pkg, result, false, nil), nil
}

func Upgrade(pkg string) (*UpgradeResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return &UpgradeResult{}, errs.ArgIsEmptyStringError("commands.Upgrade")
	}
	result, err := winget.Run("upgrade --exact --id " + pkg)
	if err != nil {
		return &UpgradeResult{}, errs.WrapError("commands.Upgrade", err)
	}
	updated := result.ExitCode == 0
	return newUpgradeResult(pkg, result, updated), nil
}
