package user

import (
	usecase "chat/application/usecase/users"
	"chat/domain/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, repo repositories.UsersRepositoryInterface) {
    userController := NewUserController(
        usecase.NewListUserUseCase(repo),
        usecase.NewCreateUserUseCase(repo),
        usecase.NewDeleteUserUseCase(repo),
        usecase.NewUpdateUserUseCase(repo),
        usecase.NewGetEmailUserUseCase(repo),
        usecase.NewGetIdUserUseCase(repo),
    )

    r.GET("/users", userController.List)
    r.GET("/users/id/:id", userController.GetIdUser)
    r.GET("/users/email/:email", userController.GetEmailUser)
    r.POST("/users", userController.Create)
    r.PUT("/users", userController.Update)
    r.DELETE("/users/:id", userController.Delete)
}