package core

import (
	"fmt"

	"example.com/mbodm/wingetupd/args"
	"example.com/mbodm/wingetupd/config"
	"example.com/mbodm/wingetupd/console"
	"example.com/mbodm/wingetupd/core"
	"example.com/mbodm/wingetupd/errs"
)

const Name = "wingetupd"
const Version = "2.0.1"
const Author = "MBODM"
const Date = "2022-06-09"

func Run() (bool, error) {
	if !args.Validate() {
		console.ShowUsage(Name, false)
		return false, nil
	}
	if args.HelpExists() {
		console.ShowUsage(Name, true)
		return false, nil
	}
	err := core.Init()
	if err != nil {
		return false, errs.WrapError("app.Run", err)
	}
	packages, err := config.ReadPackageFile()
	if err != nil {
		return false, errs.WrapError("app.Run", err)
	}
	console.ShowPackageFileEntries(packages)
	fmt.Println()
	fmt.Print("Processing ...")
	packageData, err := core.AnalyzePackages(packages, func() { fmt.Print(".") })
	if err != nil {
		return false, errs.WrapError("app.Run", err)
	}
	fmt.Println(" finished.")
	fmt.Println()
	if packageData.HasInvalidPackages() {
		console.ShowInvalidPackagesError(packageData.InvalidPackages)
		return false, nil
	}
	if packageData.HasNonInstalledPackages() {
		console.ShowNonInstalledPackagesError(packageData.NonInstalledPackages)
		return false, nil
	}
	console.ShowSummary(packageData)
	fmt.Println()
	packageData.UpdatablePackages = []string{"Mozilla.Firefox", "Microsoft.Edge"}
	if packageData.HasUpdatablePackages() {
		shallUpdate := false
		if args.NoConfirmExists() {
			shallUpdate = true
		} else {
			questionResult, err := console.AskUpdateQuestion(packageData.)
			if err != nil {
				return false, errs.WrapError("app.Run", err)
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
