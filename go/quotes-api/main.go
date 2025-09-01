// main.go
// Entry point for the Quotes API using Gin.
// It sets up the router and starts the HTTP server.

package main

import (
	"quotes-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterQuoteRoutes(router)

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Quotes API 🚀"})
	})

	// Start server on port 8080
	router.Run(":8080")
}
