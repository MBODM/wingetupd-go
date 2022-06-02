package core

import (
	"errors"

	"github.com/mbodm/wingetupd-go/commands"
)

type PackageInfo struct {
	Package          string
	IsValid          bool
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

func Analyze(packages []string, progress func()) ([]PackageInfo, error) {
	packageInfos := []PackageInfo{}
	for _, pkg := range packages {
		valid, err := commands.Search(pkg)
		if err != nil {
			return packageInfos, errors.New("todo") // todo: chain
		}
		progress()
		listResult, err := commands.List(pkg)
		if err != nil {
			return packageInfos, errors.New("todo") // todo: chain
		}
		progress()
		packageInfo := newPackageInfo(pkg, valid, listResult)
		packageInfos = append(packageInfos, *packageInfo)
	}
	return packageInfos, nil
}

func newPackageInfo(pkg string, valid bool, listResult commands.ListResult) *PackageInfo {
	return &PackageInfo{
		Package:          pkg,
		IsValid:          valid,
		IsInstalled:      listResult.IsInstalled,
		IsUpdatable:      listResult.IsUpdatable,
		InstalledVersion: listResult.InstalledVersion,
		UpdateVersion:    listResult.UpdateVersion,
	}
}
