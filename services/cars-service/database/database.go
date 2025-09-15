package database

import (
	"fmt"
	"log"
	"os"

	"github.com/carcius-rent-car/cars-service/models"
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
	err = DB.AutoMigrate(&models.Car{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Seed initial data if the table is empty
	seedInitialData()
}

func seedInitialData() {
	var count int64
	DB.Model(&models.Car{}).Count(&count)

	if count == 0 {
		cars := []models.Car{
			{
				Make:         "Toyota",
				Model:        "Camry",
				Year:         2022,
				Color:        "Silver",
				LicensePlate: "ABC123",
				PricePerDay:  50.00,
				Status:       models.CarStatusAvailable,
				Mileage:      15000,
				Seats:        5,
				Doors:        4,
				Transmission: "Automatic",
				FuelType:     "Petrol",
				ImageURL:     "https://example.com/camry.jpg",
				Description:  "Comfortable mid-size sedan",
			},
			{
				Make:         "Honda",
				Model:        "CR-V",
				Year:         2023,
				Color:        "Black",
				LicensePlate: "XYZ789",
				PricePerDay:  65.00,
				Status:       models.CarStatusAvailable,
				Mileage:      5000,
				Seats:        5,
				Doors:        5,
				Transmission: "Automatic",
				FuelType:     "Hybrid",
				ImageURL:     "https://example.com/crv.jpg",
				Description:  "Spacious and fuel-efficient SUV",
			},
		}

		for _, car := range cars {
			DB.Create(&car)
		}
	}
}
