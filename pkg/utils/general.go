package utils

import (
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"time"
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

// Calculate a numeric hash of a string
func HashNumber(source string) string {
	h := fnv.New32a()
	h.Write([]byte(source))
	return fmt.Sprintf("%d", h.Sum32()+1000)[0:4]
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

// Returns a string list with the dates between the start and end dates, inclusive of both
// The dates must be in the format "2006-01-02"
func DatesInInterval(start string, end string) []string {
	startDate, err := time.Parse("2006-01-02 15:04:05", start+" 00:00:00")
	if err != nil {
		return nil
	}
	endDate, err := time.Parse("2006-01-02 15:04:05", end+" 23:59:59")
	if err != nil {
		return nil
	}
	if startDate.After(endDate) {
		return nil
	}
	var dates []string
	for startDate.Before(endDate) || startDate.Equal(endDate) {
		dates = append(dates, startDate.Format("2006-01-02"))
		startDate = startDate.AddDate(0, 0, 1)
	}
	return dates
}
