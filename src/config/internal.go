package config

import (
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/mbodm/wingetupd-go/app"
)

func getPkgFilePath() (string, error) {
	exeFile, err := os.Executable()
	if err != nil {
		return "", app.WrapError("config.getPkgFilePath", err)
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
