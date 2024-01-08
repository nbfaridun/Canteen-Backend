package delivery

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client with the provided JSON input
// @Tags clients
// @Accept json
// @Produce json
// @Param input body models.CreateClientInput true "Client object to be created"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /clients [post]
func (h *Handler) CreateClient(c *gin.Context) {
	var input *models.CreateClientInput
	if err := c.BindJSON(&input); err != nil {
		logger.GetLogger().Error("error while binding json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input JSON"})
		return
	}

	id, err := h.useCase.Client.CreateClient(&models.Client{
		Email:            input.Email,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		ClientCategoryID: input.ClientCategoryID,
		Balance:          input.Balance,
		Age:              input.Age,
		Gender:           input.Gender,
		IsActive:         true,
	})
	if err != nil {
		logger.GetLogger().Error("error while creating client", zap.Error(err.LogError))
		c.JSON(err.StatusCode, gin.H{"error": err.FrontendMessage})
		return
	}

	logger.GetLogger().Info("client created", zap.Uint("id", id))
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// GetAllClients godoc
// @Summary Get all clients
// @Description Get all clients available
// @ID get-all-clients
// @Tags clients
// @Accept json
// @Produce json
// @Success 200 {array} models.Client "Successful response"
// @Failure 500 {string} string
// @Router /clients [get]
func (h *Handler) GetAllClients(c *gin.Context) {
	clients, err := h.useCase.Client.GetAllClients()
	if err != nil {
		logger.GetLogger().Error("error while getting all clients", zap.Error(err.LogError))
		c.JSON(err.StatusCode, gin.H{"error": err.FrontendMessage})
		return
	}

	logger.GetLogger().Info("all clients received")
	c.JSON(http.StatusOK, clients)
}

// GetClientByID godoc
// @Summary Get a client by ID
// @Description Get a client based on ID
// @ID get-client-by-id
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID" Format(int64)
// @Success 200 {object} models.Client "Successful response"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /clients/{id} [get]
func (h *Handler) GetClientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.GetLogger().Error("error while converting id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	client, customErr := h.useCase.Client.GetClientByID(uint(id))
	if customErr != nil {
		logger.GetLogger().Error("error while getting client by id", zap.Error(customErr.LogError))
		c.JSON(customErr.StatusCode, gin.H{"error": customErr.FrontendMessage})
		return
	}

	logger.GetLogger().Info("client received", zap.Uint("id", client.ID))
	c.JSON(http.StatusOK, client)
}

// UpdateClient godoc
// @Summary Update the existing client
// @Description Update the existing client with the provided JSON input
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID" Format(int64)
// @Param input body models.UpdateClientInput true "Client object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /clients/{id} [put]
func (h *Handler) UpdateClient(c *gin.Context) {
	var input *models.UpdateClientInput
	if err := c.BindJSON(&input); err != nil {
		logger.GetLogger().Error("error while binding json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client JSON"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.GetLogger().Error("error while converting id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customErr := h.useCase.Client.UpdateClient(uint(id), &models.Client{
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Age:              input.Age,
		Gender:           input.Gender,
		Email:            input.Email,
		ClientCategoryID: input.ClientCategoryID,
		Balance:          input.Balance,
		IsActive:         input.IsActive,
	})
	if customErr != nil {
		logger.GetLogger().Error("error while updating client", zap.Error(customErr.LogError))
		c.JSON(customErr.StatusCode, gin.H{"error": customErr.FrontendMessage})
		return
	}

	logger.GetLogger().Info("client updated", zap.Uint("id", uint(id)))
	c.JSON(http.StatusOK, gin.H{"message": "client updated"})
}

// DeleteClient godoc
// @Summary Delete a client by ID
// @Description Delete a client based on ID
// @ID delete-client
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID" Format(int64)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /clients/{id} [delete]
func (h *Handler) DeleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.GetLogger().Error("error while converting id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customErr := h.useCase.Client.DeleteClient(uint(id))
	if customErr != nil {
		logger.GetLogger().Error("error while deleting client", zap.Error(customErr.LogError))
		c.JSON(customErr.StatusCode, gin.H{"error": customErr.FrontendMessage})
		return
	}

	logger.GetLogger().Info("client deleted", zap.Uint("id", uint(id)))
	c.JSON(http.StatusOK, gin.H{"message": "client deleted"})
}
