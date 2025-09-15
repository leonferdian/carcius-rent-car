package main

import (
	"log"
	"os"

	"github.com/carcius-rent-car/users-service/database"
	"github.com/carcius-rent-car/users-service/handlers"
	"github.com/carcius-rent-car/users-service/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	database.InitDB()

	// Set up router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Create auth handler
	authHandler := handlers.NewAuthHandler(database.DB)

	// API routes
	api := r.Group("/api")
	{
		// Public routes
		auth := api.Group("/users")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := api.Group("/users")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/:id", authHandler.GetProfile)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port
	}

	r.Run(":" + port)
}
