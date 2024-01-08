package customValidations

import (
	"Canteen-Backend/customErr"
	"Canteen-Backend/internal/models"
)

func ValidateCreateClient(client *models.Client) *customErr.CustomError {

	if err := ValidateFirstName(client.FirstName); err != nil {
		return err
	}

	if err := ValidateLastName(client.LastName); err != nil {
		return err
	}

	if err := ValidateEmail(client.Email); err != nil {
		return err
	}

	if err := ValidateAge(client.Age); err != nil {
		return err
	}

	if err := ValidateGender(client.Gender); err != nil {
		return err
	}

	if client.Balance < 0 {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("balance"),
			FrontendMessage: customErr.FieldContentNotValid("balance").Error(),
			StatusCode:      400,
		}
	}

	return nil
}

func ValidateUpdateClient(client *models.Client) *customErr.CustomError {
	if client.FirstName != "" {
		if err := ValidateUsername(client.FirstName); err != nil {
			return err
		}
	}
	if client.LastName != "" {
		if err := ValidateUsername(client.LastName); err != nil {
			return err
		}
	}
	if client.Email != "" {
		if err := ValidateUsername(client.Email); err != nil {
			return err
		}
	}
	if client.Age != 0 {
		if err := ValidateAge(client.Age); err != nil {
			return err
		}
	}
	if client.Balance != 0 {
		if client.Balance < 0 {
			return &customErr.CustomError{
				LogError:        customErr.FieldContentNotValid("balance"),
				FrontendMessage: customErr.FieldContentNotValid("balance").Error(),
				StatusCode:      400,
			}
		}
	}

	return nil
}
