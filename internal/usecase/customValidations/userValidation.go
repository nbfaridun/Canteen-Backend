package customValidations

import (
	"Canteen-Backend/customErr"
	"Canteen-Backend/internal/models"
)

func ValidateCreateUser(user *models.User) *customErr.CustomError {
	if err := ValidateUsername(user.Username); err != nil {
		return err
	}

	if err := ValidatePassword(user.Password); err != nil {
		return err
	}

	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := ValidateFirstName(user.FirstName); err != nil {
		return err
	}

	if err := ValidateLastName(user.LastName); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateUser(user *models.User) *customErr.CustomError {
	if user.Username != "" {
		if err := ValidateUsername(user.Username); err != nil {
			return err
		}
	}
	if user.Password != "" {
		if err := ValidateUsername(user.Password); err != nil {
			return err
		}
	}
	if user.Email != "" {
		if err := ValidateUsername(user.Email); err != nil {
			return err
		}
	}
	if user.FirstName != "" {
		if err := ValidateUsername(user.FirstName); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := ValidateUsername(user.LastName); err != nil {
			return err
		}
	}

	return nil
}
