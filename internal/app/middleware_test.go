package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// getMockedContext func is just a helper to mock gin.Context
func getMockedContext(request *http.Request) *gin.Context {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}

func TestAuthMiddleware(t *testing.T) {

	t.Run("with correct API key", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/", nil)
		request.Header.Set("x-api-key", "kawabanga")
		context := getMockedContext(request)

		AuthMiddleware(context)

		assert.Equal(t, http.StatusOK, context.Writer.Status())
	})

	t.Run("with incorrect API key", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/", nil)
		request.Header.Set("x-api-key", "wrongKey")
		context := getMockedContext(request)

		AuthMiddleware(context)

		assert.Equal(t, http.StatusUnauthorized, context.Writer.Status())
	})

	t.Run("without API key", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/", nil)
		context := getMockedContext(request)

		AuthMiddleware(context)

		assert.Equal(t, http.StatusUnauthorized, context.Writer.Status())
	})
}
