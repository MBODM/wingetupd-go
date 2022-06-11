package console

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"example.com/mbodm/wingetupd/core"
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
	fmt.Println("For more information have a look at the GitHub page (https://github.com/mbodm/wingetupd)")
}

func ShowPackageFileEntries(fileEntries []string) {
	fmt.Printf("Found package-file, containing %d %s.", len(fileEntries), entryOrEntries(fileEntries))
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

func ShowSummary(pd *core.PackageData) {
	fmt.Printf("%d package-file %s processed.", len(pd.PackageInfos), entryOrEntries(pd.PackageInfos))
	fmt.Println()
	fmt.Printf("%d package-file %s validated.", len(pd.ValidPackages), entryOrEntries(pd.ValidPackages))
	fmt.Println()
	fmt.Printf("%d %s installed:", len(pd.InstalledPackages), packageOrPackages(pd.InstalledPackages))
	fmt.Println()
	listPackages(pd.InstalledPackages)
	fmt.Printf("%d %s updatable", len(pd.UpdatablePackages), packageOrPackages(pd.UpdatablePackages))
	fmt.Print()
	if pd.HasUpdatablePackages() {
		fmt.Println(":")
		listUpdateablePackages(pd.UpdatablePackageInfos)
	} else {
		fmt.Println(".")
	}
}

func AskUpdateQuestion(updatablePackages []string, fatalHandler func(error)) (bool, error) {
	fmt.Printf("Update %d %s ? [y/N]: ", len(updatablePackages), packageOrPackages(updatablePackages))
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fatalHandler(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	bytes := make([]byte, 1)
	var b byte
	for {
		count, err := os.Stdin.Read(bytes)
		if err != nil {
			fatalHandler(err)
		}
		if count != 1 {
			fatalHandler(fmt.Errorf("[console.AskUpdateQuestion] invalid os.Stdin.Read() result"))
		}
		b = bytes[0]
		// On Windows value 13 means ENTER key and value 3 means STRG+C was pressed.
		if b == 'y' || b == 'Y' || b == 'n' || b == 'N' || b == 13 || b == 3 {
			break
		}
	}
	if b == 3 {
		return false, fmt.Errorf("STRG+C") // Todo
	}
	if b == 13 {
		fmt.Println("N")
	} else {
		fmt.Println(string(b))
	}
	return b == 'y' || b == 'Y', nil
}

func ShowUpdatedPackages(updatedPackages []string) {
	fmt.Println()
	fmt.Printf("%d %s updated:", len(updatedPackages), packageOrPackages(updatedPackages))
	fmt.Println()
	listPackages(updatedPackages)
}
