package users

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"context"
	"fmt"
)	

type GetEmailUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewGetEmailUserUseCase(userRepository repositories.UsersRepositoryInterface) *GetEmailUserUseCase {
	return &GetEmailUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetEmailUserUseCase) Execute(ctx context.Context, email string) (*GetEmailUserDto, error) {
	user, err := uc.userRepository.FindByEmail(ctx, email)
    if err != nil {
        return nil, fmt.Errorf("failed to find user by email: %w", err)
    }
	if user == nil {
		return nil, fmt.Errorf("user not found with email: %s", email)
	}
	return NewGetEmailUserDto(user), nil
}

type GetEmailUserDto struct {
	ID    *uint  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGetEmailUserDto(u *entities.User) *GetEmailUserDto {
	return &GetEmailUserDto{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}