package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt encrypts a password using bcrypt
func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CompareEncrypt compares a password with a hash
func CompareEncrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
