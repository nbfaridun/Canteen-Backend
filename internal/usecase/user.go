package usecase

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/pkg/auth"
	"Canteen-Backend/pkg/customErr"
	"Canteen-Backend/pkg/helpers"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserUseCase struct {
	userRepo    repository.User
	sessionRepo repository.Session
}

func NewUserUseCase(userRepo repository.User, sessionRepo repository.Session) *UserUseCase {
	return &UserUseCase{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (u *UserUseCase) SignIn(userInput *models.User) (*models.Token, *customErr.CustomError) {
	user, err := u.userRepo.GetUserByUsername(userInput.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.UserNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	if err := helpers.CheckPassword(userInput.Password, user.Password); err != nil {
		if errors.Is(err, customErr.PasswordInvalid) {
			return nil, customErr.NewCustomError(err, customErr.PasswordInvalid.Error(), http.StatusUnauthorized)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return u.createSession(user.ID, user.UserRoleID)
}

func (u *UserUseCase) SignOut(refreshToken string) *customErr.CustomError {
	session, err := u.sessionRepo.GetSessionByRefreshToken(refreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.SessionNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	if err := u.sessionRepo.DeleteSession(session); err != nil {
		return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (u *UserUseCase) RefreshTokens(refreshToken string) (*models.Token, *customErr.CustomError) {
	session, err := u.sessionRepo.GetSessionByRefreshToken(refreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.SessionNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	if err := u.sessionRepo.DeleteSession(session); err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, customErr.NewCustomError(err, customErr.SessionExpired.Error(), http.StatusUnauthorized)
	}

	user, err := u.userRepo.GetUserByID(session.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.UserNotFound.Error(), http.StatusNotFound)
		}
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return u.createSession(user.ID, user.UserRoleID)
}

func (u *UserUseCase) CreateUser(user *models.User) (uint, *customErr.CustomError) {
	if _, err := u.userRepo.GetRoleByID(user.UserRoleID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, customErr.NewCustomError(err, customErr.RoleNotFound.Error(), http.StatusNotFound)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}
	user.Password = hashedPassword

	id, err := u.userRepo.CreateUser(user)
	if err != nil {
		if ok, columnName := customErr.IsDuplicateKeyError(err); ok {
			switch columnName {
			case "username":
				return 0, customErr.NewCustomError(err, customErr.UsernameAlreadyExists.Error(), http.StatusConflict)
			case "email":
				return 0, customErr.NewCustomError(err, customErr.EmailAlreadyExists.Error(), http.StatusConflict)
			}
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}

func (u *UserUseCase) GetAllUsers() (*[]models.User, *customErr.CustomError) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return users, nil
}

func (u *UserUseCase) GetUserByID(userID uint) (*models.User, *customErr.CustomError) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.UserNotFound.Error(), http.StatusNotFound)
		}
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return user, nil
}

func (u *UserUseCase) UpdateUser(user *models.User) *customErr.CustomError {
	if user.UserRoleID != 0 {
		if _, err := u.userRepo.GetRoleByID(user.UserRoleID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return customErr.NewCustomError(err, customErr.RoleNotFound.Error(), http.StatusNotFound)
			} else {
				return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
			}
		}
	}

	user.UpdatedAt = time.Now()

	if err := u.userRepo.UpdateUser(user); err != nil {
		if ok, columnName := customErr.IsDuplicateKeyError(err); ok {
			switch columnName {
			case "username":
				return customErr.NewCustomError(err, customErr.UsernameAlreadyExists.Error(), http.StatusConflict)
			case "email":
				return customErr.NewCustomError(err, customErr.EmailAlreadyExists.Error(), http.StatusConflict)
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.UserNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *UserUseCase) DeleteUser(id uint) *customErr.CustomError {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.UserNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *UserUseCase) createSession(userID uint, userRoleID uint) (*models.Token, *customErr.CustomError) {
	accessToken, err := auth.GenerateAccessToken(userID, userRoleID)
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	refreshToken, expTime, err := auth.GenerateRefreshToken()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	session := &models.Session{
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    expTime,
	}

	err = u.sessionRepo.CreateSession(session)
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return &models.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
