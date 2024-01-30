package helpers

import (
	"Canteen-Backend/pkg/customErr"
	"crypto/sha1"
	"fmt"
	"os"
	"time"
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

func ConvertStringToDate(date, layout string) time.Time {
	t, _ := time.Parse(layout, date)
	return t
}
