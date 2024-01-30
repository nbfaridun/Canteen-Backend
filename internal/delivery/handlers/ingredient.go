package handlers

import (
	"Canteen-Backend/internal/delivery/dto/request"
	"Canteen-Backend/internal/delivery/dto/response"
	"Canteen-Backend/internal/usecase"
	"Canteen-Backend/pkg/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initIngredientRoutes(api *gin.RouterGroup) {

	api.Use(h.authenticateUser)
	{
		ingredientCategories := api.Group("/ingredient-categories")
		{
			ingredientCategories.POST("/", h.ingredientHandler.CreateIngredientCategory)
			ingredientCategories.GET("/", h.ingredientHandler.GetAllIngredientCategories)
			ingredientCategories.GET("/:id", h.ingredientHandler.GetIngredientCategoryByID)
			ingredientCategories.PUT("/:id", h.ingredientHandler.UpdateIngredientCategory)
			ingredientCategories.DELETE("/:id", h.ingredientHandler.DeleteIngredientCategory)
		}

		ingredients := api.Group("/ingredients")
		{
			ingredients.POST("/", h.ingredientHandler.CreateIngredient)
			ingredients.GET("/", h.ingredientHandler.GetAllIngredients)
			ingredients.GET("/:id", h.ingredientHandler.GetIngredientByID)
			ingredients.PUT("/:id", h.ingredientHandler.UpdateIngredient)
			ingredients.DELETE("/:id", h.ingredientHandler.DeleteIngredient)
		}
	}

}

type IngredientHandler struct {
	ingredientUseCase usecase.Ingredient
}

func NewIngredientHandler(ingredientUseCase usecase.Ingredient) *IngredientHandler {
	return &IngredientHandler{ingredientUseCase: ingredientUseCase}
}

// CreateIngredientCategory godoc
// @Summary Create a new ingredient category
// @Description This endpoint allows you to create a new ingredient category.
// @ID create-ingredient-category
// @Tags ingredient_categories
// @Accept  json
// @Produce  json
// @Param input body request.CreateIngredientCategory true "Create PurchasedIngredient Category"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/ingredient-categories [post]
func (h *IngredientHandler) CreateIngredientCategory(c *gin.Context) {
	var input *request.CreateIngredientCategory
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, customErr := h.ingredientUseCase.CreateIngredientCategory(request.MapCreateIngredientCategoryToIngredientCategory(input))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "client created", gin.H{
		"id": id,
	})
}

// GetAllIngredientCategories godoc
// @Summary Get all ingredient categories
// @Description Get all ingredient categories available
// @ID get-all-ingredient-categories
// @Tags ingredient_categories
// @Accept json
// @Produce json
// @Success 200 {array} response.GetIngredientCategory "Successful response"
// @Failure 500 {string} string
// @Router /api/ingredient-categories [get]
func (h *IngredientHandler) GetAllIngredientCategories(c *gin.Context) {
	ingredientCategories, err := h.ingredientUseCase.GetAllIngredientCategories()
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}

	data := make([]*response.GetIngredientCategory, len(*ingredientCategories))
	for i, ingredientCategory := range *ingredientCategories {
		data[i] = response.MapIngredientCategoryToGetIngredientCategory(&ingredientCategory)
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient categories retrieved", data)
}

// GetIngredientCategoryByID godoc
// @Summary Get an ingredient category by ID
// @Description Get an ingredient category based on ID
// @ID get-ingredient-category-by-id
// @Tags ingredient_categories
// @Accept json
// @Produce json
// @Param id path int true "PurchasedIngredient category ID" Format(int64)
// @Success 200 {object} response.GetIngredientCategory "Successful response"
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /api/ingredient-categories/{id} [get]
func (h *IngredientHandler) GetIngredientCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	ingredientCategory, customErr := h.ingredientUseCase.GetIngredientCategoryByID(uint(id))
	if customErr != nil {
		fmt.Println(1)
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient category retrieved", response.MapIngredientCategoryToGetIngredientCategory(ingredientCategory))
}

// UpdateIngredientCategory godoc
// @Summary Update the existing ingredient category
// @Description Update the existing ingredient category with the provided JSON input
// @Tags ingredient_categories
// @Accept json
// @Produce json
// @Param id path int true "PurchasedIngredient category ID" Format(int64)
// @Param input body request.UpdateIngredientCategory true "PurchasedIngredient category object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /api/ingredient-categories/{id} [put]
func (h *IngredientHandler) UpdateIngredientCategory(c *gin.Context) {
	var input *request.UpdateIngredientCategory
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	ingredientCategory := request.MapUpdateIngredientCategoryToIngredientCategory(input)
	ingredientCategory.ID = uint(id)

	customErr := h.ingredientUseCase.UpdateIngredientCategory(ingredientCategory)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient category updated", nil)
}

// DeleteIngredientCategory godoc
// @Summary Delete the existing ingredient category
// @Description Delete the existing ingredient category with the provided ID
// @Tags ingredient_categories
// @Accept json
// @Produce json
// @Param id path int true "PurchasedIngredient category ID" Format(int64)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /api/ingredient-categories/{id} [delete]
func (h *IngredientHandler) DeleteIngredientCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	customErr := h.ingredientUseCase.DeleteIngredientCategory(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient category deleted", nil)
}

// CreateIngredient godoc
// @Summary Create a new ingredient
// @Description This endpoint allows you to create a new ingredient. Quantity, UnitPrice, LackLimit, PurchaseDate, ExpirationDate are optional.
// @Tags ingredients
// @Accept  json
// @Produce  json
// @Param input body request.CreateIngredient true "Create PurchasedIngredient"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/ingredients [post]
func (h *IngredientHandler) CreateIngredient(c *gin.Context) {
	var input *request.CreateIngredient
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, customErr := h.ingredientUseCase.CreateIngredient(request.MapCreateIngredientToIngredient(input))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient created", gin.H{
		"id": id,
	})
}

// GetAllIngredients godoc
// @Summary Get all ingredients
// @Description Get all ingredients available
// @Tags ingredients
// @Accept json
// @Produce json
// @Success 200 {array} response.GetIngredient "Successful response"
// @Failure 500 {string} string
// @Router /api/ingredients [get]
func (h *IngredientHandler) GetAllIngredients(c *gin.Context) {
	ingredients, err := h.ingredientUseCase.GetAllIngredients()
	if err != nil {
		NewErrorResponse(c, err.StatusCode, err.Message, err.Error, nil)
		return
	}

	data := make([]*response.GetIngredient, len(*ingredients))
	for i, ingredient := range *ingredients {
		data[i] = response.MapIngredientToGetIngredient(&ingredient)
	}

	NewSuccessResponse(c, http.StatusOK, "ingredients retrieved", data)
}

// GetIngredientByID godoc
// @Summary Get an ingredient by ID
// @Description Get an ingredient based on ID
// @Tags ingredients
// @Accept json
// @Produce json
// @Param id path int true "PurchasedIngredient ID" Format(int64)
// @Success 200 {object} response.GetIngredient "Successful response"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/ingredients/{id} [get]
func (h *IngredientHandler) GetIngredientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	ingredient, customErr := h.ingredientUseCase.GetIngredientByID(uint(id))
	if customErr != nil {
		fmt.Println(1)
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient retrieved", response.MapIngredientToGetIngredient(ingredient))
}

// UpdateIngredient godoc
// @Summary Update the existing ingredient
// @Description Update the existing ingredient with the provided JSON input
// @Tags ingredients
// @Accept json
// @Produce json
// @Param id path int true "PurchasedIngredient ID" Format(int64)
// @Param input body request.UpdateIngredient true "PurchasedIngredient object to be updated"
// @Success 200 {string} string
// @Failure 400 {string} string "Invalid input JSON"
// @Failure 404 {string} string "PurchasedIngredient not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/ingredients/{id} [put]
func (h *IngredientHandler) UpdateIngredient(c *gin.Context) {
	var input *request.UpdateIngredient
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input JSON", err, nil)
		return
	}

	if err := validator.ValidatePayload(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	ingredient := request.MapUpdateIngredientToIngredient(input)
	ingredient.ID = uint(id)

	customErr := h.ingredientUseCase.UpdateIngredient(ingredient)
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient updated", nil)
}

// DeleteIngredient godoc
// @Summary Delete the existing ingredient
// @Description Delete the existing ingredient with the provided ID
// @Tags ingredients
// @Accept json
// @Produce json
// @Param id path int true "PurchasedIngredient ID" Format(int64)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/ingredients/{id} [delete]
func (h *IngredientHandler) DeleteIngredient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id", err, nil)
		return
	}

	customErr := h.ingredientUseCase.DeleteIngredient(uint(id))
	if customErr != nil {
		NewErrorResponse(c, customErr.StatusCode, customErr.Message, customErr.Error, nil)
		return
	}

	NewSuccessResponse(c, http.StatusOK, "ingredient deleted", nil)
}
