package app

import (
	"log"
	"strings"

	"example.com/mbodm/wingetupd/commands"
	"example.com/mbodm/wingetupd/config"
	"example.com/mbodm/wingetupd/models"
	"example.com/mbodm/wingetupd/parse"
	"example.com/mbodm/wingetupd/winget"
)

var isInitialized bool

func panicHandler(err error) {
	// Todo
	log.Fatal(err)
}

func initialize() error {
	caller := "initialize"
	if !isInitialized {
		if err := winget.Exists(); err != nil {
			return wrapErrorIntoExpectedError(caller, "It seems WinGet is not installed on this machine", err)
		}
		if !config.PackageFileExists(panicHandler) {
			return createExpectedError(caller, "The package-file not exists")
		}
		isInitialized = true
	}
	return nil
}

func winGetRunner(winGetParams string) (*commands.WinGetRunnerResult, error) {
	result, err := winget.Run(winGetParams)
	if err != nil {
		return nil, chainError("runner", err)
	}
	return &commands.WinGetRunnerResult{
		ProcessCall:   result.ProcessCall,
		ConsoleOutput: result.ConsoleOutput,
		ExitCode:      result.ExitCode,
	}, nil
}

func winGetListParser(winGetListOutput string) (*commands.WinGetListParserResult, error) {
	result, err := parse.ParseVersions(winGetListOutput)
	if err != nil {
		return nil, chainError("parser", err)
	}
	return &commands.WinGetListParserResult{
		OldVersion: result.OldVersion,
		NewVersion: result.NewVersion,
	}, nil
}

func analyzePackages(packages []string, progress func()) ([]models.PackageInfo, error) {
	const caller = "analyzePackages"
	if packages == nil {
		return nil, argIsNilError(caller, "packages")
	}
	if !isInitialized {
		return nil, notInitializedError(caller)
	}
	packageInfos := []models.PackageInfo{}
	for _, pkg := range packages {
		pkg = strings.TrimSpace(pkg)
		if pkg != "" {
			searchResult, err := commands.Search(pkg, winGetRunner)
			if err != nil {
				return nil, chainError(caller, err)
			}
			if progress != nil {
				progress()
			}
			listResult, err := commands.List(pkg, winGetRunner, winGetListParser)
			if err != nil {
				return nil, chainError(caller, err)
			}
			if progress != nil {
				progress()
			}
			packageInfo := models.PackageInfo{
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

func updatePackages(packageInfos []models.PackageInfo, progress func()) ([]string, error) {
	caller := "updatePackages"
	if packageInfos == nil {
		return nil, argIsNilError(caller, "packageInfos")
	}
	if !isInitialized {
		return nil, notInitializedError(caller)
	}
	updatedPackages := []string{}
	for _, packageInfo := range packageInfos {
		if packageInfo.IsUpdatable {
			upgradeResult, err := commands.Upgrade(packageInfo.Package, winGetRunner)
			if err != nil {
				return nil, chainError(caller, err)
			}
			if upgradeResult.SuccessfullyUpdated {
				updatedPackages = append(updatedPackages, packageInfo.Package)
			}
			if progress != nil {
				progress()
			}
		}
	}
	return updatedPackages, nil
}
