package repository

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
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
