package main

import (
	"log"
	"os"

	"github.com/carcius-rent-car/bookings-service/database"
	"github.com/carcius-rent-car/bookings-service/handlers"
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

	// Create booking handler
	bookingHandler := handlers.NewBookingHandler(database.DB)

	// API routes
	api := r.Group("/api")
	{
		// Booking routes
		bookings := api.Group("/bookings")
		{
			bookings.POST("", bookingHandler.CreateBooking)
			bookings.GET("/user/:userId", bookingHandler.GetUserBookings)
			bookings.PUT("/:id/status", bookingHandler.UpdateBookingStatus)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083" // Default port
	}

	r.Run(":" + port)
}
