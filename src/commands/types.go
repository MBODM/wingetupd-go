package commands

// Using golang´s "promoted fields" feature here.
// But not promoting structs from other packages.
// I´m unsure if i like such direct dependencies.

type Basics struct {
	Package             string
	WinGetProcessCall   string
	WinGetConsoleOutput string
	WinGetExitCode      int
}

type SearchResult struct {
	Basics
	IsValid bool
}

type ListResult struct {
	Basics
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

type UpgradeResult struct {
	Basics
	SuccessfullyUpdated bool
}
