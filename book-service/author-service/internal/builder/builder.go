package builder

import (
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/http/handler"
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/http/router"
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/repository"
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a Author handler
	authorRepository := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorService)

	// Menggunakan PublicRoutes dengan kedua handler
	return router.PublicRoutes(authorHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a Author handler
	authorRepository := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(authorHandler)
}
