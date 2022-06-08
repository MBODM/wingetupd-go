package core

import (
	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/commands"
	"github.com/mbodm/wingetupd-go/config"
	"github.com/mbodm/wingetupd-go/winget"
)

var isInitialized bool

func Init() error {
	if !isInitialized {
		if !winget.IsInstalled() {
			return app.NewAppError("It seems WinGet is not installed on this machine", nil)
		}
		exists, err := config.PackageFileExists()
		if err != nil {
			return app.WrapError("core.Init", err)
		}
		if !exists {
			return app.NewAppError("The package-file not exists", nil)
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
			return packageInfos, app.WrapError("core.Analyze", err)
		}
		progress()
		listResult, err := commands.List(pkg)
		if err != nil {
			return packageInfos, app.WrapError("core.Analyze", err)
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
