package parser

import (
	"errors"
	"regexp"
	"strings"
)

type ParseResult struct {
	OldVersion string
	NewVersion string
	HasUpdate  bool
}

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

func getVersions(listOutput string) []string {
	regexp := regexp.MustCompile(`\d+(\.\d+)+`)
	versions := regexp.FindAllString(listOutput, -1)
	return versions
}

func getVersionStrings(versions []string) (oldVersion string, newVersion string) {
	switch len(versions) {
	case 1:
		return versions[0], ""
	case 2:
		return versions[0], versions[1]
	default:
		return "", ""
	}
}

func hasUpdate(listOutput string, newVersion string) bool {
	hasUpdateText := strings.Contains(listOutput, " Verfügbar ") || strings.Contains(listOutput, " Available ")
	hasNewVersion := newVersion != ""
	return hasUpdateText && hasNewVersion
}
