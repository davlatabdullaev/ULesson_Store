package check

import (
	"errors"
	"time"
	"unicode"
)

func PhoneNumber(phone string) bool {
	for _, r := range phone {
		if r == '+' {
			continue
		} else if !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not correct for car!")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password length should be more than 6")
	}

	return nil
}
