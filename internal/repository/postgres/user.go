package postgres

import (
	"Canteen-Backend/internal/constants"
	"Canteen-Backend/internal/models"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user *models.User) (uint, error) {

	result := r.db.Table(constants.UserTableName).Create(user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *UserPostgres) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	result := r.db.Table(constants.UserTableName).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *UserPostgres) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.Table(constants.UserTableName).First(&user, "user_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserPostgres) UpdateUser(user *models.User) error {
	result := r.db.Table(constants.UserTableName).Model(&models.User{}).Where("user_id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserPostgres) DeleteUser(id uint) error {
	result := r.db.Table(constants.UserTableName).Delete(&models.User{}, "user_id = ?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserPostgres) GetRoleByID(id uint) (string, error) {
	var role models.UserRole
	result := r.db.Table(constants.RoleTableName).Select("name").First(&role, "user_role_id = ?", id)
	if result.Error != nil {
		return "", result.Error
	}

	return role.Name, nil
}

func (r *UserPostgres) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Table(constants.UserTableName).First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
