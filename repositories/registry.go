package repositories

import (
	repositories "user-service/repositories/user"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IUserRepositoryRegistry interface {
	GetUser() repositories.IUserRepository
}

func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}
