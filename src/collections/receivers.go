package collections

func (e *EvalResult) HasValidPackages() bool {
	return len(e.ValidPackages) > 1
}

func (e *EvalResult) HasInvalidPackages() bool {
	return len(e.InvalidPackages) > 1
}

func (e *EvalResult) HasInstalledPackages() bool {
	return len(e.InstalledPackages) > 1
}

func (e *EvalResult) HasNonInstalledPackages() bool {
	return len(e.NonInstalledPackages) > 1
}

func (e *EvalResult) HasUpdatablePackages() bool {
	return len(e.UpdatablePackages) > 1
}

func (e *EvalResult) HasPackageInfos() bool {
	return len(e.PackageInfos) > 1
}
