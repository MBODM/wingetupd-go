package commands

import (
	"errors"
	"strings"

	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func Search(pkg string) (SearchResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return SearchResult{}, errors.New("empty string argument")
	}
	result, err := winget.Run("search --exact --id " + pkg)
	if err != nil {
		return SearchResult{}, errors.New("todo") // todo: chain
	}
	prettifyOutput(&result)
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return *newSearchResult(*newBasics(pkg, result), valid), nil
}

func List(pkg string) (ListResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return ListResult{}, errors.New("empty string argument")
	}
	result, err := winget.Run("list --exact --id " + pkg)
	if err != nil {
		return ListResult{}, errors.New("todo")
	}
	prettifyOutput(&result)
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parseResult, err := parser.ParseListOutput(result.ConsoleOutput)
		if err != nil {
			return ListResult{}, errors.New("todo") // todo: chain
		}
		return *newListResult(*newBasics(pkg, result), installed, parseResult), nil
	}
	return *newListResult(*newBasics(pkg, result), false, parser.ParseResult{}), nil
}

func Upgrade(pkg string) (UpgradeResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return UpgradeResult{}, errors.New("empty string argument")
	}
	result, err := winget.Run("upgrade --exact --id " + pkg)
	if err != nil {
		return UpgradeResult{}, errors.New("todo") // todo: chain
	}
	updated := result.ExitCode == 0
	return *newUpgradeResult(*newBasics(pkg, result), updated), nil
}
