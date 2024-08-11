package app

import (
	"github.com/cat9host/gin-air-boilerplate/internal/config"
	"github.com/cat9host/gin-air-boilerplate/internal/utils/formatter/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, response.UnauthorizedRequestError(message))
}

func AuthMiddleware(c *gin.Context) {
	requiredToken := config.AppSecretKey
	token := c.Request.Header.Get("x-api-key")

	if requiredToken == "" {
		requiredToken = "kawabanga"
	}

	if token != requiredToken {
		e := "Invalid API key"
		if token == "" {
			e = "API key required"
		}
		respondWithError(c, http.StatusUnauthorized, e)
		c.Next()
	}

	c.Next()
}
