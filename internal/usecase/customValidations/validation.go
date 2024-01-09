package customValidations

import (
	"Canteen-Backend/customErr"
	"github.com/asaskevich/govalidator"
	"net/http"
	"regexp"
)

func CheckFieldContentLength(fieldContent, fieldName string, min, max int) *customErr.CustomError {
	if len(fieldContent) <= min {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentTooShort(fieldName),
			FrontendMessage: customErr.FieldContentTooShort(fieldName).Error(),
			StatusCode:      http.StatusBadRequest,
		}
	} else if len(fieldContent) >= max {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentTooLong(fieldName),
			FrontendMessage: customErr.FieldContentTooLong(fieldName).Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}
	return nil
}

func CheckFieldContentEmpty(fieldContent, fieldName string) *customErr.CustomError {
	if fieldContent == "" {
		return &customErr.CustomError{
			LogError:        customErr.FieldRequired(fieldName),
			FrontendMessage: customErr.FieldRequired(fieldName).Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateSecondPass(fieldContent, secondFieldContent, fieldName string) *customErr.CustomError {
	if fieldContent != secondFieldContent {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid(fieldName),
			FrontendMessage: customErr.FieldContentNotValid(fieldName).Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateUsername(username string) *customErr.CustomError {
	if err := CheckFieldContentLength(username, "username", 4, 20); err != nil {
		return err
	}

	regUsernamePattern := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	if !regUsernamePattern.MatchString(username) {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("username"),
			FrontendMessage: customErr.FieldContentNotValid("username").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidatePassword(password string) *customErr.CustomError {

	if err := CheckFieldContentLength(password, "password", 8, 20); err != nil {
		return err
	}
	uppercaseRegex := regexp.MustCompile("[A-Z]")
	if !uppercaseRegex.MatchString(password) {
		return &customErr.CustomError{
			LogError:        customErr.PasswordNotContainsUpperChar,
			FrontendMessage: customErr.PasswordNotContainsUpperChar.Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}
	lowercaseRegex := regexp.MustCompile("[a-z]")
	if !lowercaseRegex.MatchString(password) {
		return &customErr.CustomError{
			LogError:        customErr.PasswordNotContainsLowerChar,
			FrontendMessage: customErr.PasswordNotContainsLowerChar.Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}
	digitRegex := regexp.MustCompile("[0-9]")
	if !digitRegex.MatchString(password) {
		return &customErr.CustomError{
			LogError:        customErr.PasswordNotContainsNumber,
			FrontendMessage: customErr.PasswordNotContainsNumber.Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}
	specialCharRegex := regexp.MustCompile("[!@#$%^&*(),.?\":{}|<>]")
	if !specialCharRegex.MatchString(password) {
		return &customErr.CustomError{
			LogError:        customErr.PasswordNotContainsSpecialChar,
			FrontendMessage: customErr.PasswordNotContainsSpecialChar.Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}
	validCharRegex := regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*(),.?\":{}|<>]+$")
	if !validCharRegex.MatchString(password) {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("password"),
			FrontendMessage: customErr.FieldContentNotValid("password").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}
	return nil
}

func ValidateEmail(email string) *customErr.CustomError {
	if !govalidator.IsEmail(email) {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("email"),
			FrontendMessage: customErr.FieldContentNotValid("email").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateFirstName(firstName string) *customErr.CustomError {
	if err := CheckFieldContentLength(firstName, "first_name", 1, 20); err != nil {
		return err
	}
	if !govalidator.IsAlpha(firstName) {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("first_name"),
			FrontendMessage: customErr.FieldContentNotValid("first_name").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateLastName(lastName string) *customErr.CustomError {
	if err := CheckFieldContentLength(lastName, "last_name", 1, 20); err != nil {
		return err
	}
	if !govalidator.IsAlpha(lastName) {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("last_name"),
			FrontendMessage: customErr.FieldContentNotValid("last_name").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateAge(age uint) *customErr.CustomError {
	if age >= 100 && age <= 0 {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("age"),
			FrontendMessage: customErr.FieldContentNotValid("age").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}

func ValidateGender(gender string) *customErr.CustomError {
	if !(gender == "male" || gender == "female") {
		return &customErr.CustomError{
			LogError:        customErr.FieldContentNotValid("gender"),
			FrontendMessage: customErr.FieldContentNotValid("gender").Error(),
			StatusCode:      http.StatusBadRequest,
		}
	}

	return nil
}
