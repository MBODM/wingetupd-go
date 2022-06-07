package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/mbodm/wingetupd-go/app"
)

const pkgFileName = "packages.txt"

func ReadPackageFile() ([]string, error) {
	packages := []string{}
	pkgFile, err := getPkgFilePath()
	if err != nil {
		return packages, app.WrapError("config.ReadPackageFile", err)
	}
	file, err := os.Open(pkgFile)
	if err != nil {
		if os.IsNotExist(err) {
			return packages, app.NewAppError("The package-file not exists", err)
		}
		return packages, app.WrapError("config.ReadPackageFile", err)
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
		return packages, app.WrapError("config.ReadPackageFile", err)
	}
	return packages, nil
}
