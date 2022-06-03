package console

import (
	"fmt"
	"strings"

	"github.com/mbodm/wingetupd-go/collections"
)

func ShowUsage(appName string, hideError bool) {
	if !hideError {
		fmt.Println("Error: Unknown parameter(s).")
		fmt.Println()
	}
	fmt.Printf("Usage: %s.exe [--no-log] [--no-confirm]", strings.ToLower(appName))
	fmt.Println()
	fmt.Println()
	fmt.Println("  --no-log      Don´t create log file (useful when running from a folder without write permissions)")
	fmt.Println("  --no-confirm  Don´t ask for update confirmation (useful for script integration)")
	fmt.Println()
	fmt.Println("For more information have a look at the GitHub page (https://github.com/MBODM/wingetupd")
}

func ShowPackageFileEntries(entries []string) {
	fmt.Printf("Found package-file, containing %d %s.", len(entries), entryOrEntries(entries))
	fmt.Println()
}

func ShowInvalidPackagesError(invalidPackages []string) {
	fmt.Println("Error: The package-file contains invalid entries.")
	fmt.Println()
	fmt.Println("The following package-file entries are not valid WinGet package id´s:")
	listPackages(invalidPackages)
	fmt.Println()
	fmt.Println("You can use 'winget search' to list all valid package id´s.")
	fmt.Println()
	fmt.Println("Please verify package-file and try again.")
}

func ShowNonInstalledPackagesError(nonInstalledPackages []string) {
	fmt.Println("Error: The package-file contains non-installed packages.")
	fmt.Println()
	fmt.Println("The following package-file entries are valid WinGet package id´s,")
	fmt.Println("but those packages are not already installed on this machine yet:")
	listPackages(nonInstalledPackages)
	fmt.Println()
	fmt.Println("You can use 'winget list' to show all installed packages and their package id´s.")
	fmt.Println()
	fmt.Println("Please verify package-file and try again.")
}

func ShowSummary(er *collections.EvalResult) {
	fmt.Printf("%d package-file %s processed.", len(er.PackageInfos), entryOrEntries(er.PackageInfos))
	fmt.Println()
	fmt.Printf("%d package-file %s validated.", len(er.ValidPackages), entryOrEntries(er.ValidPackages))
	fmt.Println()
	fmt.Printf("%d %s installed:", len(er.InstalledPackages), packageOrPackages(er.InstalledPackages))
	fmt.Println()
	listPackages(er.InstalledPackages)
	fmt.Printf("%d %s updatable:", len(er.UpdatablePackages), packageOrPackages(er.UpdatablePackages))
	fmt.Print()
	if er.HasUpdatablePackages() {
		fmt.Println(":")
		listUpdateablePackages(er.UpdatablePackages)
	} else {
		fmt.Println(".")
	}
}
