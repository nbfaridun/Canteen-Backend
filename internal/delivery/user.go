package delivery

import (
	_ "Canteen-Backend/docs"
	"Canteen-Backend/internal/models"
	"Canteen-Backend/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON input
// @Tags users
// @Accept json
// @Produce json
// @Param input body models.CreateUserInput true "User object to be created"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var input *models.CreateUserInput
	if err := c.BindJSON(&input); err != nil {
		logger.GetLogger().Error("error while binding json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input JSON"})
		return
	}
	id, err := h.useCase.User.CreateUser(&models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		RoleID:    input.RoleID,
		IsActive:  true,
	})
	if err != nil {
		logger.GetLogger().Error("error while creating user", zap.Error(err.LogError))
		c.JSON(err.StatusCode, gin.H{"error": err.FrontendMessage})
		return
	}

	logger.GetLogger().Info("user created", zap.Uint("id", id))
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users available
// @ID get-all-users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User "Successful response"
// @Failure 500 {string} string
// @Router /users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	var users *[]models.User
	users, err := h.useCase.User.GetAllUsers()
	if err != nil {
		logger.GetLogger().Error("error while getting all users", zap.Error(err.LogError))
		c.JSON(err.StatusCode, gin.H{"error": err.FrontendMessage})
		return
	}

	logger.GetLogger().Info("all users retrieved")
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user based on ID
// @ID get-user-by-id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} models.User "Successful response"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.GetLogger().Error("error while converting id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, customErr := h.useCase.User.GetUserByID(uint(id))
	if customErr != nil {
		logger.GetLogger().Error("error while getting user by id", zap.Error(customErr.LogError))
		c.JSON(customErr.StatusCode, gin.H{"error": customErr.FrontendMessage})
		return
	}

	logger.GetLogger().Info("user retrieved", zap.Uint("id", user.ID))
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update the existing user
// @Description Update the existing user with the provided JSON input
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Param input body models.UpdateUserInput true "User object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var user *models.UpdateUserInput
	if err := c.BindJSON(&user); err != nil {
		logger.GetLogger().Error("error while binding json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user JSON"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.GetLogger().Error("error while converting id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customErr := h.useCase.User.UpdateUser(uint(id), &models.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsActive:  user.IsActive,
		RoleID:    user.RoleID,
	})
	if customErr != nil {
		logger.GetLogger().Error("error while updating user", zap.Error(customErr.LogError))
		c.JSON(customErr.StatusCode, gin.H{"error": customErr.FrontendMessage})
		return
	}

	logger.GetLogger().Info("user updated", zap.Uint("id", uint(id)))
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user based on ID
// @ID delete-user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.GetLogger().Error("error while converting id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customErr := h.useCase.User.DeleteUser(uint(id))
	if customErr != nil {
		logger.GetLogger().Error("error while deleting user", zap.Error(customErr.LogError))
		c.JSON(customErr.StatusCode, gin.H{"error": customErr.FrontendMessage})
		return
	}

	logger.GetLogger().Info("user deleted", zap.Uint("id", uint(id)))
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
