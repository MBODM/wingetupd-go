package core

type PackageInfo struct {
	Package          string
	IsValid          bool
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

type PackageData struct {
	ValidPackages         []string
	InvalidPackages       []string
	InstalledPackages     []string
	NonInstalledPackages  []string
	UpdatablePackages     []string
	UpdatablePackageInfos []PackageInfo
	PackageInfos          []PackageInfo
}
