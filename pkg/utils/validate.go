package utils

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	// check if all numbers are the same
	for i := 0; i < len(cpf)-1; i++ {
		if cpf[i] != cpf[i+1] {
			break
		}
		if i == len(cpf)-2 {
			return false
		}
	}
	if allCaractersEqual(cpf) {
		return false
	}
	// check first digit
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	rest := sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}
	if rest != int(cpf[9]-'0') {
		return false
	}
	// check second digit
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	rest = sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}
	if rest != int(cpf[10]-'0') {
		return false
	}
	return true
}

func allCaractersEqual(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1] {
			return false
		}
	}
	return true
}

func ValidateUUID(subject string) bool {
	_, err := uuid.Parse(subject)
	return err == nil
}

func ValidateTimestamp(timestamp string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", timestamp)
	return err == nil
}

func ValidateDDD(ddd string) bool {
	if len(ddd) != 2 {
		return false
	}
	if HasOnlyNumbers(ddd) {
		return true
	}
	return false
}

func ValidatePhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) > PHONENUMBER_MAX_LENGTH {
		return false
	}
	if len(phoneNumber) < PHONENUMBER_MIN_LENGTH {
		return false
	}
	if HasOnlyNumbers(phoneNumber) {
		return true
	}
	return false
}

func HasOnlyNumbers(str string) bool {
	for _, c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func ValidateCellphoneNumber(phoneNumber string) error {
	if len(phoneNumber) != 11 {
		return errors.New("número de celular inválido")
	}
	if phoneNumber[2] != '9' {
		return errors.New("número de celular inválido")
	}
	return nil
}

func ValidateLandlineNumber(phoneNumber string) error {
	if len(phoneNumber) != 10 {
		return errors.New("número de telefone fixo inválido")
	}
	if phoneNumber[2] == '9' || phoneNumber[2] == '8' || phoneNumber[2] == '7' {
		return errors.New("número de telefone fixo inválido")
	}
	return nil
}

func ValidateEmail(email string) bool {
	if (len(email) < 5) || (len(email) > 254) {
		return false
	}
	if !strings.Contains(email, ".") {
		return false
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	match, err := regexp.Compile(VALID_EMAIL_REGEX)
	if err != nil {
		return false
	}
	if !match.MatchString(email) {
		return false
	}
	return true
}

func ValidateURL(link string) bool {
	_, err := url.ParseRequestURI(link)
	return err == nil
}
