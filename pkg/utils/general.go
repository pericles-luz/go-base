package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func isTesting() bool {
	currentPath, err := os.Getwd()
	if err != nil {
		return false
	}
	return strings.HasSuffix(currentPath, "_test")
}

func Hash256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
func GetOnlyNumbers(str string) string {
	var numbers string
	for _, char := range str {
		if char >= '0' && char <= '9' {
			numbers += string(char)
		}
	}
	return numbers
}

func CompleteWithZeros(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}
