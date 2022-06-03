package core

import (
	"errors"

	"github.com/mbodm/wingetupd-go/commands"
	"github.com/mbodm/wingetupd-go/config"
	"github.com/mbodm/wingetupd-go/winget"
)

var isInitialized bool

func Init() error {
	if !isInitialized {
		if !winget.IsInstalled() {
			return errors.New("it seems WinGet is not installed on this machine")
		}
		exists, err := config.PackageFileExists()
		if err != nil {
			return errors.New("todo") // todo: chain
		}
		if !exists {
			return errors.New("the package-file not exists")
		}
		isInitialized = true
	}
	return nil
}

func Analyze(packages []string, progress func()) ([]PackageInfo, error) {
	packageInfos := []PackageInfo{}
	for _, pkg := range packages {
		searchResult, err := commands.Search(pkg)
		if err != nil {
			return packageInfos, errors.New("todo") // todo: chain
		}
		progress()
		listResult, err := commands.List(pkg)
		if err != nil {
			return packageInfos, errors.New("todo") // todo: chain
		}
		progress()
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
	return packageInfos, nil
}
