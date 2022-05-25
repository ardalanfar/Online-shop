package hash

import (
	"golang.org/x/crypto/bcrypt"
)

//Intractor package hash

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}