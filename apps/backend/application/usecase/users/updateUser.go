package users

import (
	"chat/domain/entities"
	"chat/domain/password"
	"chat/domain/repositories"
	"context"
	"fmt"
)

type UpdateUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewUpdateUserUseCase(userRepository repositories.UsersRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, userDto *UpdateUserDto) (*entities.User, error) {
	user, _ := uc.userRepository.FindById(ctx, userDto.ID)
	if user == nil {
		return nil, fmt.Errorf("user not found with id: %d", userDto.ID)
	}
	passwordHash, err := password.NewPasswordHash(userDto.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	newUser := &entities.User{
		ID:       &userDto.ID,
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: passwordHash,
	}
	newUser.Update()

	user, err = uc.userRepository.Update(ctx, newUser)

	if err != nil {
		return nil, fmt.Errorf("failed to Update user: %w", err)
	}

	return user, nil
}

type UpdateUserDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewUpdateUserDto(id uint, name, email, password string) *UpdateUserDto {
	return &UpdateUserDto{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}