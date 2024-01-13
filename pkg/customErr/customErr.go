package customErr

import (
	"errors"
)

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Error      error
	Message    string
	StatusCode int
}

var UsernameAlreadyExists = errors.New("username already exists")
var EmailAlreadyExists = errors.New("email already exists")
var PasswordInvalid = errors.New("password invalid")
var SessionExpired = errors.New("session expired")
var SessionNotFound = errors.New("session not found")

var RoleNotFound = errors.New("role not found")
var UserNotFound = errors.New("user not found")
var ClientCategoryNotFound = errors.New("client category not found")
var ClientNotFound = errors.New("client not found")

var ServerError = errors.New("server error")
