package users

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"chat/domain/password"
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
    passObj := password.NewPasswordCheck(userDto.Password)
    if err := passObj.Validate(); err != nil {
        return nil, fmt.Errorf("invalid password: %w", err)
    }

	passwordHash, err := password.NewPasswordHash(userDto.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &entities.User{
		Name:  userDto.Name,
		Email: userDto.Email,
		Password: passwordHash,
	}

	user, err := uc.userRepository.Create(ctx, newUser)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateUserDto(name, email, password string) *CreateUserDto {
    return &CreateUserDto{
        Name:     name,
        Email:    email,
        Password: password,
    }
}