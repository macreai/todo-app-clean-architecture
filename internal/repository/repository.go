package repository

import "github.com/macreai/todo-app-clean-architecture/internal/domain"

type ActivityUserRepository interface {
	Create(user *domain.ActivityUser) error
	GetByID(id uint) (*domain.ActivityUser, error)
	GetAll() ([]*domain.ActivityUser, error)
	Update(user *domain.ActivityUser) error
	Delete(id uint) error
}

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
}
