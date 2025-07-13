package users

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"context"
	"fmt"
)

type CreateUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewCreateUserUseCase(userRepository repositories.UsersRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, userDto *CreateUserDto) (*entities.User, error) {
	newUser, err := entities.NewUser(
		userDto.Name,
		userDto.Email,
		userDto.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new user entity: %w", err)
	}

	user, err := uc.userRepository.Create(ctx, newUser)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

type CreateUserDto struct {
	ID       *uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateUserDto(u *entities.User) *CreateUserDto {
	return &CreateUserDto{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}