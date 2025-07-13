package users

import (
	"chat/domain/entities"
	"chat/domain/repositories"
	"context"
)

type ListUserUseCase struct {
	userRepository repositories.UsersRepositoryInterface
}

func NewListUserUseCase(userRepository repositories.UsersRepositoryInterface) *ListUserUseCase {
	return &ListUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *ListUserUseCase) Execute(ctx context.Context) ([]ListUserDto, error) {
	users, err := uc.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var userModels []ListUserDto
	for _, u := range users {
		userModels = append(userModels, NewListUserDto(u))
	}

	return userModels, nil
}

type ListUserDto struct {
	ID    *uint  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewListUserDto(u entities.User) ListUserDto {
    return ListUserDto{
        ID:    u.ID,
        Name:  u.Name,
        Email: u.Email,
    }
}