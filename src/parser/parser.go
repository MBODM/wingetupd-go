package parser

import (
	"errors"
	"strings"
)

func ParseListOutput(listOutput string) (ParseResult, error) {
	listOutput = strings.TrimSpace(listOutput)
	if listOutput == "" {
		return ParseResult{}, errors.New("WinGet list output is empty")
	}
	versions := getVersions(listOutput)
	if len(versions) < 1 {
		return ParseResult{}, errors.New("WinGet list output not contains any version numbers")
	}
	if len(versions) > 2 {
		return ParseResult{}, errors.New("WinGet list output contains more than 2 version numbers")
	}
	oldVersion, newVersion := getVersionStrings(versions)
	hasUpdate := hasUpdate(listOutput, newVersion)
	return ParseResult{oldVersion, newVersion, hasUpdate}, nil
}
