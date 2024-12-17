package utils

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"golang.org/x/exp/rand"
)

// EncodeSha256
func EncodeSha256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", hash)
}

// MakePassword
func GeneratePassword(password string, salt string) string {
	return EncodeSha256(password + salt)
}

// GenerateSalt
func GenerateSalt() string {
	return strconv.Itoa(rand.Intn(1000000))
}

// CheckPassword
func CheckPassword(password string, salt string, hash string) bool {
	return GeneratePassword(password, salt) == hash
}
