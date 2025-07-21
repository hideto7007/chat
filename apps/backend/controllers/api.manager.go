package controllers

import (
	"chat/controllers/user"
	factory "chat/infrastructure/repositories/factory"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiManagerWithDB(r *gin.Engine, db *gorm.DB) {
    api := r.Group("/api")
    repos := factory.NewRepositoryFactory(db)
    user.RegisterUserRoutes(api, repos.UserRepository)
}