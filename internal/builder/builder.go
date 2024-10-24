package builder

import (
	"github.com/Sandhya-Pratama/Libary-API/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/internal/http/router"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) []*router.Route {

	return router.PublicRoutes()
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) []*router.Route {
	// Create a user handler

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler)
}
