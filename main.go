package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Version of the application
const Version = "1.0.0"

func setupRouter() *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Health endpoint
	// Shows the status, timestamp, and version number
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
			"version":   Version,
		})
	})

	// Route: Messages endpoint
	r.GET("/messages", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "Hello from messages route!",
			"description": "This is a simple dummy endpoint for testing purposes.",
			"success":     true,
		})
	})

	// Route: Users endpoint
	r.GET("/users", func(c *gin.Context) {
		users := []gin.H{
			{"id": 1, "name": "Alice Smith", "email": "alice@example.com"},
			{"id": 2, "name": "Bob Jones", "email": "bob@example.com"},
			{"id": 3, "name": "Charlie Brown", "email": "charlie@example.com"},
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    users,
			"total":   len(users),
			"success": true,
		})
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": string("dev:" + Version)})
	})

	return r
}

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it, using system environment variables")
	}

	// Set Gin mode based on environment or default to release/debug
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	// Setup host and port
	hostIP := os.Getenv("HOST_IP")
	if hostIP == "" {
		hostIP = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()

	address := hostIP + ":" + port
	log.Printf("Server starting on %s", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
