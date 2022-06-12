package parse

import (
	"strings"
)

func ParseListOutput(listOutput string) (*ParseResult, error) {
	listOutput = strings.TrimSpace(listOutput)
	if listOutput == "" {
		return nil, createError("ParseListOutput", "given WinGet list output is empty string")
	}
	versions := getVersions(listOutput)
	if len(versions) < 1 {
		return nil, createError("ParseListOutput", "given WinGet list output not contains any version numbers")
	}
	if len(versions) > 2 {
		return nil, createError("ParseListOutput", "given WinGet list output contains more than 2 version numbers")
	}
	oldVersion, newVersion := getVersionStrings(versions)
	hasUpdate := hasUpdate(listOutput, newVersion)
	return &ParseResult{oldVersion, newVersion, hasUpdate}, nil
}
