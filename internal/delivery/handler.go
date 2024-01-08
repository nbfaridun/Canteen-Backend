package delivery

import (
	"Canteen-Backend/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	useCase *usecase.UseCase
}

func NewHandler(useCase *usecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.POST("/", h.CreateUser)
		users.GET("/", h.GetAllUsers)
		users.GET("/:id", h.GetUserByID)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}

	clients := router.Group("/clients")
	{
		clients.POST("/", h.CreateClient)
		clients.GET("/", h.GetAllClients)
		clients.GET("/:id", h.GetClientByID)
		clients.PUT("/:id", h.UpdateClient)
		clients.DELETE("/:id", h.DeleteClient)
	}

	return router
}
