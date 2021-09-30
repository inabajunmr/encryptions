package caesar

import (
	"errors"
)

// Encrypt by caesar
func Encrypt(v string, key int) (string, error) {
	result := []rune{}
	for _, r := range v {
		if isLowerCase(r) {
			result = append(result, encryptLowerCase(r, key))
			continue
		}

		if isUpperCase(r) {
			result = append(result, encryptUpperCase(r, key))
			continue
		}

		if isSpace(r) {
			result = append(result, r)
			continue
		}

		return "", errors.New("this method can encrypt only [a-zA-Z]")
	}

	return string(result), nil
}

// Decrypt by caesar
func Decrypt(v string, key int) (string, error) {
	result := []rune{}
	for _, r := range v {
		if isLowerCase(r) {
			result = append(result, decryptLowerCase(r, key))
			continue
		}

		if isUpperCase(r) {
			result = append(result, decryptUpperCase(r, key))
			continue
		}

		if isSpace(r) {
			result = append(result, r)
			continue
		}

		return "", errors.New("this method can encrypt only [a-zA-Z]")
	}

	return string(result), nil
}

func encryptLowerCase(v rune, key int) rune {
	sum := int(v) + key
	if sum < 122 {
		return rune(sum)
	}

	return rune((sum)%122 + 96)
}

func encryptUpperCase(v rune, key int) rune {
	sum := int(v) + key
	if sum < 90 {
		return rune(sum)
	}

	return rune((sum)%90 + 64)
}

func decryptLowerCase(v rune, key int) rune {
	diff := int(v) - key
	if diff > 96 {
		return rune(diff)
	}

	return rune(122 - (96 - diff))
}

func decryptUpperCase(v rune, key int) rune {
	diff := int(v) - key
	if diff > 64 {
		return rune(diff)
	}

	return rune(90 - (64 - diff))
}

func isLowerCase(v rune) bool {
	return v >= 97 && v <= 122
}

func isUpperCase(v rune) bool {
	return v >= 65 && v <= 90
}

func isSpace(v rune) bool {
	return v == 32
}
