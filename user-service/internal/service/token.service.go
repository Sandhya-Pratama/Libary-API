package service

import (
	"context"
	"time"

	"github.com/Sandhya-Pratama/Libary-API/common"
	"github.com/Sandhya-Pratama/Libary-API/entity"
	"github.com/Sandhya-Pratama/Libary-API/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(ctx context.Context, username *entity.User) (string, error)
}

type TokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{cfg: cfg}
}

func (s *TokenService) GenerateAccessToken(ctx context.Context, username *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(10 * time.Minute).Unix()

	claims := common.JwtCustomClaims{
		ID:    username.ID,
		Email: username.Email,
		Role:  username.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiredTime, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
