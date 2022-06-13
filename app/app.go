package app

import (
	"fmt"

	"example.com/mbodm/wingetupd/args"
	"example.com/mbodm/wingetupd/config"
	"example.com/mbodm/wingetupd/console"
	"example.com/mbodm/wingetupd/models"
)

const Name = "wingetupd"
const Version = "2.0.1"
const Author = "MBODM"
const Date = "2022-06-09"

func Run() (bool, error) {
	caller := "Run"
	if !args.Validate() {
		console.ShowUsage(Name, false)
		return false, nil
	}
	if args.HelpExists() {
		console.ShowUsage(Name, true)
		return true, nil
	}
	if err := initialize(); err != nil {
		return false, chainError(caller, err)
	}
	panicHandler := func(err error) {
		// Todo
	}
	packages, err := config.ReadPackageFile(panicHandler)
	if err != nil {
		return false, chainError(caller, err)
	}
	console.ShowPackageFileEntries(packages)
	fmt.Println()
	fmt.Print("Processing ...")
	progress := func() {
		fmt.Print(".")
	}
	packageInfos, err := analyzePackages(packages, progress)
	if err != nil {
		return false, chainError(caller, err)
	}
	packageData := models.NewPackageData(packageInfos)
	fmt.Println(" finished.")
	fmt.Println()
	if packageData.ContainsInvalid() {
		console.ShowInvalidPackagesError(packageData.GetInvalidPackages())
		return false, nil
	}
	if packageData.ContainsNonInstalled() {
		console.ShowNonInstalledPackagesError(packageData.GetNonInstalledPackages())
		return false, nil
	}
	updates := []console.Update{}
	for _, packageInfo := range packageData.GetUpdatable() {
		updates = append(updates, console.Update{
			Package:          packageInfo.Package,
			InstalledVersion: packageInfo.InstalledVersion,
			UpdateVersion:    packageInfo.UpdateVersion,
		})
	}
	console.ShowSummary(packageData.GetAllPackages(), packageData.GetValidPackages(), packageData.GetInstalledPackages(), updates)
	fmt.Println()
	if packageData.ContainsUpdatable() {
		shallUpdate := false
		if args.NoConfirmExists() {
			shallUpdate = true
		} else {
			fatalHandler := func(err error) {
				// Todo
			}
			questionResult, err := console.AskUpdateQuestion(packageData.GetUpdatablePackages(), fatalHandler)
			if err != nil {
				return false, chainError(caller, err)
			}
			shallUpdate = questionResult
		}
		if shallUpdate {
			fmt.Println()
			fmt.Print("Updating ...")
			progress = func() {
				fmt.Print("...")
			}
			updatedPackages, err := updatePackages(packageData.GetUpdatable(), progress)
			if err != nil {
				return false, chainError(caller, err)
			}
			fmt.Println(" finished.")
			//updatedPackages := []string{"Mozilla.Firefox", "Microsoft.Edge"}
			console.ShowUpdatedPackages(updatedPackages)
		} else {
			fmt.Println()
			fmt.Println("Canceled, no packages updated.")
		}
		fmt.Println()
	}
	return true, nil
}
