package service

import (
	"context"
	"errors"
	"log"

	"github.com/Sandhya-Pratama/Libary-API/entity"
	"golang.org/x/crypto/bcrypt" // Untuk hash password
	"gorm.io/gorm"
)

type LoginUseCase interface {
	Login(ctx context.Context, username string, password string) (*entity.User, error)
}

type LoginRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type loginService struct {
	repository LoginRepository
}

func NewLoginService(repository LoginRepository) *loginService {
	return &loginService{
		repository: repository,
	}
}

func (s *loginService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	// Cari user berdasarkan username
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		// Tangani error jika user tidak ditemukan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not found:", username)
			return nil, errors.New("user not found")
		}
		log.Println("Error fetching user:", err)
		return nil, err
	}

	// Verifikasi password menggunakan bcrypy
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect login credentials")
	}

	// Jika semua berhasil, kembalikan user
	return user, nil
}
