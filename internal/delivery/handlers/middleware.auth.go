package handlers

import (
	"Canteen-Backend/pkg/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) authenticateUser(c *gin.Context) {
	userId, userRoleId, err := parseAuthHeader(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error(), err, nil)
		c.Abort()
		return
	}

	c.Set("user_id", userId)
	c.Set("user_role_id", userRoleId)
}

func parseAuthHeader(c *gin.Context) (uint, uint, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return 0, 0, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, 0, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return 0, 0, errors.New("token is empty")
	}

	return auth.ParseToken(headerParts[1])
}
