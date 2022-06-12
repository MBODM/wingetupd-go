package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const PackageFileName = "packages.txt"

func GetPackageFilePath(panicHandler func(error)) string {
	appDataPackageFile := filepath.Join(getAppDataPath(panicHandler), PackageFileName)
	exePathPackageFile := filepath.Join(getExeFilePath(panicHandler), PackageFileName)
	if fileExists(appDataPackageFile, panicHandler) {
		return appDataPackageFile
	}
	if fileExists(exePathPackageFile, panicHandler) {
		return exePathPackageFile
	}
	return ""
}

func PackageFileExists(panicHandler func(error)) bool {
	return GetPackageFilePath(panicHandler) != ""
}

func ReadPackageFile(panicHandler func(error)) ([]string, error) {
	// Todo: Move this comment in some common section or the README file!
	// Since golang has nothing like C#´s "nameof" operator and reflection is too
	// much bucks for the bang, hardcoded string seems ok, even when a bit filthy.
	const caller = "ReadPackageFile"
	const notFound = "could not found package-file"
	packageFilePath := GetPackageFilePath(panicHandler)
	if packageFilePath == "" {
		return nil, createError(caller, notFound)
	}
	file, err := os.Open(packageFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, wrapError(caller, notFound, err)
		}
		return nil, wrapError(caller, "could not open package-file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	packages := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = handleBOM(line)
		line = strings.TrimSpace(line)
		if line != "" {
			packages = append(packages, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, wrapError(caller, "could not read package-file", err)
	}
	return packages, nil
}
