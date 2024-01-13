package helpers

import (
	"Canteen-Backend/pkg/customErr"
	"crypto/sha1"
	"fmt"
	"os"
)

func HashPassword(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT")))), nil
}

func CheckPassword(password, hashedPassword string) error {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT")))) != hashedPassword {
		return customErr.PasswordInvalid
	}

	return nil
}
