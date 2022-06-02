package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
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

func getPkgFilePath() (string, error) {
	exeFile, err := os.Executable()
	if err != nil {
		return "", errors.New("todo") // todo: chain
	}
	exePath := filepath.Dir(exeFile)
	pkgFile := filepath.Join(exePath, pkgFileName)
	return pkgFile, nil
}

func handleBOM(s string) string {
	// If the file is a UTF-8 text file with BOM, like Windows Notepad does,
	// skip BOM. Text files with BOM have "\ufeff" as their first text char.
	if strings.Contains(s, "\ufeff") {
		_, i := utf8.DecodeRuneInString(s)
		return s[i:]
	}
	return s
}
