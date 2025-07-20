package users

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"context"
	"fmt"
)

type LoginUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewLoginUserUseCase(userRepository repositories.UsersRepositoryInterface) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, email, password string) (*entities.User, error) {
	user, err := uc.userRepository.FindByEmail(ctx, email)

	if user == nil {
		return nil, fmt.Errorf("invalid email")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if verify := user.Password.Verify(password); verify != true {
		return nil, fmt.Errorf("password verification failed")
	}

	return user, nil
}