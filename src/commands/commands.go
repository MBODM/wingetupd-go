package commands

import (
	"errors"
	"strings"

	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func Search(pkg string) (bool, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return false, errors.New("empty string argument")
	}
	result, err := winget.Run("search --exact --id " + pkg)
	if err != nil {
		return false, errors.New("todo") // todo: chain
	}
	return result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg), nil
}

type ListResult struct {
	Package          string
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
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
	installed := result.ExitCode == 0 && strings.Contains(result.ConsoleOutput, pkg)
	if installed {
		parseResult, err := parser.ParseListOutput(result.ConsoleOutput)
		if err != nil {
			return ListResult{}, errors.New("todo") // todo: chain
		}
		return *newListResult(pkg, installed, parseResult), nil
	}
	return *newListResult(pkg, false, parser.ParseResult{}), nil
}

func Upgrade(pkg string) (bool, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		return false, errors.New("empty string argument")
	}
	result, err := winget.Run("upgrade --exact --id " + pkg)
	if err != nil {
		return false, errors.New("todo") // todo: chain
	}
	return result.ExitCode == 0, nil
}

func newListResult(pkg string, installed bool, parseResult parser.ParseResult) *ListResult {
	return &ListResult{
		Package:          pkg,
		IsInstalled:      installed,
		IsUpdatable:      parseResult.HasUpdate,
		InstalledVersion: parseResult.OldVersion,
		UpdateVersion:    parseResult.NewVersion,
	}
}
