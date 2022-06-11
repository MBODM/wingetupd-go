package console

import (
	"fmt"

	"example.com/mbodm/wingetupd/core"
)

func entryOrEntries[T any](slice []T) string {
	return singularOrPlural(slice, "entry", "entries")
}

func packageOrPackages[T any](slice []T) string {
	return singularOrPlural(slice, "package", "packages")
}

func singularOrPlural[T any](slice []T, singular string, plural string) string {
	len := len(slice)
	if len == 1 {
		return singular
	}
	return plural
}

func listPackages(packages []string) {
	for _, pkg := range packages {
		fmt.Printf("  - %s", pkg)
		fmt.Println()
	}
}

func listUpdateablePackages(packageInfos []core.PackageInfo) {
	for _, pi := range packageInfos {
		fmt.Printf("  - %s %s => %s", pi.Package, pi.InstalledVersion, pi.UpdateVersion)
		fmt.Println()
	}
}
