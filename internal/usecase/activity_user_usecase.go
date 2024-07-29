package usecase

import (
	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"github.com/macreai/todo-app-clean-architecture/internal/repository"
)

type ActivityUserUsecase struct {
	repo repository.ActivityUserRepository
}

func NewActivityUserUsecase(repo repository.ActivityUserRepository) *ActivityUserUsecase {
	return &ActivityUserUsecase{repo: repo}
}

func (u *ActivityUserUsecase) Create(user *domain.ActivityUser) error {
	return u.repo.Create(user)
}

func (u *ActivityUserUsecase) GetByID(id uint) (*domain.ActivityUser, error) {
	return u.repo.GetByID(id)
}

func (u *ActivityUserUsecase) GetAll() ([]*domain.ActivityUser, error) {
	return u.repo.GetAll()
}

func (u *ActivityUserUsecase) Update(user *domain.ActivityUser) error {
	return u.repo.Update(user)
}

func (u *ActivityUserUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
