package console

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
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
	fmt.Printf("Found package-file, containing %d %s.", len(fileEntries), entryOrEntries(len(fileEntries)))
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

func ShowSummary(fileEntries, validPackages, installedPackages []string, updates []Update) {
	fmt.Printf("%d package-file %s processed.", len(fileEntries), entryOrEntries(len(fileEntries)))
	fmt.Println()
	fmt.Printf("%d package-file %s validated.", len(validPackages), entryOrEntries(len(validPackages)))
	fmt.Println()
	fmt.Printf("%d %s installed:", len(installedPackages), packageOrPackages(len(installedPackages)))
	fmt.Println()
	listPackages(installedPackages)
	fmt.Printf("%d %s updatable", len(updates), packageOrPackages(len(updates)))
	fmt.Print()
	if len(updates) > 0 {
		fmt.Println(":")
		for _, update := range updates {
			fmt.Printf("  - %s %s => %s", update.Package, update.InstalledVersion, update.UpdateVersion)
			fmt.Println()
		}
	} else {
		fmt.Println(".")
	}
}

func AskUpdateQuestion(updatablePackages []string, fatalHandler func(error)) (bool, error) {
	fmt.Printf("Update %d %s ? [y/N]: ", len(updatablePackages), packageOrPackages(len(updatablePackages)))
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
	fmt.Printf("%d %s updated:", len(updatedPackages), packageOrPackages(len(updatedPackages)))
	fmt.Println()
	listPackages(updatedPackages)
}
