package users

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"context"
	"fmt"
)

type GetIdUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewGetIdUserUseCase(userRepository repositories.UsersRepositoryInterface) *GetIdUserUseCase {
	return &GetIdUserUseCase{
		userRepository: userRepository,
	}
}
func (uc *GetIdUserUseCase) Execute(ctx context.Context, id string) (*GetIdUserDto, error) {
	user, err := uc.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found with id: %s", id)
	}

	return NewGetIdUserDto(user), nil
}

type GetIdUserDto struct {
	ID    *uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGetIdUserDto(u *entities.User) *GetIdUserDto {
	return &GetIdUserDto{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}