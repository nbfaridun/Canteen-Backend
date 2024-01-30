package request

import "Canteen-Backend/internal/models"

type SignIn struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type CreateUser struct {
	Username   string `json:"username" validate:"required,min=4,max=20,alphanum"`
	UserRoleID uint   `json:"user_role_id" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8,max=20,any_uppercase,any_lowercase,any_digit,english_chars,any_special_char"`
	FirstName  string `json:"first_name" validate:"required,min=1,max=20,alpha"`
	LastName   string `json:"last_name" validate:"required,min=1,max=20,alpha"`
}

type UpdateUser struct {
	Username   string `json:"username" validate:"omitempty,min=4,max=20,alphanum"`
	UserRoleID uint   `json:"user_role_id"`
	Email      string `json:"email" validate:"omitempty,email"`
	Password   string `json:"password" validate:"omitempty,min=8,max=20,any_digit,any_uppercase,any_lowercase,english_chars,any_special_char"`
	FirstName  string `json:"first_name" validate:"omitempty,min=1,max=20,alpha"`
	LastName   string `json:"last_name" validate:"omitempty,min=1,max=20,alpha"`
	IsActive   bool   `json:"is_active"`
}

func MapSignInToUser(input *SignIn) *models.User {
	return &models.User{
		Username: input.Username,
		Password: input.Password,
	}
}

func MapCreateUserToUser(input *CreateUser) *models.User {
	return &models.User{
		Username:   input.Username,
		Email:      input.Email,
		Password:   input.Password,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		UserRoleID: input.UserRoleID,
		IsActive:   true,
	}
}

func MapUpdateUserToUser(input *UpdateUser) *models.User {
	return &models.User{
		Username:   input.Username,
		Email:      input.Email,
		Password:   input.Password,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		UserRoleID: input.UserRoleID,
		IsActive:   input.IsActive,
	}
}
