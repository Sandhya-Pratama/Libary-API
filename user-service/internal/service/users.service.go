package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserService struct {
	repository UserRepository
}

// func untuk UserRepository
func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository}
}

// untuk memanggil repository
func (s *UserService) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	return s.repository.GetAllUsers(ctx)
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repository.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUSer(ctx context.Context, id int64) error {
	return s.repository.DeleteUser(ctx, id)
}
