package parse

import (
	"strings"
)

func ParseVersions(winGetListOutput string) (*ParseResult, error) {
	caller := "ParseVersions"
	winGetListOutput = strings.TrimSpace(winGetListOutput)
	if winGetListOutput == "" {
		return nil, createError(caller, "given WinGet list output is empty string")
	}
	versions := getVersions(winGetListOutput)
	if len(versions) < 1 {
		return nil, createError(caller, "given WinGet list output not contains any version numbers")
	}
	if len(versions) > 2 {
		return nil, createError(caller, "given WinGet list output contains more than 2 version numbers")
	}
	oldVersion, newVersion := getVersionStrings(versions)
	return &ParseResult{oldVersion, newVersion}, nil
}
