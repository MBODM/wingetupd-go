package shibby

func GetValidPackages(packageInfos []PackageInfo) []string {
	return getFilteredPackages(packageInfos, func(pi PackageInfo) bool { return pi.IsValid })
}

func GetInvalidPackages(packageInfos []PackageInfo) []string {
	return getFilteredPackages(packageInfos, func(pi PackageInfo) bool { return !pi.IsValid })
}

func GetInstalledPackages(packageInfos []PackageInfo) []string {
	return getFilteredPackages(packageInfos, func(pi PackageInfo) bool { return pi.IsInstalled })
}

func GetNonInstalledPackages(packageInfos []PackageInfo) []string {
	return getFilteredPackages(packageInfos, func(pi PackageInfo) bool { return !pi.IsInstalled })
}

func GetUpdatablePackages(packageInfos []PackageInfo) []string {
	return getFilteredPackages(packageInfos, func(pi PackageInfo) bool { return pi.IsUpdatable })
}

func ContainsValidPackages(packageInfos []PackageInfo) bool {
	return packageInfosContains(packageInfos, func(pi PackageInfo) bool { return pi.IsValid })
}

func ContainsInvalidPackages(packageInfos []PackageInfo) bool {
	return packageInfosContains(packageInfos, func(pi PackageInfo) bool { return !pi.IsValid })
}

func ContainsInstalledPackages(packageInfos []PackageInfo) bool {
	return packageInfosContains(packageInfos, func(pi PackageInfo) bool { return pi.IsInstalled })
}

func ContainsNonInstalledPackages(packageInfos []PackageInfo) bool {
	return packageInfosContains(packageInfos, func(pi PackageInfo) bool { return !pi.IsInstalled })
}

func ContainsUpdatablePackages(packageInfos []PackageInfo) bool {
	return packageInfosContains(packageInfos, func(pi PackageInfo) bool { return pi.IsUpdatable })
}

func FilterUpdatables(packageInfos []PackageInfo) []PackageInfo {
	return getFilteredPackageInfos(packageInfos, func(pi PackageInfo) bool { return pi.IsValid })
}
