package models

type PackageInfo struct {
	Package          string
	IsValid          bool
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

// No need for an interface here imo, since there is
// no benefit in re-implementing data access methods.

type PackageData struct {
	all                  []PackageInfo
	valid                []PackageInfo
	invalid              []PackageInfo
	installed            []PackageInfo
	nonInstalled         []PackageInfo
	updatable            []PackageInfo
	allPackages          []string
	validPackages        []string
	invalidPackages      []string
	installedPackages    []string
	nonInstalledPackages []string
	updatablePackages    []string
	allCount             int
	validCount           int
	invalidCount         int
	installedCount       int
	nonInstalledCount    int
	updatableCount       int
}
