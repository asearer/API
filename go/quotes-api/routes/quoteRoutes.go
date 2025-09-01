// routes/quoteRoutes.go
// Defines the API routes for managing quotes.

package routes

import (
	"quotes-api/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterQuoteRoutes sets up all routes for quote management
func RegisterQuoteRoutes(router *gin.Engine) {
	quoteRoutes := router.Group("/quotes")
	{
		quoteRoutes.GET("/", controllers.GetQuotes)            // Get all quotes
		quoteRoutes.GET("/random", controllers.GetRandomQuote) // Get random quote
		quoteRoutes.GET("/:id", controllers.GetQuoteByID)      // Get quote by ID
		quoteRoutes.POST("/", controllers.CreateQuote)         // Add new quote
		quoteRoutes.PUT("/:id", controllers.UpdateQuote)       // Update existing quote
		quoteRoutes.DELETE("/:id", controllers.DeleteQuote)    // Delete quote
	}
}
