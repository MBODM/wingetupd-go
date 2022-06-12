package models

// Infos

func (pd *PackageData) GetAll() []PackageInfo {
	return pd.all
}

func (pd *PackageData) GetValid() []PackageInfo {
	return pd.valid
}

func (pd *PackageData) GetInvalid() []PackageInfo {
	return pd.invalid
}

func (pd *PackageData) GetInstalled() []PackageInfo {
	return pd.installed
}

func (pd *PackageData) GetNonInstalled() []PackageInfo {
	return pd.nonInstalled
}

func (pd *PackageData) GetUpdatable() []PackageInfo {
	return pd.updatable
}

// Names

func (pd *PackageData) GetAllPackages() []string {
	return pd.allPackages
}

func (pd *PackageData) GetValidPackages() []string {
	return pd.validPackages
}

func (pd *PackageData) GetInvalidPackages() []string {
	return pd.invalidPackages
}

func (pd *PackageData) GetInstalledPackages() []string {
	return pd.installedPackages
}

func (pd *PackageData) GetNonInstalledPackages() []string {
	return pd.nonInstalledPackages
}

func (pd *PackageData) GetUpdatablePackages() []string {
	return pd.updatablePackages
}

// Counts

func (pd *PackageData) CountAll() int {
	return len(pd.all)
}

func (pd *PackageData) CountValid() int {
	return len(pd.valid)
}

func (pd *PackageData) CountInvalid() int {
	return len(pd.invalid)
}

func (pd *PackageData) CountInstalled() int {
	return len(pd.installed)
}

func (pd *PackageData) CountNonInstalled() int {
	return len(pd.nonInstalled)
}

func (pd *PackageData) CountUpdatable() int {
	return len(pd.updatable)
}

// Contains

func (pd *PackageData) ContainsValid() bool {
	return pd.validCount > 0
}

func (pd *PackageData) ContainsInvalid() bool {
	return pd.invalidCount > 0
}

func (pd *PackageData) ContainsInstalled() bool {
	return pd.installedCount > 0
}

func (pd *PackageData) ContainsNonInstalled() bool {
	return pd.nonInstalledCount > 0
}

func (pd *PackageData) ContainsUpdatable() bool {
	return pd.updatableCount > 0
}
