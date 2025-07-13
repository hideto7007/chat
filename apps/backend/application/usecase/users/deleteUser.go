package users

import (
	"chat/domain/repositories"
	"context"
	"fmt"
)

type DeleteUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewDeleteUserUseCase(userRepository repositories.UsersRepositoryInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, id string) error {
	if err := uc.userRepository.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
