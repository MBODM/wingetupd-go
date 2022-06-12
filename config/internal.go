package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func getExeFilePath(panicHandler func(error)) string {
	exe, err := os.Executable()
	if err != nil {
		unrecoverableError(panicHandler, err)
	}
	return filepath.Dir(exe)
}

func getAppDataPath(panicHandler func(error)) string {
	dir, err := os.UserCacheDir()
	if err != nil {
		unrecoverableError(panicHandler, err)
	}
	return dir
}

func fileExists(filePath string, panicHandler func(error)) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		unrecoverableError(panicHandler, err)
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
