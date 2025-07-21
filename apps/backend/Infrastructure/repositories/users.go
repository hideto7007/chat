package repositories

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"chat/domain/valueObject/passwordHash"
	"chat/models"
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
    var userModels []models.User // スライスにする
    if err := r.db.WithContext(ctx).Find(&userModels).Error; err != nil {
        return nil, err
    }

    var users []entities.User
    for _, um := range userModels {
        users = append(users, entities.User{
            ID:       &um.ID,
            Name:     um.Name,
            Email:    um.Email,
            // Password: バリューオブジェクトに変換する場合はここで変換
        })
    }
    return users, nil
}

func (r *UserRepository) FindById(ctx context.Context, id uint) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		return nil, err
	}
	return &entities.User{
		ID:       &userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
	}, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error; err != nil {
		return nil, err
	}
	hashedPassword := passwordHash.NewPasswordHashFromHash(userModel.Password)
	return &entities.User{
		ID:       &userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: hashedPassword,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	var userModel models.User
	userModel.Name = user.Name
	userModel.Email = user.Email
	userModel.Password = user.Password.ToString()
	if err := r.db.WithContext(ctx).Create(&userModel).Error; err != nil {
		return nil, err
	}
    createdUser := &entities.User{
        ID:       &userModel.ID,
        Name:     userModel.Name,
        Email:    userModel.Email,
        Password: user.Password,
    }
    return createdUser, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	var userModel models.User
	userModel.ID = *user.ID
	userModel.Name = user.Name
	userModel.Email = user.Email
	userModel.Password = user.Password.ToString()
	if err := r.db.WithContext(ctx).Save(&userModel).Error; err != nil {
		return nil, err
	}
	return &entities.User{
		ID:       &userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
	}, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	var userModel models.User
	if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Delete(&userModel).Error; err != nil {
		return err
	}
	return nil
}