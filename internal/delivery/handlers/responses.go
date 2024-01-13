package handlers

import (
	"Canteen-Backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	if data != nil {
		logger.GetLogger().Info(message, zap.Any("data", data))
	} else {
		logger.GetLogger().Info(message)
	}

	c.JSON(statusCode, data)

}

func NewErrorResponse(c *gin.Context, statusCode int, message string, err error, data interface{}) {
	if data != nil {
		logger.GetLogger().Error(message, zap.Error(err), zap.Any("data", data))
	} else {
		logger.GetLogger().Error(message, zap.Error(err))
	}

	c.JSON(statusCode, gin.H{
		"error": message,
	})
}
