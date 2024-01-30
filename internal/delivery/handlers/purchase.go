package handlers

import (
	"Canteen-Backend/internal/delivery/dto/request"
	"Canteen-Backend/internal/delivery/dto/response"
	"Canteen-Backend/internal/usecase"
	"Canteen-Backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initPurchaseRoutes(api *gin.RouterGroup) {

	api.Use(h.authenticateUser)
	{
		suppliers := api.Group("/suppliers")
		{
			suppliers.POST("/", h.purchaseHandler.CreateSupplier)
			suppliers.GET("/", h.purchaseHandler.GetAllSuppliers)
			suppliers.GET("/:id", h.purchaseHandler.GetSupplierByID)
			suppliers.PUT("/:id", h.purchaseHandler.UpdateSupplier)
			suppliers.DELETE("/:id", h.purchaseHandler.DeleteSupplier)
		}

		purchases := api.Group("/purchases")
		{
			purchases.POST("/", h.purchaseHandler.CreatePurchase)
			//purchases.GET("/", h.purchaseHandler.GetAllPurchases)
			//purchases.GET("/:id", h.purchaseHandler.GetPurchaseByID)
			//purchases.PUT("/:id", h.purchaseHandler.UpdatePurchase)
			//purchases.DELETE("/:id", h.purchaseHandler.DeletePurchase)
		}
	}
}

type PurchaseHandler struct {
	purchaseUseCase usecase.Purchase
}

func NewPurchaseHandler(purchaseUseCase usecase.Purchase) *PurchaseHandler {
	return &PurchaseHandler{purchaseUseCase: purchaseUseCase}
}

// CreateSupplier godoc
// @Summary Create a supplier
// @Description Create a supplier with the provided JSON input
// @Tags suppliers
// @Accept json
// @Produce json
// @Param input body request.CreateSupplier true "Supplier creation object"
// @Success 201 {object} request.CreateSupplier "Successful response"
// @Failure 400 {string} string "Invalid input JSON"
// @Failure 500 {string} string "Internal server error"
// @Router /api/suppliers [post]
func (h *PurchaseHandler) CreateSupplier(c *gin.Context) {
	var input *request.CreateSupplier
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "validation error", err, nil)
		return
	}

	supplier := request.MapCreateSupplierToSupplier(input)
	id, customErr := h.purchaseUseCase.CreateSupplier(supplier)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusCreated, "supplier created successfully", gin.H{"id": id})
}

// GetAllSuppliers godoc
// @Summary Get all suppliers
// @Description Get all suppliers
// @Tags suppliers
// @Produce json
// @Success 200 {array} response.GetSupplier "Successful response"
// @Failure 500 {string} string "Internal server error"
// @Router /api/suppliers [get]
func (h *PurchaseHandler) GetAllSuppliers(c *gin.Context) {
	suppliers, customErr := h.purchaseUseCase.GetAllSuppliers()
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	data := make([]*response.GetSupplier, len(*suppliers))
	for i, supplier := range *suppliers {
		data[i] = response.MapSupplierToGetSupplier(&supplier)
	}

	NewSuccessResponse(c, http.StatusOK, "suppliers retrieved successfully", data)
}

// GetSupplierByID godoc
// @Summary Get a supplier by id
// @Description Get a supplier by id
// @Tags suppliers
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} response.GetSupplier "Successful response"
// @Failure 400 {string} string "Invalid supplier id"
// @Failure 404 {string} string "Supplier not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/suppliers/{id} [get]
func (h *PurchaseHandler) GetSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid supplier id", err, nil)
		return
	}

	supplier, customErr := h.purchaseUseCase.GetSupplierByID(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "supplier retrieved successfully", response.MapSupplierToGetSupplier(supplier))
}

// UpdateSupplier godoc
// @Summary Update a supplier
// @Description Update a supplier with the provided JSON input
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Param input body request.UpdateSupplier true "Supplier update object"
// @Success 200 {string} string "Successful response"
// @Failure 400 {string} string "Invalid input JSON"
// @Failure 404 {string} string "Supplier not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/suppliers/{id} [put]
func (h *PurchaseHandler) UpdateSupplier(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid supplier id", err, nil)
		return
	}

	var input *request.UpdateSupplier
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "validation error", err, nil)
		return
	}

	supplier := request.MapUpdateSupplierToSupplier(input)
	supplier.ID = uint(id)

	customErr := h.purchaseUseCase.UpdateSupplier(supplier)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "supplier updated successfully", nil)
}

// DeleteSupplier godoc
// @Summary Delete a supplier
// @Description Delete a supplier
// @Tags suppliers
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Successful response"
// @Failure 400 {string} string "Invalid supplier id"
// @Failure 404 {string} string "Supplier not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/suppliers/{id} [delete]
func (h *PurchaseHandler) DeleteSupplier(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid supplier id", err, nil)
		return
	}

	customErr := h.purchaseUseCase.DeleteSupplier(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "supplier deleted successfully", nil)
}

func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var input *request.CreatePurchase
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), err, nil)
		return
	}

	purchase := request.MapCreatePurchaseToPurchase(input)
	id, customErr := h.purchaseUseCase.CreatePurchase(purchase)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusCreated, "purchase created successfully", gin.H{"id": id})
}
