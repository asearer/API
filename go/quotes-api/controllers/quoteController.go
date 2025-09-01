/// controllers/quoteController.go
// Business logic for handling quotes: CRUD + random retrieval.

package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	"quotes-api/models"

	"github.com/gin-gonic/gin"
)

// getRandomIndex wraps rand.Intn so it can be mocked in tests
var getRandomIndex = func(n int) int {
	return rand.Intn(n)
}

// GetQuotes returns all quotes from the shared models.QuoteStore
func GetQuotes(c *gin.Context) {
	c.JSON(http.StatusOK, models.QuoteStore)
}

// GetRandomQuote returns a single random quote
func GetRandomQuote(c *gin.Context) {
	if len(models.QuoteStore) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No quotes available"})
		return
	}

	quote := models.QuoteStore[getRandomIndex(len(models.QuoteStore))]
	c.JSON(http.StatusOK, quote)
}

// GetQuoteByID finds a quote by its ID
func GetQuoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, q := range models.QuoteStore {
		if q.ID == id {
			c.JSON(http.StatusOK, q)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Quote not found"})
}

// CreateQuote adds a new quote
func CreateQuote(c *gin.Context) {
	var newQuote models.Quote

	if err := c.ShouldBindJSON(&newQuote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newQuote.ID = len(models.QuoteStore) + 1
	models.QuoteStore = append(models.QuoteStore, newQuote)

	c.JSON(http.StatusCreated, newQuote)
}

// UpdateQuote modifies an existing quote
func UpdateQuote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, q := range models.QuoteStore {
		if q.ID == id {
			if err := c.ShouldBindJSON(&models.QuoteStore[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			models.QuoteStore[i].ID = id
			c.JSON(http.StatusOK, models.QuoteStore[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Quote not found"})
}

// DeleteQuote removes a quote by ID
func DeleteQuote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, q := range models.QuoteStore {
		if q.ID == id {
			models.QuoteStore = append(models.QuoteStore[:i], models.QuoteStore[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Quote deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Quote not found"})
}
