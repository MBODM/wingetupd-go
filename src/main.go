package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/args"
	"github.com/mbodm/wingetupd-go/config"
	"github.com/mbodm/wingetupd-go/console"
	"github.com/mbodm/wingetupd-go/core"
)

func main() {
	fmt.Println()
	title := fmt.Sprintf("%s %s (by %s %s)", app.Name, app.Version, app.Author, app.Date)
	fmt.Println(title)
	fmt.Println()
	result, err := run()
	if err != nil {
		var appError *app.AppError
		if errors.As(err, &appError) {
			if appError.Msg != "STRG+C" {
				fmt.Println("Error: " + appError.Msg + ".")
			}
		} else {
			fmt.Println("Unexpected error(s):", err)
		}
		os.Exit(1)
	}
	if !result {
		os.Exit(1)
	}
	fmt.Println("Have a nice day.")
	os.Exit(0)
}

func run() (bool, error) {
	if !args.Validate() {
		console.ShowUsage(app.Name, false)
		return false, nil
	}
	if args.HelpExists() {
		console.ShowUsage(app.Name, true)
		return false, nil
	}
	err := core.Init()
	if err != nil {
		return false, app.WrapError("main.run", err)
	}
	packages, err := config.ReadPackageFile()
	if err != nil {
		return false, app.WrapError("main.run", err)
	}
	console.ShowPackageFileEntries(packages)
	fmt.Println()
	fmt.Print("Processing ...")
	packageData, err := core.Analyze(packages, func() { fmt.Print(".") })
	if err != nil {
		return false, app.WrapError("main.run", err)
	}
	fmt.Println(" finished.")
	fmt.Println()
	if packageData.HasInvalidPackages() {
		console.ShowInvalidPackagesError(packageData.InvalidPackages)
		os.Exit(1)
	}
	if packageData.HasNonInstalledPackages() {
		console.ShowNonInstalledPackagesError(packageData.NonInstalledPackages)
		os.Exit(1)
	}
	console.ShowSummary(&packageData)
	fmt.Println()
	packageData.UpdatablePackages = []string{"Mozilla.Firefox", "Microsoft.Edge"}
	if packageData.HasUpdatablePackages() {
		shallUpdate := false
		if args.NoConfirmExists() {
			shallUpdate = true
		} else {
			questionResult, err := console.AskUpdateQuestion(packageData.UpdatablePackages)
			if err != nil {
				return false, app.WrapError("main.run", err)
			}
			shallUpdate = questionResult
		}
		if shallUpdate {
			fmt.Println()
			fmt.Print("Updating ......")
			//updatedPackages := core.UpdatePackages(packages, func() { fmt.Print(".") })
			updatedPackages := []string{"hallo", "welt"}
			fmt.Println(" finished.")
			console.ShowUpdatedPackages(updatedPackages)
		} else {
			fmt.Println()
			fmt.Println("Canceled, no packages updated.")
		}
		fmt.Println()
	}
	return true, nil
}
