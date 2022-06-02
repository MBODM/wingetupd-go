package main

import (
	"fmt"
	"os"

	"github.com/mbodm/wingetupd-go/args"
	"github.com/mbodm/wingetupd-go/collections"
	"github.com/mbodm/wingetupd-go/config"
	"github.com/mbodm/wingetupd-go/console"
	"github.com/mbodm/wingetupd-go/core"
)

const AppName = "wingetupd"
const AppVersion = "2.0.1"
const AppAuthor = "MBODM"
const AppDate = "2022-06-02"

func main() {
	fmt.Println()
	title := fmt.Sprintf("%s %s (by %s %s)", AppName, AppVersion, AppAuthor, AppDate)
	fmt.Println(title)
	fmt.Println()
	if !args.Validate() {
		console.ShowUsage(AppName, false)
		os.Exit(1)
	}
	if args.HelpExists() {
		console.ShowUsage(AppName, true)
		os.Exit(0)
	}
	err := core.Init()
	if err != nil {
		// todo
		os.Exit(1)
	}
	packages, err := config.ReadPackageFile()
	if err != nil {
		// todo
		os.Exit(1)
	}
	console.ShowPackageFileEntries(packages)
	fmt.Println()
	fmt.Print("Processing ...")
	packageInfos, err := core.Analyze(packages, func() { fmt.Print(".") })
	if err != nil {
		// todo
		os.Exit(1)
	}
	fmt.Println(" finished.")
	fmt.Println()
	evalResult := collections.Eval(packageInfos)
	if evalResult.HasInvalidPackages() {
		console.ShowInvalidPackagesError(evalResult.InvalidPackages)
		os.Exit(1)
	}
	if evalResult.HasNonInstalledPackages() {
		console.ShowNonInstalledPackagesError(evalResult.NonInstalledPackages)
		os.Exit(1)
	}
	console.ShowSummary(*evalResult)
	fmt.Println()
}
