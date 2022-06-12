package models

// In most languages this errors, if predicate is missing. So a zero result seems ok.
func filter(packageInfos []PackageInfo, predicate func(PackageInfo) bool) []PackageInfo {
	result := []PackageInfo{}
	if predicate == nil {
		return result
	}
	for _, packageInfo := range packageInfos {
		if predicate(packageInfo) {
			result = append(result, packageInfo)
		}
	}
	return result
}

// In most languages this is called "map", but "map" is a language keyword in golang.
func transform(packageInfos []PackageInfo, predicate func(PackageInfo) bool) []string {
	result := []string{}
	if predicate == nil {
		return result
	}
	for _, packageInfo := range packageInfos {
		if predicate(packageInfo) {
			result = append(result, packageInfo.Package)
		}
	}
	return result
}

// Golang REALLY! is missing some functional and generic stuff.
// It´s rather hard, when you have to write such loops in 2022.
// In other languages we use "any()" or "filter()" since years.
// And that Go community is in some weird way even proud of it.
// But hey, Rust´s ownership/borrowchecker is happytime too. :)
