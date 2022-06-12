package core

import (
	"log"
	"strings"

	"example.com/mbodm/wingetupd/config"
	"example.com/mbodm/wingetupd/errs"
	"example.com/mbodm/wingetupd/shibby"
	"example.com/mbodm/wingetupd/winget"
)

var isInitialized bool

func panicHandler(err error) {
	// Todo
	log.Fatal(err)
}

func Init() error {
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

func AnalyzePackages(packages []string, progress func()) (shibby.IPackageData, error) {
	const caller = "AnalyzePackages"
	if packages == nil {
		return nil, createArgIsNilError(caller, "packages")
	}
	if !isInitialized {
		return nil, createError(caller, "core is not initialized")
	}
	packageInfos := []shibby.PackageInfo{}
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
			packageInfo := shibby.PackageInfo{
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
	return shibby.NewPackageData(packageInfos), nil
}

func UpdatePackagesAsync(packageData *PackageData, progress func()) ([]string, error) {
	if packageData == nil {
		return nil, errs.ArgIsNilError("core.UpdatePackagesAsync")
	}
	if !packageData.HasPackageInfos() {
		return nil, errs.CreateError("core.UpdatePackagesAsync", "argument contains empty slice of package infos")
	}
	if !isInitialized {
		return nil, errs.CreateError("core.UpdatePackagesAsync", "core is not initialized")
	}
	updatedPackages := []string{}
	for _, packageInfo := range packageData.UpdatablePackageInfos {
		upgradeResult, err := upgrade(packageInfo.Package)
		if err != nil {
			return []string{}, errs.WrapError("core.UpdatePackagesAsync", err)
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
