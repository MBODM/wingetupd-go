package commands

// Functional dependencies

type WinGetListParser func(winGetListOutput string) (*WinGetListParserResult, error)

type WinGetListParserResult struct {
	OldVersion string
	NewVersion string
}

type WinGetRunner func(winGetParams string) (*WinGetRunnerResult, error)

type WinGetRunnerResult struct {
	ProcessCall   string
	ConsoleOutput string
	ExitCode      int
}

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
