package parse

import (
	"regexp"
)

func getVersions(winGetListOutput string) []string {
	regexp := regexp.MustCompile(`\d+(\.\d+)+`)
	versions := regexp.FindAllString(winGetListOutput, -1)
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
