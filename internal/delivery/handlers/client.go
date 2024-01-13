package handlers

import (
	"Canteen-Backend/internal/delivery/dto/request"
	"Canteen-Backend/internal/delivery/dto/response"
	"Canteen-Backend/internal/usecase"
	"Canteen-Backend/pkg/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initClientRoutes(api *gin.RouterGroup) {

	clients := api.Group("/clients")
	{
		clients.POST("/", h.clientHandler.CreateClient)
		clients.GET("/", h.clientHandler.GetAllClients)
		clients.GET("/:id", h.clientHandler.GetClientByID)
		clients.PUT("/:id", h.clientHandler.UpdateClient)
		clients.DELETE("/:id", h.clientHandler.DeleteClient)

		clients.PUT("/:id/modify-balance", h.clientHandler.ModifyBalanceByClientID)
	}
}

type ClientHandler struct {
	clientUseCase usecase.Client
}

func NewClientHandler(clientUseCase usecase.Client) *ClientHandler {
	return &ClientHandler{clientUseCase: clientUseCase}
}

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client with the provided JSON input
// @Tags clients
// @Accept json
// @Produce json
// @Param input body request.CreateClient true "Client object to be created"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/clients [post]
func (h *ClientHandler) CreateClient(c *gin.Context) {
	var input *request.CreateClient
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, err := h.clientUseCase.CreateClient(request.MapCreateClientToClient(input))
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client created", gin.H{
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
// @Success 200 {array} response.GetClient "Successful response"
// @Failure 500 {string} string
// @Router /api/clients [get]
func (h *ClientHandler) GetAllClients(c *gin.Context) {
	clients, err := h.clientUseCase.GetAllClients()
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}
	data := make([]*response.GetClient, len(*clients))
	for i, client := range *clients {
		data[i] = response.MapClientToGetClient(&client)
	}
	NewSuccessResponse(c, http.StatusOK, "all clients received", data)
}

// GetClientByID godoc
// @Summary Get a client by ID
// @Description Get a client based on ID
// @ID get-client-by-id
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID" Format(int64)
// @Success 200 {object} response.GetClient "Successful response"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/clients/{id} [get]
func (h *ClientHandler) GetClientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	client, customErr := h.clientUseCase.GetClientByID(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}

	data := response.MapClientToGetClient(client)
	NewSuccessResponse(c, http.StatusOK, "client received", data)
}

// UpdateClient godoc
// @Summary Update the existing client
// @Description Update the existing client with the provided JSON input
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID" Format(int64)
// @Param input body request.UpdateClient true "Client object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/clients/{id} [put]
func (h *ClientHandler) UpdateClient(c *gin.Context) {
	var input *request.UpdateClient
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid client JSON", err, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, gin.H{"id": id})
		return
	}

	customErr := h.clientUseCase.UpdateClient(uint(id), request.MapUpdateClientToClient(input))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client updated", nil)
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
// @Router /api/clients/{id} [delete]
func (h *ClientHandler) DeleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	customErr := h.clientUseCase.DeleteClient(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client deleted", nil)
}

// ModifyBalanceByClientID godoc
// @Summary Modify the balance of a client by ID
// @Description Modify the balance of a client based on ID and provided JSON input
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID" Format(int64)
// @Param input body request.ModifyBalance true "Balance modification object"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/clients/{id}/modify-balance [put]
func (h *ClientHandler) ModifyBalanceByClientID(c *gin.Context) {
	var input *request.ModifyBalance
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

	customErr := h.clientUseCase.ModifyBalanceByClientID(uint(id), input.Difference)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, gin.H{"id": id})
		return
	}

	NewSuccessResponse(c, http.StatusOK, "money withdrawn from client", nil)
}
