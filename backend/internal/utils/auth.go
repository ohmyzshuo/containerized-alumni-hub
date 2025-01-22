package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ExtractLoginName(matricNo string) (string, error) {
	parts := strings.Split(matricNo, "/")
	log.Println("marticNo", matricNo)
	if len(parts) > 0 {
		return strings.ToLower(parts[0]), nil
	}
	return "", errors.New("invalid matric number format")
}
