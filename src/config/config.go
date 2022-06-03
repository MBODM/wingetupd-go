package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const pkgFileName = "packages.txt"

func PackageFileExists() (bool, error) {
	pkgFile, err := getPkgFilePath()
	if err != nil {
		return false, errors.New("todo") // todo: chain
	}
	_, err = os.Stat(pkgFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, errors.New("todo") // todo: chain
		}
	}
	return true, nil
}

func ReadPackageFile() ([]string, error) {
	packages := []string{}
	pkgFile, err := getPkgFilePath()
	if err != nil {
		return packages, errors.New("todo") // todo: chain
	}
	file, err := os.Open(pkgFile)
	if err != nil {
		return packages, errors.New("todo") // todo: chain
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
		return packages, errors.New("todo") // todo: chain
	}
	return packages, nil
}
