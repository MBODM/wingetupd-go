package shibby

type PackageInfo struct {
	Package          string
	IsValid          bool
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

type IPackageData interface {
	HasValidPackages() bool
	HasInvalidPackages() bool
	HasInstalledPackages() bool
	HasNonInstalledPackages() bool
	HasUpdatablePackages() bool
	HasUpdatablePackageInfos() bool
	HasPackageInfos() bool
	ValidPackages() []string
	InvalidPackages() []string
	InstalledPackages() []string
	NonInstalledPackages() []string
	UpdatablePackages() []string
	UpdatablePackageInfos() []PackageInfo
	PackageInfos() []PackageInfo
}

type PackageData struct {
	packageInfos []PackageInfo
}
