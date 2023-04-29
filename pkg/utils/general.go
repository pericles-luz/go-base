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

// Calculate the SHA256 hash of a string
func Hash256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// Calculate the SHA256 hash of a byte sequence
func Hash256FromBytes(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// Returns all numbers from a string and only the numbers
func GetOnlyNumbers(str string) string {
	var numbers string
	for _, char := range str {
		if char >= '0' && char <= '9' {
			numbers += string(char)
		}
	}
	return numbers
}

// Complete a string with left zeros until it reaches the desired length
func CompleteWithZeros(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}
