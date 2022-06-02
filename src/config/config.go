package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const packageFile = "packages.txt"

func PackageFileExists() (bool, error) {
	exePath, err := os.Executable()
	if err != nil {
		return false, errors.New("todo") // todo: chain
	}
	pkgFile := filepath.Join(exePath, packageFile)
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
	file, err := os.Open("packages.txt")
	if err != nil {
		return packages, errors.New("todo") // todo: chain
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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
