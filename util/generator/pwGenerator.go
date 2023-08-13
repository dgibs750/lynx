package generator

import (
	"crypto/rand"
	"math/big"
)

const (
	lowerCaseLetters  = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters = "~!@#$%^&*+`-=,."
	digits            = "0123456789"
)

func GeneratePassword(length int, includeLower bool, includeUpper bool, includeSpecial bool, includeDig bool, numDigits int, numSpecial int) (string, error) {
	var password string
	var letters string

	char := length - numDigits - numSpecial

	if includeLower {
		letters += lowerCaseLetters
	}
	if includeUpper {
		letters += upperCaseLetters
	}
	if len(letters) == 0 {
		letters += lowerCaseLetters
	}

	// Generates random letters for the password based on length - the amount of numbers and special characters
	for i := 0; i < char; i++ {
		val, err := randomValue(letters)
		if err != nil {
			return "", err
		}
		password, err = randomPlace(password, val)
		if err != nil {
			return "", err
		}
	}

	// Generates random numbers
	for i := 0; i < numDigits; i++ {
		num, err := randomValue(digits)
		if err != nil {
			return "", err
		}
		password, err = randomPlace(password, num)
		if err != nil {
			return "", nil
		}
	}

	// Generates random special characters
	for i := 0; i < numSpecial; i++ {
		specialChar, err := randomValue(specialCharacters)
		if err != nil {
			return "", nil
		}
		password, err = randomPlace(password, specialChar)
		if err != nil {
			return "", nil
		}
	}

	return password, nil
}

func randomPlace(s string, val string) (string, error) {
	if s == "" {
		return val, nil
	}
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(s)+1)))
	if err != nil {
		return "", err
	}
	i := num.Int64()
	return s[0:i] + val + s[i:], nil
}

func randomValue(s string) (string, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(s))))
	if err != nil {
		return "", err
	}
	return string(s[num.Int64()]), nil
}
