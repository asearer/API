package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"quotes-api/models"
)

// setupRouter creates a Gin router with all routes for testing.
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/quotes", GetQuotes)
	router.GET("/quotes/random", GetRandomQuote)
	router.GET("/quotes/:id", GetQuoteByID)
	router.POST("/quotes", CreateQuote)
	router.PUT("/quotes/:id", UpdateQuote)
	router.DELETE("/quotes/:id", DeleteQuote)
	return router
}

// ------------------------------
// HAPPY PATH TESTS
// ------------------------------

func TestGetQuotes(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var quotes []models.Quote
	err := json.Unmarshal(w.Body.Bytes(), &quotes)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(quotes), 1)
}

func TestGetQuoteByID_ValidID(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/quotes/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var quote models.Quote
	err := json.Unmarshal(w.Body.Bytes(), &quote)
	assert.NoError(t, err)
	assert.Equal(t, 1, quote.ID)
}

func TestCreateQuote_Valid(t *testing.T) {
	router := setupRouter()
	body := []byte(`{"text":"Test Quote","author":"Tester"}`)
	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var quote models.Quote
	err := json.Unmarshal(w.Body.Bytes(), &quote)
	assert.NoError(t, err)
	assert.NotZero(t, quote.ID)
}

func TestUpdateQuote_Valid(t *testing.T) {
	router := setupRouter()
	body := []byte(`{"text":"Updated Quote","author":"Tester"}`)
	req, _ := http.NewRequest("PUT", "/quotes/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var quote models.Quote
	err := json.Unmarshal(w.Body.Bytes(), &quote)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Quote", quote.Text)
}

func TestDeleteQuote_Valid(t *testing.T) {
	backup := make([]models.Quote, len(models.QuoteStore))
	copy(backup, models.QuoteStore)
	defer func() { models.QuoteStore = backup }()

	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/quotes/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Quote deleted", resp["message"])
}

// ------------------------------
// EDGE CASE / INVALID INPUT TESTS
// ------------------------------

func TestGetQuoteByID_InvalidID(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/quotes/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Contains(t, resp["error"], "Invalid ID")
}

func TestUpdateQuote_InvalidID(t *testing.T) {
	router := setupRouter()
	body := []byte(`{"text":"Updated","author":"Tester"}`)
	req, _ := http.NewRequest("PUT", "/quotes/abc", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteQuote_InvalidID(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/quotes/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetRandomQuote_Empty(t *testing.T) {
	backup := make([]models.Quote, len(models.QuoteStore))
	copy(backup, models.QuoteStore)
	models.QuoteStore = []models.Quote{}
	defer func() { models.QuoteStore = backup }()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/quotes/random", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Contains(t, resp["error"], "No quotes")
}

func TestCreateQuote_BadJSON(t *testing.T) {
	router := setupRouter()
	body := []byte(`{invalid json}`)
	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateQuote_BadJSON(t *testing.T) {
	router := setupRouter()
	body := []byte(`{invalid json}`)
	req, _ := http.NewRequest("PUT", "/quotes/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ------------------------------
// ADDITIONAL TESTS TO COVER REMAINING BRANCHES
// ------------------------------

func TestUpdateQuote_NonExistentID(t *testing.T) {
	backup := make([]models.Quote, len(models.QuoteStore))
	copy(backup, models.QuoteStore)
	defer func() { models.QuoteStore = backup }()

	router := setupRouter()
	body := []byte(`{"text":"Nonexistent","author":"Nobody"}`)
	req, _ := http.NewRequest("PUT", "/quotes/9999", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Quote not found", resp["error"])
}

func TestDeleteQuote_NonExistentID(t *testing.T) {
	backup := make([]models.Quote, len(models.QuoteStore))
	copy(backup, models.QuoteStore)
	defer func() { models.QuoteStore = backup }()

	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/quotes/9999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Quote not found", resp["error"])
}

func TestCreateQuote_EmptyBody(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var newQuote models.Quote
	json.Unmarshal(w.Body.Bytes(), &newQuote)
	assert.NotZero(t, newQuote.ID)
}

// ------------------------------
// MOCKED RANDOM QUOTE TEST
// ------------------------------

func TestGetRandomQuote_Mocked(t *testing.T) {
	backup := make([]models.Quote, len(models.QuoteStore))
	copy(backup, models.QuoteStore)
	defer func() { models.QuoteStore = backup }()

	// Seed with multiple quotes
	models.QuoteStore = []models.Quote{
		{ID: 1, Text: "Quote 1", Author: "Author A"},
		{ID: 2, Text: "Quote 2", Author: "Author B"},
	}

	// Mock getRandomIndex to always return 1 (second quote)
	originalRandom := getRandomIndex
	getRandomIndex = func(n int) int { return 1 }
	defer func() { getRandomIndex = originalRandom }()

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/quotes/random", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var quote models.Quote
	err := json.Unmarshal(w.Body.Bytes(), &quote)
	assert.NoError(t, err)
	assert.Equal(t, 2, quote.ID)
}
