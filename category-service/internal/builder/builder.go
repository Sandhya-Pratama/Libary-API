package builder

import (
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/http/handler"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/http/router"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/repository"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a category handler
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Menggunakan PublicRoutes dengan kedua handler
	return router.PublicRoutes(categoryHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a category handler
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(categoryHandler)
}
