package builder

import (
	"github.com/Sandhya-Pratama/Libary-API/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/internal/http/handler"
	"github.com/Sandhya-Pratama/Libary-API/internal/http/router"
	"github.com/Sandhya-Pratama/Libary-API/internal/repository"
	"github.com/Sandhya-Pratama/Libary-API/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)

	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)

	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)
	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler)
}
