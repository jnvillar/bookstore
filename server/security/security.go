package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	password := []byte(pwd)

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func ComparePassWords(inputPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(inputPassword)); err != nil {
		return false
	}
	return true
}
