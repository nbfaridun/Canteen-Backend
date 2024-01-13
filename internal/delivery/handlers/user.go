package handlers

import (
	_ "Canteen-Backend/docs"
	"Canteen-Backend/internal/delivery/dto/request"
	"Canteen-Backend/internal/delivery/dto/response"
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/usecase"
	"Canteen-Backend/pkg/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.userHandler.SignIn)
		auth.POST("/sign-out", h.userHandler.SignOut)
		auth.POST("/refresh-token", h.userHandler.RefreshToken)
	}

	users := api.Group("/users")
	{
		users.Use(h.authenticateUser)
		{
			users.POST("/", h.userHandler.CreateUser)
			users.GET("/", h.userHandler.GetAllUsers)
			users.GET("/:id", h.userHandler.GetUserByID)
			users.PUT("/:id", h.userHandler.UpdateUser)
			users.DELETE("/:id", h.userHandler.DeleteUser)
		}
	}
}

type UserHandler struct {
	userUseCase usecase.User
}

func NewUserHandler(userUseCase usecase.User) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// SignIn godoc
// @Summary Sign in a user
// @Description Sign in a user with the provided JSON input
// @Tags auth
// @Accept json
// @Produce json
// @Param input body request.SignIn true "User sign in object"
// @Success 200 {object} models.Token "Successful response"
// @Failure 400 {string} string "Invalid input JSON"
// @Failure 401 {string} string "Unauthorized"
// @Router /api/auth/sign-in [post]
func (h *UserHandler) SignIn(c *gin.Context) {
	var input *request.SignIn
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	tokens, customErr := h.userUseCase.SignIn(request.MapSignInToUser(input))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "user signed in", tokens)
}

// SignOut godoc
// @Summary Sign out a user
// @Description Sign out a user with the provided refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body request.RefreshToken true "Refresh token object"
// @Success 200 {string} string "User signed out"
// @Failure 400 {string} string "Invalid input JSON"
// @Failure 401 {string} string "Unauthorized"
// @Router /api/auth/sign-out [post]
func (h *UserHandler) SignOut(c *gin.Context) {
	var input *request.RefreshToken
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	if customErr := h.userUseCase.SignOut(input.RefreshToken); customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "user signed out", nil)
}

// RefreshToken godoc
// @Summary Refresh access and refresh tokens
// @Description Refresh access and refresh tokens with the provided refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body request.RefreshToken true "Refresh token object"
// @Success 200 {object} models.Token "Successful response"
// @Failure 400 {string} string "Invalid input JSON"
// @Failure 401 {string} string "Unauthorized"
// @Router /api/auth/refresh-token [post]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var input *request.RefreshToken
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	accessToken, customErr := h.userUseCase.RefreshTokens(input.RefreshToken)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "token refreshed", accessToken)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON input
// @Tags users
// @Accept json
// @Produce json
// @Param input body request.CreateUser true "User object to be created"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input *request.CreateUser
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, err := h.userUseCase.CreateUser(request.MapCreateUserToUser(input))
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "user created", gin.H{
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
// @Success 200 {array} response.GetUser "Successful response"
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /api/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users *[]models.User
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}
	data := make([]*response.GetUser, len(*users))
	for i, user := range *users {
		data[i] = response.MapUserToGetUser(&user)
	}
	NewSuccessResponse(c, http.StatusOK, "all users retrieved", data)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user based on ID
// @ID get-user-by-id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} response.GetUser "Successful response"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	user, customErr := h.userUseCase.GetUserByID(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}
	data := response.MapUserToGetUser(user)
	NewSuccessResponse(c, http.StatusOK, "user retrieved", data)
}

// UpdateUser godoc
// @Summary Update the existing user
// @Description Update the existing user with the provided JSON input
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Param input body request.UpdateUser true "User object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input *request.UpdateUser
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, gin.H{"id": id})
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	customErr := h.userUseCase.UpdateUser(uint(id), request.MapUpdateUserToUser(input))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}

	NewSuccessResponse(c, http.StatusOK, "user updated", nil)
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
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	customErr := h.userUseCase.DeleteUser(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}

	NewSuccessResponse(c, http.StatusOK, "user deleted", nil)
}
