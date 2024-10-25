package builder

import (
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/handler"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/router"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/repository"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a book handler
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// Menggunakan PublicRoutes dengan kedua handler
	return router.PublicRoutes(bookHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a book handler
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(bookHandler)
}
