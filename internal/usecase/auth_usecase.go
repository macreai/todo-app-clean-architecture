package usecase

import (
	"errors"

	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"github.com/macreai/todo-app-clean-architecture/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo repository.UserRepository
}

func NewAuthUseCase(repo repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.repo.CreateUser(user)
}

func (u *AuthUsecase) Login(username, password string) (*domain.User, error) {
	user, err := u.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	return user, nil
}
