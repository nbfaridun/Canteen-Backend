package handlers

import (
	"Canteen-Backend/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	userHandler           *UserHandler
	clientHandler         *ClientHandler
	clientCategoryHandler *ClientCategoryHandler
}

func NewHandler(useCase *usecase.UseCase) *Handler {
	userHandler := NewUserHandler(useCase.User)
	clientHandler := NewClientHandler(useCase.Client)
	clientCategoryHandler := NewClientCategoryHandler(useCase.ClientCategory)
	return &Handler{userHandler: userHandler, clientHandler: clientHandler, clientCategoryHandler: clientCategoryHandler}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		h.initUserRoutes(api)
		h.initClientRoutes(api)
	}

	return router
}
