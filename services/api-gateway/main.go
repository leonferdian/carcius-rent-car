package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	usersServiceURL    string
	carsServiceURL     string
	bookingsServiceURL string
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get service URLs from environment variables
	usersServiceURL = getEnv("USERS_SERVICE_URL", "http://users-service:8081")
	carsServiceURL = getEnv("CARS_SERVICE_URL", "http://cars-service:8082")
	bookingsServiceURL = getEnv("BOOKINGS_SERVICE_URL", "http://bookings-service:8083")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := r.Group("/api")
	{
		// Users service routes
		users := v1.Group("/users")
		{
			users.POST("/register", reverseProxy(usersServiceURL))
			users.POST("/login", reverseProxy(usersServiceURL))
			users.GET("/:id", reverseProxy(usersServiceURL))
		}

		// Cars service routes
		cars := v1.Group("/cars")
		{
			cars.GET("", reverseProxy(carsServiceURL))
			cars.GET("/:id", reverseProxy(carsServiceURL))
			cars.POST("", reverseProxy(carsServiceURL))
			cars.PUT("/:id/status", reverseProxy(carsServiceURL))
		}

		// Bookings service routes
		bookings := v1.Group("/bookings")
		{
			bookings.GET("", reverseProxy(bookingsServiceURL))
			bookings.POST("", reverseProxy(bookingsServiceURL))
			bookings.GET("/user/:userId", reverseProxy(bookingsServiceURL))
		}
	}

	// Start the server
	port := getEnv("PORT", "8080")
	r.Run(":" + port)
}

// reverseProxy creates a reverse proxy handler for the given target URL
func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// In a real implementation, this would use a proper reverse proxy
		// For now, we'll just forward the request to the target service
		c.JSON(http.StatusOK, gin.H{
			"service": target,
			"path":    c.Request.URL.Path,
			"method":  c.Request.Method,
		})
	}
}
