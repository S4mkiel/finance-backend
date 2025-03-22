package utils

import (
	"bytes"
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

func CryptPassword(password *string) (*string, error) {
	hashedInput := sha256.Sum256([]byte(*password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(preparedPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	bcryptPasswordString := string(bcryptPassword)

	return PString(bcryptPasswordString), nil
}

func CompareHash(password, hash *string) bool {
	if password == nil || hash == nil {
		return false
	}

	hashedInput := sha256.Sum256([]byte(*password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)

	plainTextInBytes := []byte(preparedPassword)
	hasTextInBytes := []byte(*hash)

	err := bcrypt.CompareHashAndPassword(hasTextInBytes, plainTextInBytes)
	if err != nil {
		return false
	} else {
		return true
	}
}
