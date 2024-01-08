package usecase

import (
	"Canteen-Backend/customErr"
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/internal/usecase/customValidations"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserUseCase struct {
	repo repository.User
}

func NewUserUseCase(repo repository.User) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(user *models.User) (uint, *customErr.CustomError) {

	// User Credentials Validation for CreateUser
	if err := customValidations.ValidateCreateUser(user); err != nil {
		return 0, err
	}

	if _, err := u.repo.GetRoleByID(user.RoleID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.RoleNotFound.Error(),
				StatusCode:      http.StatusNotFound,
			}
		} else {
			return 0, &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.ServerError.Error(),
				StatusCode:      http.StatusInternalServerError,
			}
		}

	}

	id, err := u.repo.CreateUser(user)
	if err != nil {
		if ok, columnName := customErr.IsDuplicateKeyError(err); ok {
			switch columnName {
			case "username":
				return 0, &customErr.CustomError{
					LogError:        err,
					FrontendMessage: customErr.UsernameAlreadyExists.Error(),
					StatusCode:      http.StatusConflict,
				}
			case "email":
				return 0, &customErr.CustomError{
					LogError:        err,
					FrontendMessage: customErr.EmailAlreadyExists.Error(),
					StatusCode:      http.StatusConflict,
				}
			}
		} else {
			return 0, &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.ServerError.Error(),
				StatusCode:      http.StatusInternalServerError,
			}
		}
	}

	return id, nil
}

func (u *UserUseCase) GetAllUsers() (*[]models.User, *customErr.CustomError) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		return nil, &customErr.CustomError{
			LogError:        err,
			FrontendMessage: customErr.ServerError.Error(),
			StatusCode:      http.StatusInternalServerError,
		}
	}

	return users, nil
}

func (u *UserUseCase) GetUserByID(userID uint) (*models.User, *customErr.CustomError) {
	user, err := u.repo.GetUserByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.UserNotFound.Error(),
				StatusCode:      http.StatusNotFound,
			}
		}
		return nil, &customErr.CustomError{
			LogError:        err,
			FrontendMessage: customErr.ServerError.Error(),
			StatusCode:      http.StatusInternalServerError,
		}
	}

	return user, nil
}

func (u *UserUseCase) UpdateUser(id uint, user *models.User) *customErr.CustomError {

	// User Credentials Validation for UpdateUser
	if err := customValidations.ValidateUpdateUser(user); err != nil {
		return err
	}

	if user.RoleID != 0 {
		if _, err := u.repo.GetRoleByID(user.RoleID); err != nil {
			if err == gorm.ErrRecordNotFound {
				return &customErr.CustomError{
					LogError:        err,
					FrontendMessage: customErr.RoleNotFound.Error(),
					StatusCode:      http.StatusNotFound,
				}
			} else {
				return &customErr.CustomError{
					LogError:        err,
					FrontendMessage: customErr.ServerError.Error(),
					StatusCode:      http.StatusInternalServerError,
				}
			}
		}
	}

	user.UpdatedAt = time.Now()

	if err := u.repo.UpdateUser(id, user); err != nil {
		if ok, columnName := customErr.IsDuplicateKeyError(err); ok {
			switch columnName {
			case "username":
				return &customErr.CustomError{
					LogError:        err,
					FrontendMessage: customErr.UsernameAlreadyExists.Error(),
					StatusCode:      http.StatusConflict,
				}
			case "email":
				return &customErr.CustomError{
					LogError:        err,
					FrontendMessage: customErr.EmailAlreadyExists.Error(),
					StatusCode:      http.StatusConflict,
				}
			}
		} else if err == gorm.ErrRecordNotFound {
			return &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.UserNotFound.Error(),
				StatusCode:      http.StatusNotFound,
			}
		} else {
			return &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.ServerError.Error(),
				StatusCode:      http.StatusInternalServerError,
			}
		}
	}

	return nil
}

func (u *UserUseCase) DeleteUser(id uint) *customErr.CustomError {

	err := u.repo.DeleteUser(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.UserNotFound.Error(),
				StatusCode:      http.StatusNotFound,
			}
		} else {
			return &customErr.CustomError{
				LogError:        err,
				FrontendMessage: customErr.ServerError.Error(),
				StatusCode:      http.StatusInternalServerError,
			}
		}
	}

	return nil
}
