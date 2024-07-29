package postgres

import (
	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"github.com/macreai/todo-app-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type PostgresActivityUserRepository struct {
	DB *gorm.DB
}

func NewPostgresActivityUserRepository(db *gorm.DB) repository.ActivityUserRepository {
	return &PostgresActivityUserRepository{DB: db}
}

// Create implements repository.ActivityUserRepository.
func (p *PostgresActivityUserRepository) Create(user *domain.ActivityUser) error {
	return p.DB.Create(user).Error
}

// Delete implements repository.ActivityUserRepository.
func (p *PostgresActivityUserRepository) Delete(id uint) error {
	return p.DB.Delete(&domain.ActivityUser{}, id).Error
}

// GetByID implements repository.ActivityUserRepository.
func (p *PostgresActivityUserRepository) GetByID(id uint) (*domain.ActivityUser, error) {
	var activityUser domain.ActivityUser
	err := p.DB.First(&activityUser, id).Error
	return &activityUser, err
}

// GetAll implements repository.ActivityUserRepository.
func (p *PostgresActivityUserRepository) GetAll() ([]*domain.ActivityUser, error) {
	var activityUser []*domain.ActivityUser
	err := p.DB.Find(&activityUser).Error
	return activityUser, err
}

// Update implements repository.ActivityUserRepository.
func (p *PostgresActivityUserRepository) Update(user *domain.ActivityUser) error {
	return p.DB.Save(&user).Error
}
