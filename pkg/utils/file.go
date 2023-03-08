package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetBaseDirectory(directory string) string {
	directory = strings.ReplaceAll(directory, "..", "")
	path, err := os.Getwd()
	if err != nil {
		log.Println("Erro ao obter diretorio atual", err)
		return ""
	}
	if isTesting() {
		path = getBaseDirectoryOnTesting()
	}
	if directory == "" {
		return path
	}
	if directory[0] == '/' {
		directory = directory[1:]
	}
	if directory[len(directory)-1] == '/' {
		directory = directory[:len(directory)-1]
	}
	path += string(filepath.Separator) + directory
	return path
}

func getBaseDirectoryOnTesting() string {
	base, err := os.Getwd()
	if err != nil {
		log.Println("Erro ao obter diretorio atual", err)
		return ""
	}
	for !strings.HasSuffix(filepath.Dir(base), "/go") && len(base) > 1 {
		base = filepath.Dir(base)
	}
	return base
}

func FileExists(path string) bool {
	if path == "" {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
