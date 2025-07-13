package repositories

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"context"

	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UsersRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *UserRepository) FindById(ctx context.Context, id uint) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func (r *UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}