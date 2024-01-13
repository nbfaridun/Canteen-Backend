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

func (h *Handler) initClientCategoryRoutes(api *gin.RouterGroup) {
	clientCategories := api.Group("/client-categories")
	{
		clientCategories.POST("/", h.clientCategoryHandler.CreateClientCategory)
		clientCategories.GET("/", h.clientCategoryHandler.GetAllClientCategories)
		clientCategories.GET("/:id", h.clientCategoryHandler.GetClientCategoryByID)
		clientCategories.PUT("/:id", h.clientCategoryHandler.UpdateClientCategory)
		clientCategories.DELETE("/:id", h.clientCategoryHandler.DeleteClientCategory)
	}
}

type ClientCategoryHandler struct {
	clientCategoryUseCase usecase.ClientCategory
}

func NewClientCategoryHandler(clientCategoryUseCase usecase.ClientCategory) *ClientCategoryHandler {
	return &ClientCategoryHandler{clientCategoryUseCase: clientCategoryUseCase}
}

// CreateClientCategory godoc
// @Summary Create a new client category
// @Description Create a new client category with the provided JSON input
// @Tags client_categories
// @Accept json
// @Produce json
// @Param input body request.CreateClientCategory true "Client category object to be created"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/client-categories [post]
func (h *ClientCategoryHandler) CreateClientCategory(c *gin.Context) {
	var input *request.CreateClientCategory
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, err := h.clientCategoryUseCase.CreateClientCategory(request.MapCreateClientCategoryToClientCategory(input))
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client category created", gin.H{"id": id})

}

// GetAllClientCategories godoc
// @Summary Get all client categories
// @Description Get all client categories available
// @ID get-all-client-categories
// @Tags client_categories
// @Accept json
// @Produce json
// @Success 200 {array} response.GetClientCategory "Successful response"
// @Failure 500 {string} string
// @Router /api/client-categories [get]
func (h *ClientCategoryHandler) GetAllClientCategories(c *gin.Context) {
	clientCategories, err := h.clientCategoryUseCase.GetAllClientCategories()
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}

	data := make([]*response.GetClientCategory, len(*clientCategories))
	for i, clientCategory := range *clientCategories {
		data[i] = response.MapClientCategoryToGetClientCategory(&clientCategory)
	}
	NewSuccessResponse(c, http.StatusOK, "client categories retrieved", data)
}

// GetClientCategoryByID godoc
// @Summary Get a client category by ID
// @Description Get a client category based on ID
// @ID get-client-category-by-id
// @Tags client_categories
// @Accept json
// @Produce json
// @Param id path int true "Client category ID" Format(int64)
// @Success 200 {object} response.GetClientCategory "Successful response"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/client-categories/{id} [get]
func (h *ClientCategoryHandler) GetClientCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	clientCategory, customErr := h.clientCategoryUseCase.GetClientCategoryByID(uint(id))
	if err != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client category retrieved", response.MapClientCategoryToGetClientCategory(clientCategory))
}

// UpdateClientCategory godoc
// @Summary Update the existing client category
// @Description Update the existing client category with the provided JSON input
// @Tags client_categories
// @Accept json
// @Produce json
// @Param id path int true "Client category ID" Format(int64)
// @Param input body request.UpdateClientCategory true "Client category object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/client-categories/{id} [put]
func (h *ClientCategoryHandler) UpdateClientCategory(c *gin.Context) {
	var input *request.UpdateClientCategory
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validators.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	customErr := h.clientCategoryUseCase.UpdateClientCategory(uint(id), request.MapUpdateClientCategoryToClientCategory(input))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client category updated", nil)
}

// DeleteClientCategory godoc
// @Summary Delete a client category by ID
// @Description Delete a client category based on ID
// @ID delete-client-category
// @Tags client_categories
// @Accept json
// @Produce json
// @Param id path int true "Client category ID" Format(int64)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/client-categories/{id} [delete]
func (h *ClientCategoryHandler) DeleteClientCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	customErr := h.clientCategoryUseCase.DeleteClientCategory(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client category deleted", nil)
}
