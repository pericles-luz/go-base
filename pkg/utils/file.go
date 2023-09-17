package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// Identify the root directory of the project. The default is the current one.
// If there is a running test, the root directory is the first one that contains
// the directory "config" or the first one inside the directory "go".
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

// Verify if a file or directory exists
func FileExists(path string) bool {
	if path == "" {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// Calculate the MD5 hash of a file
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

// Calculate the SHA256 hash of a file
func SHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Extract the file from a base64 string in a slice of bytes
// The base64 string must be in the format "data:{mimeType};base64,...."
// The file is saved in a temporary directory and the path is returned
// The file name is a UUID to avoid conflicts
func ExtractFileFromBase64(r io.Reader) (*os.File, error) {
	base64Data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	start, err := FileStartFromBase64(bytes.NewReader(base64Data))
	if err != nil {
		return nil, err
	}
	mimeType, err := MimeTypeFromBase64(bytes.NewReader(base64Data))
	if err != nil {
		return nil, err
	}
	fileName := uuid.NewString() + "." + strings.Split(mimeType, "/")[1]
	file, err := os.CreateTemp("", fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	preReader := bytes.NewReader(base64Data)
	discard := make([]byte, start)
	count, err := preReader.Read(discard)
	if err != nil {
		return nil, err
	}
	if count != start {
		return nil, errors.New("invalid base64")
	}
	reader := base64.NewDecoder(base64.StdEncoding, preReader)
	if _, err := io.Copy(file, reader); err != nil {
		return nil, err
	}
	return file, nil
}

func MimeTypeFromBase64(r io.Reader) (string, error) {
	buffer := make([]byte, 512)
	_, err := r.Read(buffer)
	if err != nil {
		return "", err
	}
	idx, err := FileStartFromBase64(bytes.NewReader(buffer))
	if err != nil {
		return "", err
	}
	if idx < 5 {
		return "", errors.New("invalid base64")
	}
	if bytes.Index(buffer, []byte("data:")) != 0 {
		return "", errors.New("invalid base64 start")
	}
	mimeType := string(buffer[5 : idx-8])
	return mimeType, err
}

func FileStartFromBase64(r io.Reader) (int, error) {
	buffer := make([]byte, 512)
	_, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}
	idx := bytes.Index(buffer, []byte(";base64,"))
	if idx == -1 {
		return 0, errors.New("invalid base64")
	}
	return idx + 8, nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("file not found")
	}
	defer file.Close()
	return io.ReadAll(file)
}

func FileCopy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	return err
}

func FileMimeType(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return ""
	}
	result := http.DetectContentType(buffer)
	// strip the charset
	if idx := strings.Index(result, ";"); idx != -1 {
		result = result[:idx]
	}
	return result
}

func FileExtension(path string) string {
	return filepath.Ext(path)
}

func FileSize(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return 0
	}
	return stat.Size()
}

// Generate base64 string from a file
// the string must start with "data:{mimeType};base64,"
// the file is read in chunks of 510 bytes because needs to be multiple of 3
func FileToBase64(path string) (string, error) {
	mimeType := FileMimeType(path)
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer file.Close()
	buffer := make([]byte, 510)
	base64String := "data:" + mimeType + ";base64,"
	for {
		count, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		base64String += base64.StdEncoding.EncodeToString(buffer[:count])
	}
	return base64String, nil
}

// Returns a list with the fullpath of all files in a directory
func GetFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("error walking path", path, err)
			return nil
		}
		if info == nil {
			return nil
		}
		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}
