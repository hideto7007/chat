package repositories

import (
	domainRepo "chat/domain/repositories"

	"gorm.io/gorm"
)

type RepositoryFactory struct {
    UserRepository domainRepo.UsersRepositoryInterface
    // 他のリポジトリも追加可能
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
    return &RepositoryFactory{
        UserRepository: NewUserRepository(db),
    }
}