package main

import (
	"log"
	"os"

	"github.com/carcius-rent-car/cars-service/database"
	"github.com/carcius-rent-car/cars-service/handlers"
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

	// Create car handler
	carHandler := handlers.NewCarHandler(database.DB)

	// API routes
	api := r.Group("/api")
	{
		// Car routes
		cars := api.Group("/cars")
		{
			cars.GET("", carHandler.GetCars)
			cars.GET("/:id", carHandler.GetCar)
			cars.POST("", carHandler.CreateCar)
			cars.PUT("/:id/status", carHandler.UpdateCarStatus)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // Default port
	}

	r.Run(":" + port)
}
