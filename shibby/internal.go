package shibby

func getFilteredPackages(packageInfos []PackageInfo, predicate func(PackageInfo) bool) []string {
	result := []string{}
	for _, pi := range packageInfos {
		if predicate != nil {
			if predicate(pi) {
				result = append(result, pi.Package)
			}
		}
	}
	return result
}

func getFilteredPackageInfos(packageInfos []PackageInfo, predicate func(PackageInfo) bool) []PackageInfo {
	result := []PackageInfo{}
	for _, pi := range packageInfos {
		if predicate != nil {
			if predicate(pi) {
				result = append(result, pi)
			}
		}
	}
	return result
}

func packageInfosContains(packageInfos []PackageInfo, predicate func(PackageInfo) bool) bool {
	for _, pi := range packageInfos {
		if predicate != nil {
			if predicate(pi) {
				return true
			}
		}
	}
	return false
}

// Golang REALLY! is missing some functional and generic stuff.
// It´s rather hard, when you have to write such loops in 2022.
// In other languages we use "any()" or "filter()" since years.
// And that Go community is in some weird way even proud of it.
// But hey, Rust´s ownership/borrowchecker is happytime too. :)
