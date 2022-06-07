package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mbodm/wingetupd-go/app"
	"github.com/mbodm/wingetupd-go/args"
	"github.com/mbodm/wingetupd-go/collections"
	"github.com/mbodm/wingetupd-go/config"
	"github.com/mbodm/wingetupd-go/console"
	"github.com/mbodm/wingetupd-go/core"
)

func main() {
	fmt.Println()
	title := fmt.Sprintf("%s %s (by %s %s)", app.Name, app.Version, app.Author, app.Date)
	fmt.Println(title)
	fmt.Println()
	result, err := run()
	if err != nil {
		var appError *app.AppError
		if errors.As(err, &appError) {
			fmt.Println("Error: " + appError.Msg + ".")
		} else {
			fmt.Println("Unexpected error(s):", err)
		}
		os.Exit(1)
	}
	if !result {
		os.Exit(1)
	}
	fmt.Println("Have a nice day.")
	os.Exit(0)
}

func run() (bool, error) {
	if !args.Validate() {
		console.ShowUsage(app.Name, false)
		return false, nil
	}
	if args.HelpExists() {
		console.ShowUsage(app.Name, true)
		return false, nil
	}
	e := core.Init()
	if e != nil {
		return false, app.WrapError("main.run", e)
	}
	packages, err := config.ReadPackageFile()
	if err != nil {
		return false, app.WrapError("main.run", err)
	}
	console.ShowPackageFileEntries(packages)
	fmt.Println()
	fmt.Print("Processing ...")
	packageInfos, err := core.Analyze(packages, func() { fmt.Print(".") })
	if err != nil {
		return false, app.WrapError("main.run", err)
	}
	fmt.Println(" finished.")
	fmt.Println()
	evalResult := collections.Eval(packageInfos)
	if evalResult.HasInvalidPackages() {
		console.ShowInvalidPackagesError(evalResult.InvalidPackages)
		os.Exit(1)
	}
	if evalResult.HasNonInstalledPackages() {
		console.ShowNonInstalledPackagesError(evalResult.NonInstalledPackages)
		os.Exit(1)
	}
	console.ShowSummary(&evalResult)
	fmt.Println()
	return true, nil
}
