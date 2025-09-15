package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/carcius-rent-car/bookings-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(&models.Booking{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Seed initial data if the table is empty
	seedInitialData()
}

func seedInitialData() {
	var count int64
	DB.Model(&models.Booking{}).Count(&count)

	if count == 0 {
		// Sample bookings (assuming users and cars with these IDs exist)
		bookings := []models.Booking{
			{
				UserID:     1,
				CarID:      1,
				StartDate:  time.Now().Add(24 * time.Hour),
				EndDate:    time.Now().Add(72 * time.Hour),
				Status:     models.BookingStatusConfirmed,
				TotalCost:  150.00,
			},
			{
				UserID:     2,
				CarID:      2,
				StartDate:  time.Now().Add(24 * time.Hour),
				EndDate:    time.Now().Add(48 * time.Hour),
				Status:     models.BookingStatusPending,
				TotalCost:  130.00,
			},
		}

		for _, booking := range bookings {
			DB.Create(&booking)
		}
	}
}
