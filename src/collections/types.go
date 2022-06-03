package collections

import "github.com/mbodm/wingetupd-go/core"

type EvalResult struct {
	ValidPackages        []string
	InvalidPackages      []string
	InstalledPackages    []string
	NonInstalledPackages []string
	UpdatablePackages    []core.PackageInfo
	PackageInfos         []core.PackageInfo
}
