package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), err
}

func IsValidPassword(password string, hash string) (bool, error) {
	byteHash := []byte(hash)
	bytePass := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)

	if err != nil {
		return false, err
	}
	return true, nil
}
