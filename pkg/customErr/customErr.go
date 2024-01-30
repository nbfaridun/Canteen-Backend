package customErr

import (
	"errors"
)

// CustomError is a custom error type that contains an error, a message and a status code.
type CustomError struct {
	Error      error
	Message    string
	StatusCode int
}

func NewCustomError(err error, message string, statusCode int) *CustomError {
	return &CustomError{
		Error:      err,
		Message:    message,
		StatusCode: statusCode,
	}
}

var UsernameAlreadyExists = errors.New("username already exists")
var EmailAlreadyExists = errors.New("email already exists")
var IngredientCategoryAlreadyExists = errors.New("ingredient category already exists")
var IngredientAlreadyExists = errors.New("ingredient already exists")
var SupplierAlreadyExists = errors.New("supplier already exists")
var ClientCategoryAlreadyExists = errors.New("client category already exists")
var PurchaseAlreadyExists = errors.New("purchase already exists")

var PasswordInvalid = errors.New("password invalid")
var SessionExpired = errors.New("session expired")
var SessionNotFound = errors.New("session not found")
var RoleNotFound = errors.New("role not found")
var UserNotFound = errors.New("user not found")
var SupplierNotFound = errors.New("supplier not found")
var ClientCategoryNotFound = errors.New("client category not found")
var ClientNotFound = errors.New("client not found")
var IngredientCategoryNotFound = errors.New("ingredient category not found")
var IngredientNotFound = errors.New("ingredient not found")

var ServerError = errors.New("server error")
