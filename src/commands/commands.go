package commands

import (
	"strings"

	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func Search(pkg string) (SearchResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return SearchResult{}, app.EmptyStrArgError("commands.Search")
	}
	result, err := winget.Run("search --exact --id " + pkg)
	if err != nil {
		return SearchResult{}, app.WrapError("commands.Search", err)
	}
	prettifyOutput(&result)
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return *newSearchResult(*newBasics(pkg, result), valid), nil
}

func List(pkg string) (ListResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return ListResult{}, app.EmptyStrArgError("commands.List")
	}
	result, err := winget.Run("list --exact --id " + pkg)
	if err != nil {
		return ListResult{}, app.WrapError("commands.List", err)
	}
	prettifyOutput(&result)
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parseResult, err := parser.ParseListOutput(result.ConsoleOutput)
		if err != nil {
			return ListResult{}, app.WrapError("commands.List", err)
		}
		return *newListResult(*newBasics(pkg, result), installed, parseResult), nil
	}
	return *newListResult(*newBasics(pkg, result), false, parser.ParseResult{}), nil
}

func Upgrade(pkg string) (UpgradeResult, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return UpgradeResult{}, app.EmptyStrArgError("commands.Upgrade")
	}
	result, err := winget.Run("upgrade --exact --id " + pkg)
	if err != nil {
		return UpgradeResult{}, app.WrapError("commands.Upgrade", err)
	}
	updated := result.ExitCode == 0
	return *newUpgradeResult(*newBasics(pkg, result), updated), nil
}
