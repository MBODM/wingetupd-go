package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/mbodm/wingetupd-go/errs"
)

const packageFileName = "packages.txt"

func PackageFileExists() bool {
	return getPackageFilePath() != ""
}

func ReadPackageFile() ([]string, error) {
	packages := []string{}
	packageFile := getPackageFilePath()
	notExistsErrMsg := "Package-file not exists"
	if packageFile == "" {
		return packages, errs.NewExpectedError(notExistsErrMsg, nil)
	}
	file, err := os.Open(packageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return packages, errs.NewExpectedError(notExistsErrMsg, err)
		}
		return packages, errs.NewExpectedError("Could not open package-file", err)
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
		return packages, errs.NewExpectedError("Unknown problem while reading package-file", err)
	}
	return packages, nil
}
