package commands

// Using golang´s "promoted fields" feature here.

type basics struct {
	Package             string
	WinGetProcessCall   string
	WinGetConsoleOutput string
	WinGetExitCode      int
}

type SearchResult struct {
	basics
	IsValid bool
}

type ListResult struct {
	basics
	IsInstalled      bool
	IsUpdatable      bool
	InstalledVersion string
	UpdateVersion    string
}

type UpgradeResult struct {
	basics
	SuccessfullyUpdated bool
}
