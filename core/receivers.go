package core

func (e *PackageData) HasValidPackages() bool {
	return len(e.ValidPackages) > 1
}

func (e *PackageData) HasInvalidPackages() bool {
	return len(e.InvalidPackages) > 1
}

func (e *PackageData) HasInstalledPackages() bool {
	return len(e.InstalledPackages) > 1
}

func (e *PackageData) HasNonInstalledPackages() bool {
	return len(e.NonInstalledPackages) > 1
}

func (e *PackageData) HasUpdatablePackages() bool {
	return len(e.UpdatablePackages) > 1
}

func (e *PackageData) HasUpdatablePackageInfos() bool {
	return len(e.UpdatablePackageInfos) > 1
}

func (e *PackageData) HasPackageInfos() bool {
	return len(e.PackageInfos) > 1
}
