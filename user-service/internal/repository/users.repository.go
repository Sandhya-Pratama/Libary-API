package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sandhya-Pratama/Libary-API/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	//menggunakan db untuk melakukan query ke database
	err := r.db.WithContext(ctx).Create(&user).Error // pada line ini akan melakukan query "INSERT INTO users"
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User

	// Melakukan query untuk mencari user berdasarkan username
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		// Cek apakah errornya karena record tidak ditemukan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Bisa return error yang lebih deskriptif atau custom error
			return nil, fmt.Errorf("user dengan username %s tidak ditemukan", username)
		}
		// Jika error lain, kembalikan error tersebut
		return nil, err
	}

	// Jika user ditemukan, kembalikan data user
	return &user, nil
}
