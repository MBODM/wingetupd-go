package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/mbodm/wingetupd-go/errs"
)

const pkgFileName = "packages.txt"

func PackageFileExists() (bool, error) {
	path, err := getPackageFilePath()
	if err != nil {
		return false, errs.WrapError("config.PackageFileExists", err)
	}
	result, err := fileExists(path)
	if err != nil {
		return false, errs.WrapError("config.PackageFileExists", err)
	}
	return result, nil
}

func ReadPackageFile() ([]string, error) {
	packages := []string{}
	pkgFile, err := getPkgFilePath()
	if err != nil {
		return packages, errs.WrapError("config.ReadPackageFile", err)
	}
	file, err := os.Open(pkgFile)
	if err != nil {
		if os.IsNotExist(err) {
			return packages, errs.NewExpectedError("Package-file not exists", err)
		}
		return packages, errs.WrapError("config.ReadPackageFile", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = handleBOM(line)
		line = strings.TrimSpace(line)
		if line != "" {
			packages = append(packages, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return packages, errs.WrapError("config.ReadPackageFile", err)
	}
	return packages, nil
}
