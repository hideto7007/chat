package repositories

import (
	"chat/domain/entities"
	"context"
)

type UsersRepositoryInterface interface {
	FindAll(ctx context.Context) ([]entities.User, error)
	FindById(ctx context.Context, id uint) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Delete(ctx context.Context, id uint) error
}