package commands

import (
	"strings"
)

func Search(pkg string, runner WinGetRunner) (*SearchResult, error) {
	caller := "Search"
	pkg, err := guard(caller, pkg, runner)
	if err != nil {
		return nil, err
	}
	result, err := runner(createCommand("search", pkg))
	if err != nil {
		return nil, runnerError(caller, err)
	}
	valid := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	return newSearchResult(pkg, result, valid), nil
}

func List(pkg string, runner WinGetRunner, parser WinGetListParser) (*ListResult, error) {
	caller := "List"
	pkg, err := guard(caller, pkg, runner)
	if err != nil {
		return nil, err
	}
	if parser == nil {
		return nil, argIsNilError(caller, "parser")
	}
	result, err := runner(createCommand("list", pkg))
	if err != nil {
		return nil, runnerError(caller, err)
	}
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parserResult, err := parser(result.ConsoleOutput)
		if err != nil {
			return nil, parserError(caller, err)
		}
		return newListResult(pkg, result, installed, parserResult), nil
	}
	return newListResult(pkg, result, false, nil), nil
}

func Upgrade(pkg string, runner WinGetRunner) (*UpgradeResult, error) {
	caller := "Upgrade"
	pkg, err := guard(caller, pkg, runner)
	if err != nil {
		return nil, err
	}
	result, err := runner(createCommand("upgrade", pkg))
	if err != nil {
		return nil, runnerError(caller, err)
	}
	updated := result.ExitCode == 0
	return newUpgradeResult(pkg, result, updated), nil
}
