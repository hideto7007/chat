package main

import (
	libGorm "chat/lib/gorm"
	models "chat/models"
)

func main() {
    db := models.InitDB()
    libGorm.Migrate(db)
}