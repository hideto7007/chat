package main

import (
	"chat/controllers"
	libSwagger "chat/lib/swagger"
	models "chat/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    db := models.InitDB()
    controllers.ApiManagerWithDB(r, db) 
    libSwagger.Register(r)
    log.Println("Server is running on port 8080")
    r.Run(":8080")
}

