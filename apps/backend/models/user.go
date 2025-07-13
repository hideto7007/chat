package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"-"`
    CreatedAt time.Time `gorm:"default:now()"`
    UpdatedAt time.Time `gorm:"default:now()"`
}