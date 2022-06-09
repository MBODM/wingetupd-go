package parser

import (
	"strings"

	"github.com/mbodm/wingetupd-go/errs"
)

func ParseListOutput(listOutput string) (*ParseResult, error) {
	listOutput = strings.TrimSpace(listOutput)
	if listOutput == "" {
		return &ParseResult{}, errs.NewExpectedError("WinGet list output is empty", nil)
	}
	versions := getVersions(listOutput)
	if len(versions) < 1 {
		return &ParseResult{}, errs.NewExpectedError("WinGet list output not contains any version numbers", nil)
	}
	if len(versions) > 2 {
		return &ParseResult{}, errs.NewExpectedError("WinGet list output contains more than 2 version numbers", nil)
	}
	oldVersion, newVersion := getVersionStrings(versions)
	hasUpdate := hasUpdate(listOutput, newVersion)
	return &ParseResult{oldVersion, newVersion, hasUpdate}, nil
}
