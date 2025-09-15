package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName     string `json:"full_name" gorm:"not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
}
