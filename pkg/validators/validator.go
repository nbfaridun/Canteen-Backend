package validators

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

func ValidatePayload(payload interface{}) error {
	validate := validator.New()
	err := validate.RegisterValidation("digit", containsDigit)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("englishchars", containsOnlyEnglishLetters)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("specialchar", containsSpecialChar)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("gender", validateGender)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("uppercase", containsUppercase)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("lowercase", containsLowercase)
	if err != nil {
		return err
	}

	var field, param, tag string
	err = validate.Struct(payload)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field = e.Field()
			tag = e.Tag()
			param = e.Param()

			return errors.New(handleError(field, tag, param))
		}
	}

	return nil
}

func handleError(field, tag, param string) string {
	field = strings.ToLower(field)
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required field", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only english letters and digits", field)
	case "alpha":
		return fmt.Sprintf("%s must contain only english letters", field)
	case "email":
		return "email address is not valid"
	case "lowercase":
		return fmt.Sprintf("%s must contain at least one lowercase letter", field)
	case "uppercase":
		return fmt.Sprintf("%s must contain at least one uppercase letter", field)
	case "digit":
		return fmt.Sprintf("%s must contain at least one digit", field)
	case "englishchars":
		return fmt.Sprintf("%s must contain only english letters", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, param)
	case "max":
		return fmt.Sprintf("%s must be less than %s characters long", field, param)
	case "specialchar":
		return fmt.Sprintf("%s must contain at least one special character of '!@#$%%^&*()?'", field)
	case "gender":
		return fmt.Sprintf("%s must be male of female", field)
	}
	return fmt.Sprintf("%s: %s %s", field, tag, param)
}

func containsDigit(fl validator.FieldLevel) bool {
	return strings.ContainsAny(fl.Field().String(), "0123456789")
}

func containsOnlyEnglishLetters(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	for _, r := range field {
		if unicode.IsLetter(r) && (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func containsSpecialChar(fl validator.FieldLevel) bool {
	return strings.ContainsAny(fl.Field().String(), "!@#$%^&*()?")
}

func validateGender(fl validator.FieldLevel) bool {
	return fl.Field().String() == "male" || fl.Field().String() == "female"
}

func containsUppercase(fl validator.FieldLevel) bool {
	return strings.ContainsAny(fl.Field().String(), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func containsLowercase(fl validator.FieldLevel) bool {
	return strings.ContainsAny(fl.Field().String(), "abcdefghijklmnopqrstuvwxyz")
}
