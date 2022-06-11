package core

import (
	"strings"

	"example.com/mbodm/wingetupd/config"
	"example.com/mbodm/wingetupd/errs"
	"example.com/mbodm/wingetupd/winget"
)

var isInitialized bool

func Init() error {
	if !isInitialized {
		if !winget.Exists() {
			return errs.NewExpectedError("It seems WinGet is not installed on this machine", nil)
		}
		if !config.PackageFileExists() {
			return errs.NewExpectedError("The package-file not exists", nil)
		}
		isInitialized = true
	}
	return nil
}

func AnalyzePackages(packages []string, progress func()) (*PackageData, error) {
	if packages == nil {
		return nil, errs.ArgIsNilError("core.AnalyzePackages")
	}
	if !isInitialized {
		return nil, errs.CreateError("core.AnalyzePackages", "core is not initialized")
	}
	packageInfos := []PackageInfo{}
	for _, pkg := range packages {
		pkg = strings.TrimSpace(pkg)
		if pkg != "" {
			searchResult, err := search(pkg)
			if err != nil {
				return &PackageData{}, errs.WrapError("core.AnalyzePackages", err)
			}
			if progress != nil {
				progress()
			}
			listResult, err := list(pkg)
			if err != nil {
				return &PackageData{}, errs.WrapError("core.AnalyzePackages", err)
			}
			if progress != nil {
				progress()
			}
			packageInfo := PackageInfo{
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
	return createPackageData(packageInfos), nil
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
