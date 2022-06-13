package commands

import (
	"strings"
)

func Search(pkg string, runner WinGetRunner) (*SearchResult, error) {
	caller := "Search"
	if runner == nil {
		return nil, createArgIsNilError(caller, "runner")
	}
	result := runner(createCommand("search", pkg))
	if result == nil {
		return nil, createRunnerError(caller)
	}
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return newSearchResult(pkg, result, valid), nil
}

func List(pkg string, runner WinGetRunner, parser WinGetListParser) (*ListResult, error) {
	caller := "List"
	if runner == nil {
		return nil, createArgIsNilError(caller, "runner")
	}
	if parser == nil {
		return nil, createArgIsNilError(caller, "parser")
	}
	result := runner(createCommand("list", pkg))
	if result == nil {
		return nil, createRunnerError(caller)
	}
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parserResult := parser(result.ConsoleOutput)
		if parserResult == nil {
			return nil, createParserError(caller)
		}
		return newListResult(pkg, result, installed, parserResult), nil
	}
	return newListResult(pkg, result, false, nil), nil
}

func Upgrade(pkg string, runner WinGetRunner) (*UpgradeResult, error) {
	caller := "Upgrade"
	if runner == nil {
		return nil, createArgIsNilError(caller, "runner")
	}
	result := runner(createCommand("upgrade", pkg))
	if result == nil {
		return nil, createRunnerError(caller)
	}
	updated := result.ExitCode == 0
	return newUpgradeResult(pkg, result, updated), nil
}
