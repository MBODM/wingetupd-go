package core

type PackageInfo struct {
	Package          string
	IsValid          bool
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}
