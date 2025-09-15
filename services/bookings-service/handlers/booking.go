package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/carcius-rent-car/bookings-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingHandler struct {
	DB *gorm.DB
}

func NewBookingHandler(db *gorm.DB) *BookingHandler {
	return &BookingHandler{DB: db}
}

type CreateBookingRequest struct {
	UserID    uint      `json:"user_id" binding:"required"`
	CarID     uint      `json:"car_id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

type UpdateBookingStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=confirmed cancelled completed"`
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate dates
	if req.StartDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start date must be in the future"})
		return
	}

	if req.EndDate.Before(req.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End date must be after start date"})
		return
	}

	// Check car availability
	var existingBookings int64
	h.DB.Model(&models.Booking{}).
		Where("car_id = ? AND status = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			req.CarID, models.BookingStatusConfirmed,
			req.StartDate, req.StartDate,
			req.EndDate, req.EndDate,
			req.StartDate, req.EndDate).
		Count(&existingBookings)

	if existingBookings > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Car is not available for the selected dates"})
		return
	}

	// In a real application, we would fetch the car's price per day from the Cars service
	// For now, we'll use a default value
	pricePerDay := 50.0
	duration := req.EndDate.Sub(req.StartDate).Hours() / 24
	if duration < 1 {
		duration = 1
	}
	totalCost := pricePerDay * duration

	booking := models.Booking{
		UserID:    req.UserID,
		CarID:     req.CarID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    models.BookingStatusPending,
		TotalCost: totalCost,
	}

	if err := h.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func (h *BookingHandler) GetUserBookings(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var bookings []models.Booking
	if err := h.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func (h *BookingHandler) UpdateBookingStatus(c *gin.Context) {
	bookingID := c.Param("id")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking ID is required"})
		return
	}

	var req UpdateBookingStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var booking models.Booking
	if err := h.DB.First(&booking, bookingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	booking.Status = models.BookingStatus(req.Status)

	if err := h.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	c.JSON(http.StatusOK, booking)
}
