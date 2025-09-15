package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
	BookingStatusCompleted BookingStatus = "completed"
)

type Booking struct {
	gorm.Model
	UserID    uint          `json:"user_id" gorm:"not null"`
	CarID     uint          `json:"car_id" gorm:"not null"`
	StartDate time.Time     `json:"start_date" gorm:"not null"`
	EndDate   time.Time     `json:"end_date" gorm:"not null"`
	Status    BookingStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
	TotalCost float64       `json:"total_cost" gorm:"not null"`
}
