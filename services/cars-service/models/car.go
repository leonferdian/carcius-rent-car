package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type CarStatus string

const (
	CarStatusAvailable CarStatus = "available"
	CarStatusRented   CarStatus = "rented"
	CarStatusMaintenance CarStatus = "maintenance"
)

type Car struct {
	gorm.Model
	Make         string    `json:"make" gorm:"not null"`
	Model        string    `json:"model" gorm:"not null"`
	Year         int       `json:"year" gorm:"not null"`
	Color        string    `json:"color"`
	LicensePlate string    `json:"license_plate" gorm:"unique;not null"`
	PricePerDay  float64   `json:"price_per_day" gorm:"not null"`
	Status       CarStatus `json:"status" gorm:"type:varchar(20);default:'available'"`
	Mileage      int       `json:"mileage"`
	LastService  time.Time `json:"last_service"`
	NextService  time.Time `json:"next_service"`
	ImageURL     string    `json:"image_url"`
	Description  string    `json:"description"`
	Seats        int       `json:"seats"`
	Doors        int       `json:"doors"`
	Transmission string    `json:"transmission"` // automatic, manual
	FuelType     string    `json:"fuel_type"`    // petrol, diesel, electric, hybrid
	Features     string    `json:"features"`     // JSON string of features
}
