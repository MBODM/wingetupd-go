package core

func createPackageData(packageInfos []PackageInfo) *PackageData {
	result := &PackageData{PackageInfos: packageInfos}
	for _, pi := range packageInfos {
		if pi.IsValid {
			result.ValidPackages = append(result.ValidPackages, pi.Package)
		} else {
			result.InvalidPackages = append(result.InvalidPackages, pi.Package)
		}
		if pi.IsInstalled {
			result.InstalledPackages = append(result.InstalledPackages, pi.Package)
		} else {
			result.NonInstalledPackages = append(result.NonInstalledPackages, pi.Package)
		}
		if pi.IsUpdatable {
			result.UpdatablePackages = append(result.UpdatablePackages, pi.Package)
			result.UpdatablePackageInfos = append(result.UpdatablePackageInfos, pi)
		}
	}
	return result
}

// Golang REALLY! is missing some functional and generic stuff.
// It´s rather hard, when you have to write such loops in 2022.
// In other languages we use "any()" or "filter()" since years.
// And that Go community is in some weird way even proud of it.
// But hey, Rust´s ownership/borrowchecker is happytime too. :)
