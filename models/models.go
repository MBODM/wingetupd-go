package models

// Since this object is designed to be immutable, there is no benefit in
// running the filters, transforms and counts on every access, as mostly
// done in such cases. Instead running them once and caching the results.

func NewPackageData(packageInfos []PackageInfo) *PackageData {
	result := &PackageData{
		all:                  packageInfos,
		valid:                filter(packageInfos, func(pi PackageInfo) bool { return pi.IsValid }),
		invalid:              filter(packageInfos, func(pi PackageInfo) bool { return !pi.IsValid }),
		installed:            filter(packageInfos, func(pi PackageInfo) bool { return pi.IsInstalled }),
		nonInstalled:         filter(packageInfos, func(pi PackageInfo) bool { return !pi.IsInstalled }),
		updatable:            filter(packageInfos, func(pi PackageInfo) bool { return pi.IsUpdatable }),
		allPackages:          transform(packageInfos, func(pi PackageInfo) bool { return true }),
		validPackages:        transform(packageInfos, func(pi PackageInfo) bool { return pi.IsValid }),
		invalidPackages:      transform(packageInfos, func(pi PackageInfo) bool { return !pi.IsValid }),
		installedPackages:    transform(packageInfos, func(pi PackageInfo) bool { return pi.IsInstalled }),
		nonInstalledPackages: transform(packageInfos, func(pi PackageInfo) bool { return !pi.IsInstalled }),
		updatablePackages:    transform(packageInfos, func(pi PackageInfo) bool { return pi.IsUpdatable }),
	}
	result.allCount = len(result.all)
	result.validCount = len(result.valid)
	result.invalidCount = len(result.invalid)
	result.installedCount = len(result.installed)
	result.nonInstalledCount = len(result.nonInstalled)
	result.updatableCount = len(result.updatable)
	return result
}
