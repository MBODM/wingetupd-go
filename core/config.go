package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"example.com/mbodm/wingetupd/errs"
)

const packageFileName = "packages.txt"

func GetPackageFilePath() string {
	appDataPackageFile := filepath.Join(getAppDataPath(), packageFileName)
	exePathPackageFile := filepath.Join(getExeFilePath(), packageFileName)
	if fileExists(appDataPackageFile) {
		return appDataPackageFile
	}
	if fileExists(exePathPackageFile) {
		return exePathPackageFile
	}
	return ""
}

func PackageFileExists() bool {
	return GetPackageFilePath() != ""
}

func ReadPackageFile() ([]string, error) {
	packages := []string{}
	packageFile := GetPackageFilePath()
	notExistsErrMsg := "Could not found package-file"
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

func getExeFilePath() string {
	exe, err := os.Executable()
	if err != nil {
		errs.Fatal(err)
	}
	return filepath.Dir(exe)
}

func getAppDataPath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		errs.Fatal(err)
	}
	return dir
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		errs.Fatal(err)
	}
	return true
}

func handleBOM(s string) string {
	// If the file is a UTF-8 text file with BOM, like Windows Notepad does,
	// skip BOM. Text files with BOM have "\ufeff" as their first text char.
	if strings.Contains(s, "\uFEFF") {
		_, i := utf8.DecodeRuneInString(s)
		return s[i:]
	}
	return s
}
