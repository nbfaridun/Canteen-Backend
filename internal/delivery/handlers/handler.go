package handlers

import (
	"Canteen-Backend/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	userHandler       *UserHandler
	clientHandler     *ClientHandler
	ingredientHandler *IngredientHandler
	purchaseHandler   *PurchaseHandler
}

func NewHandler(useCase *usecase.UseCase) *Handler {
	userHandler := NewUserHandler(useCase.User)
	clientHandler := NewClientHandler(useCase.Client)
	ingredientHandler := NewIngredientHandler(useCase.Ingredient)
	purchaseHandler := NewPurchaseHandler(useCase.Purchase)

	return &Handler{userHandler: userHandler, clientHandler: clientHandler, ingredientHandler: ingredientHandler, purchaseHandler: purchaseHandler}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		h.initUserRoutes(api)
		h.initClientRoutes(api)
		h.initIngredientRoutes(api)
		h.initPurchaseRoutes(api)
	}

	return router
}
