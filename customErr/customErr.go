package customErr

import (
	"errors"
)

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	LogError        error
	FrontendMessage string
	StatusCode      int
}

var UsernameAlreadyExists = errors.New("username already exists")
var EmailAlreadyExists = errors.New("email already exists")

var RoleNotFound = errors.New("role not found")
var UserNotFound = errors.New("user not found")
var ClientCategoryNotFound = errors.New("client category not found")
var ClientNotFound = errors.New("client not found")

var FieldContentTooShort = func(fieldName string) error { return errors.New(fieldName + " too short") }
var FieldContentTooLong = func(fieldName string) error { return errors.New(fieldName + " too long") }
var FieldRequired = func(fieldName string) error { return errors.New(fieldName + " is required") }
var FieldContentNotValid = func(fieldName string) error { return errors.New(fieldName + " not valid") }

var PasswordNotContainsSpecialChar = errors.New("password not contains special char")
var PasswordNotContainsNumber = errors.New("password not contains number")
var PasswordNotContainsUpperChar = errors.New("password not contains upper char")
var PasswordNotContainsLowerChar = errors.New("password not contains lower char")

var ServerError = errors.New("server error")
