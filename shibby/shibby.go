package shibby

func NewPackageData(packageInfos []PackageInfo) IPackageData {
	return &PackageData{packageInfos}
}
