package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	routerMain, _, _ := SetupRouter(false)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	routerMain.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
