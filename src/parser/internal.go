package parser

import (
	"regexp"
	"strings"
)

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
