package postgres

import (
	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"github.com/macreai/todo-app-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type PostgreUserRepository struct {
	DB *gorm.DB
}

func NewPostgreUserRepository(db *gorm.DB) repository.UserRepository {
	return &PostgreUserRepository{DB: db}
}

// CreateUser implements repository.UserRepository.
func (p *PostgreUserRepository) CreateUser(user *domain.User) error {
	return p.DB.Create(user).Error
}

// GetUserByUsername implements repository.UserRepository.
func (p *PostgreUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := p.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
