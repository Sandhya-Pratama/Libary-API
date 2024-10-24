package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/entity"
)

type UserUseCase interface {
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
}

// interface untuk repository
// untuk memanggil repository
// GetAll = untuk menampilkan semua data user, dan itu harus sama dengan yang ada di repository
type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
}

// code di line 23 merupakan dependency injection, karena repository tidak langsung di panggil.
// karena repository dipanggil melalui code pada line 18
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

// func dibawah ini untuk type user usecase
// ini untuk membuat data user
func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repository.CreateUser(ctx, user)
}

// untuk update data user
func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.repository.UpdateUser(ctx, user)
}
