package core

// Using golang´s "promoted fields" feature here.

type basics struct {
	Package             string
	WinGetProcessCall   string
	WinGetConsoleOutput string
	WinGetExitCode      int
}

type SearchResult struct {
	basics
	IsValid bool
}

type ListResult struct {
	basics
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

type UpgradeResult struct {
	basics
	SuccessfullyUpdated bool
}

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
