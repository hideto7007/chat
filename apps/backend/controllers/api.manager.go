package controllers

import (
	"chat/Infrastructure/repositories"
	"chat/controllers/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiManagerWithDB(r *gin.Engine, db *gorm.DB) {
    api := r.Group("/api")
    repos := repositories.NewRepositoryFactory(db)
    user.RegisterUserRoutes(api, repos.UserRepository)
}