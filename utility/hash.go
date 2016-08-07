package utility

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	result := ""
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return result, err
	}

	result = hex.EncodeToString(passwordBytes)
	return result, nil
}

func IsPasswordValid(hashedPassword, plaintextPassword string) bool {
	passwordBytes, _ := hex.DecodeString(hashedPassword)
	err := bcrypt.CompareHashAndPassword(passwordBytes, []byte(plaintextPassword))

	if err == nil {
		return true
	}

	return false
}
