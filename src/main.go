package main

import (
	"fmt"

	"github.com/mbodm/wingetupd-go/args"
	"github.com/mbodm/wingetupd-go/console"
)

const AppName = "wingetupd"
const AppVersion = "2.0.1"
const AppAuthor = "MBODM"
const AppDate = "2022-06-02"

func main() {
	fmt.Println()
	fmt.Println("%s %s (by %s %s)", AppName, AppVersion, AppAuthor, AppDate)
	fmt.Println()
	if !args.Validate() {
		console.ShowUsage(AppName, false)
		Exit(1)
	}
	if args.HelpExists() {
		console.ShowUsage(AppName, true)
		Exit(0)
	}
	// wgr, err := winget.Run("search --exact --id Mozilla.Firefox")
	// if err != nil {
	// 	if wgr.ExitCode != 0 {
	// 		fmt.Println("WinGet exitcode was", wgr.ExitCode)
	// 		fmt.Printf("Error: %v.", err)
	// 	} else {
	// 		fmt.Printf("Error was: %v", err)
	// 	}
	// 	return
	// }
	// fmt.Println()
	// fmt.Println("ProcessCall: ", wgr.ProcessCall)
	// fmt.Println()
	// fmt.Println("ConsoleOutput:")
	// fmt.Println(wgr.ConsoleOutput)
	// fmt.Println("ExitCode: ", wgr.ExitCode)
	// fmt.Println()
	// pr, err := parser.ParseListOutput(wgr.ConsoleOutput)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	// if pr.NewVersion == "" {
	// 	pr.NewVersion = "---"
	// }
	// fmt.Println("Old version: ", pr.OldVersion)
	// fmt.Println("New version: ", pr.NewVersion)
	// fmt.Println("Has update: ", pr.HasUpdate)
}
