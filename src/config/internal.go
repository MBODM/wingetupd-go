package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/mbodm/wingetupd-go/errs"
)

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

func getPackageFilePath() string {
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

func handleBOM(s string) string {
	// If the file is a UTF-8 text file with BOM, like Windows Notepad does,
	// skip BOM. Text files with BOM have "\ufeff" as their first text char.
	if strings.Contains(s, "\uFEFF") {
		_, i := utf8.DecodeRuneInString(s)
		return s[i:]
	}
	return s
}
