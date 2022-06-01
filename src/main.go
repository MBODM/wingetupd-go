package main

import (
	"fmt"

	"github.com/mbodm/wingetupd-go/parser"
	"github.com/mbodm/wingetupd-go/winget"
)

func main() {
	wgr, err := winget.Run("search --exact --id Mozilla.Firefox")
	if err != nil {
		if wgr.ExitCode != 0 {
			fmt.Println("WinGet exitcode was", wgr.ExitCode)
			fmt.Printf("Error: %v.", err)
		} else {
			fmt.Printf("Error was: %v", err)
		}
		return
	}
	fmt.Println()
	fmt.Println("ProcessCall: ", wgr.ProcessCall)
	fmt.Println()
	fmt.Println("ConsoleOutput:")
	fmt.Println(wgr.ConsoleOutput)
	fmt.Println("ExitCode: ", wgr.ExitCode)
	fmt.Println()
	pr, err := parser.ParseListOutput(wgr.ConsoleOutput)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if pr.NewVersion == "" {
		pr.NewVersion = "---"
	}
	fmt.Println("Old version: ", pr.OldVersion)
	fmt.Println("New version: ", pr.NewVersion)
	fmt.Println("Has update: ", pr.HasUpdate)
}
