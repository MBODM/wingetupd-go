package console

import "fmt"

func entryOrEntries(count int) string {
	return singularOrPlural(count, "entry", "entries")
}

func packageOrPackages(count int) string {
	return singularOrPlural(count, "package", "packages")
}

func singularOrPlural(count int, singular string, plural string) string {
	if count == 1 {
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
