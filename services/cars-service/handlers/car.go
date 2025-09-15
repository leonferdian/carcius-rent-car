package handlers

import (
	"net/http"
	"strconv"

	"github.com/carcius-rent-car/cars-service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarHandler struct {
	DB *gorm.DB
}

func NewCarHandler(db *gorm.DB) *CarHandler {
	return &CarHandler{DB: db}
}

type CreateCarRequest struct {
	Make         string  `json:"make" binding:"required"`
	Model        string  `json:"model" binding:"required"`
	Year         int     `json:"year" binding:"required"`
	Color        string  `json:"color"`
	LicensePlate string  `json:"license_plate" binding:"required"`
	PricePerDay  float64 `json:"price_per_day" binding:"required"`
	Status       string  `json:"status"`
	Mileage      int     `json:"mileage"`
	Seats        int     `json:"seats"`
	Doors        int     `json:"doors"`
	Transmission string  `json:"transmission"`
	FuelType     string  `json:"fuel_type"`
	ImageURL     string  `json:"image_url"`
	Description  string  `json:"description"`
}

type UpdateCarStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=available rented maintenance"`
}

func (h *CarHandler) GetCars(c *gin.Context) {
	var cars []models.Car
	if err := h.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cars"})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) GetCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car

	if err := h.DB.First(&car, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch car"})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	var req CreateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := models.Car{
		Make:         req.Make,
		Model:        req.Model,
		Year:         req.Year,
		Color:        req.Color,
		LicensePlate: req.LicensePlate,
		PricePerDay:  req.PricePerDay,
		Status:       models.CarStatus(req.Status),
		Mileage:      req.Mileage,
		Seats:        req.Seats,
		Doors:        req.Doors,
		Transmission: req.Transmission,
		FuelType:     req.FuelType,
		ImageURL:     req.ImageURL,
		Description:  req.Description,
	}

	if err := h.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	c.JSON(http.StatusCreated, car)
}

func (h *CarHandler) UpdateCarStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	var req UpdateCarStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var car models.Car
	if err := h.DB.First(&car, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car status"})
		return
	}

	car.Status = models.CarStatus(req.Status)

	if err := h.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car status"})
		return
	}

	c.JSON(http.StatusOK, car)
}
