package response

import "Canteen-Backend/internal/models"

type GetUser struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	IsActive   bool   `json:"is_active"`
	UserRoleID uint   `json:"user_role"`
}

func MapUserToGetUser(user *models.User) *GetUser {
	return &GetUser{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		UserRoleID: user.UserRoleID,
	}
}
