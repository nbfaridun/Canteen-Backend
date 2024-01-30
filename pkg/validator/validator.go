package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

func ValidatePayload(payload interface{}) error {
	validate := validator.New()
	err := validate.RegisterValidation("any_digit", containsDigit)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("english_chars", containsOnlyEnglishLetters)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("any_special_char", containsSpecialChar)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("gender", validateGender)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("any_uppercase", containsUppercase)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("any_lowercase", containsLowercase)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("alphanumunicode_and_space", isAlphaNumSpace)

	var field, tag string
	err = validate.Struct(payload)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field = e.Field()
			tag = e.Tag()

			return errors.New(fmt.Sprintf("%s: %s", field, tag))
		}
	}

	return nil
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

func isAlphaNumSpace(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	for _, r := range field {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
