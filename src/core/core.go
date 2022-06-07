package core

import (
	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/commands"
	"github.com/mbodm/wingetupd-go/winget"
)

var isInitialized bool

func Init() error {
	if !isInitialized {
		_, e := winget.Run("--fuzz")
		if e != nil {
			return app.WrapError("", e)
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
