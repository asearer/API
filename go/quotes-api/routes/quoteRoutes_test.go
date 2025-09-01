package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestRouter sets up Gin router with all quote routes
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	RegisterQuoteRoutes(r)
	return r
}

// TestRoutesRegistered ensures the main routes exist
func TestRoutesRegistered(t *testing.T) {
	router := setupTestRouter()

	endpoints := []string{"/quotes", "/quotes/random", "/quotes/1"}
	for _, ep := range endpoints {
		req, _ := http.NewRequest("GET", ep, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusNotFound, w.Code, ep+" should be registered")
	}
}
