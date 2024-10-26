package builder

import (
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/http/handler"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/http/router"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/repository"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a BorrowingBook handler
	borrowingBookRepository := repository.NewBorrowingBookRepository(db)
	borrowingBookService := service.NewBorrowingBookService(borrowingBookRepository)
	borrowingBookHandler := handler.NewBorrowingBookHandler(borrowingBookService)

	// Menggunakan PublicRoutes dengan kedua handler
	return router.PublicRoutes(borrowingBookHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a BorrowingBook handler
	borrowingBookRepository := repository.NewBorrowingBookRepository(db)
	borrowingBookService := service.NewBorrowingBookService(borrowingBookRepository)
	borrowingBookHandler := handler.NewBorrowingBookHandler(borrowingBookService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(borrowingBookHandler)
}
