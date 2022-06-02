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

func (e EvalResult) HasValidPackages() bool {
	return len(e.ValidPackages) > 1
}

func (e EvalResult) HasInvalidPackages() bool {
	return len(e.InvalidPackages) > 1
}

func (e EvalResult) HasInstalledPackages() bool {
	return len(e.InstalledPackages) > 1
}

func (e EvalResult) HasNonInstalledPackages() bool {
	return len(e.NonInstalledPackages) > 1
}

func (e EvalResult) HasUpdatablePackages() bool {
	return len(e.UpdatablePackages) > 1
}

func (e EvalResult) HasPackageInfos() bool {
	return len(e.PackageInfos) > 1
}

func Eval(packageInfos []core.PackageInfo) *EvalResult {
	result := EvalResult{PackageInfos: packageInfos}
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
			result.UpdatablePackages = append(result.UpdatablePackages, pi)
		}
	}
	return &result
}

// Golang REALLY! is missing some functional and generic stuff.
// It´s rather hard, when you have to write such loops in 2022.
// In other languages we use "any()" or "filter()" since years.
// And that Go community is in some weird way even proud of it.
// But hey, Rust´s ownership/borrowchecker is happytime too. :)