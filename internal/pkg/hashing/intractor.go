package hashing

import "golang.org/x/crypto/bcrypt"

//Intractor Hash

func HashPassword(password string) (string, error) {

	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}
