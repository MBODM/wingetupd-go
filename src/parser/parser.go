package parser

import (
	"strings"

	"github.com/mbodm/wingetupd-go/app"
)

func ParseListOutput(listOutput string) (ParseResult, error) {
	listOutput = strings.TrimSpace(listOutput)
	if listOutput == "" {
		return ParseResult{}, app.NewAppError("WinGet list output is empty", nil)
	}
	versions := getVersions(listOutput)
	if len(versions) < 1 {
		return ParseResult{}, app.NewAppError("WinGet list output not contains any version numbers", nil)
	}
	if len(versions) > 2 {
		return ParseResult{}, app.NewAppError("WinGet list output contains more than 2 version numbers", nil)
	}
	oldVersion, newVersion := getVersionStrings(versions)
	hasUpdate := hasUpdate(listOutput, newVersion)
	return ParseResult{oldVersion, newVersion, hasUpdate}, nil
}
