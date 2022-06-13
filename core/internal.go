package core

import (
	"log"
	"strings"

	"example.com/mbodm/wingetupd/commands"
	"example.com/mbodm/wingetupd/config"
	"example.com/mbodm/wingetupd/domain"
	"example.com/mbodm/wingetupd/models"
	"example.com/mbodm/wingetupd/winget"
)

var isInitialized bool

func panicHandler(err error) {
	// Todo
	log.Fatal(err)
}

func initCore() error {
	if !isInitialized {
		if err := winget.Exists(); err != nil {
			return createExpectedError("Init", "It seems WinGet is not installed on this machine")
		}
		if !config.PackageFileExists(panicHandler) {
			return createExpectedError("Init", "The package-file not exists")
		}
		isInitialized = true
	}
	return nil
}

func analyzePackages(packages []string, progress func()) ([]models.PackageInfo, error) {
	const caller = "AnalyzePackages"
	if packages == nil {
		return nil, createArgIsNilError(caller, "packages")
	}
	if !isInitialized {
		return nil, createError(caller, "core is not initialized")
	}
	packageInfos := []models.PackageInfo{}
	for _, pkg := range packages {
		pkg = strings.TrimSpace(pkg)
		if pkg != "" {
			searchResult, err := search(pkg)
			if err != nil {
				return nil, chainError(caller, err)
			}
			if progress != nil {
				progress()
			}
			listResult, err := list(pkg)
			if err != nil {
				return nil, chainError(caller, err)
			}
			if progress != nil {
				progress()
			}
			packageInfo := domain.PackageInfo{
				Package:          pkg,
				IsValid:          searchResult.IsValid,
				IsInstalled:      listResult.IsInstalled,
				IsUpdatable:      listResult.IsUpdatable,
				InstalledVersion: listResult.InstalledVersion,
				UpdateVersion:    listResult.UpdateVersion,
			}
			packageInfos = append(packageInfos, packageInfo)
		}
	}
	return packageInfos, nil
}

func updatePackages(packageData *models.PackageData, progress func()) ([]string, error) {
	caller := "updatePackages"
	if packageData == nil {
		return nil, createArgIsNilError(caller, "packageData")
	}
	if len(packageData.GetAll()) <= 0 {
		return nil, createError(caller, "argument contains empty slice of package infos")
	}
	if !isInitialized {
		return nil, createError(caller, "core is not initialized")
	}
	updatedPackages := []string{}
	for _, packageInfo := range packageData.GetUpdatable() {
		// Todo: Runner
		upgradeResult, err := commands.Upgrade(packageInfo.Package)
		if err != nil {
			return nil, wrapError(caller, "todo", err)
		}
		if upgradeResult.SuccessfullyUpdated {
			updatedPackages = append(updatedPackages, packageInfo.Package)
		}
		if progress != nil {
			progress()
		}
	}
	return updatedPackages, nil
}
