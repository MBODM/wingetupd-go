package core

import (
	"strings"

	"github.com/mbodm/wingetupd-go/commands"
	"github.com/mbodm/wingetupd-go/config"
	"github.com/mbodm/wingetupd-go/errs"
	"github.com/mbodm/wingetupd-go/winget"
)

var isInitialized bool

func Init() error {
	if !isInitialized {
		if !winget.IsInstalled() {
			return errs.NewExpectedError("It seems WinGet is not installed on this machine", nil)
		}
		exists, err := config.PackageFileExists()
		if err != nil {
			return errs.WrapError("core.Init", err)
		}
		if !exists {
			return errs.NewExpectedError("The package-file not exists", nil)
		}
		isInitialized = true
	}
	return nil
}

func AnalyzePackages(packages []string, progress func()) (*PackageData, error) {
	if packages == nil {
		return nil, errs.ArgIsNilError("core.AnalyzePackages")
	}
	packageInfos := []PackageInfo{}
	for _, pkg := range packages {
		pkg = strings.TrimSpace(pkg)
		if pkg != "" {
			searchResult, err := commands.Search(pkg)
			if err != nil {
				return &PackageData{}, errs.WrapError("core.AnalyzePackages", err)
			}
			if progress != nil {
				progress()
			}
			listResult, err := commands.List(pkg)
			if err != nil {
				return &PackageData{}, errs.WrapError("core.AnalyzePackages", err)
			}
			if progress != nil {
				progress()
			}
			packageInfo := PackageInfo{
				Package:          pkg,
				IsValid:          searchResult.IsValid,
				IsInstalled:      listResult.IsInstalled,
				IsUpdatable:      listResult.IsUpdatable,
				InstalledVersion: listResult.InstalledVersion,
				UpdateVersion:    listResult.UpdateVersion,
			}
			packageInfos = append(packageInfos, packageInfo)
		}
	}
	return createPackageData(packageInfos), nil
}

// func UpdatePackagesAsync(packageData PackageData, progress func()) ([]string, error) {

// }

// public async Task<IEnumerable<string>> UpdatePackagesAsync(
// 	IEnumerable<PackageInfo> packageInfos,
// 	IProgress<PackageProgressData>? progress = default,
// 	CancellationToken cancellationToken = default)
// {
// 	if (packageInfos is null)
// 	{
// 		throw new ArgumentNullException(nameof(packageInfos));
// 	}

// 	if (!packageInfos.Any())
// 	{
// 		throw new ArgumentException("Given list of package infos is empty.", nameof(packageInfos));
// 	}

// 	if (!isInitialized)
// 	{
// 		throw new BusinessLogicException($"{nameof(BusinessLogic)} not initialized.");
// 	}

// 	// We can not use a concurrent logic here, by using some typical Task.WhenAll() approach. Because
// 	// WinGet fails with "Failed in attempting to update the source" errors, when running in parallel.
// 	// Therefore we sadly have to stick here with the non-concurrent, sequential, way slower approach.
// 	// Nonetheless, all parts and modules of this App are designed with a concurrent approach in mind.
// 	// So, if WinGet may change it´s behaviour in future, we are ready to use the concurrent approach.

// 	var updatedPackages = new List<string>();

// 	foreach (var packageInfo in packageInfos)
// 	{
// 		var (package, updated) = await UpdatePackageAsync(packageInfo, progress, cancellationToken).ConfigureAwait(false);

// 		if (updated)
// 		{
// 			updatedPackages.Add(package);
// 		}
// 	}

// 	return updatedPackages;
// }
