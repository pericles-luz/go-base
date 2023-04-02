package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
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
	directory = strings.TrimPrefix(directory, "/")
	directory = strings.TrimSuffix(directory, "/")
	path += string(filepath.Separator) + directory
	return path
}

func getBaseDirectoryOnTesting() string {
	base, err := os.Getwd()
	if err != nil {
		log.Println("Erro ao obter diretorio atual", err)
		return ""
	}
	for !(FileExists(base+"/config") || strings.HasSuffix(base, "/go")) && len(base) > 1 {
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

func MD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
