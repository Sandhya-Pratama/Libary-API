package main

import (
	"github.com/Sandhya-Pratama/Libary-API/API-gateway/config"
	"github.com/Sandhya-Pratama/Libary-API/API-gateway/internal/router"
	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi konfigurasi
	cfg := config.LoadConfig()

	// Inisialisasi layanan
	authorService := service.NewAuthorService(cfg.AuthorServiceURL)
	bookService := service.NewBookService(cfg.BookServiceURL)
	categoryService := service.NewCategoryService(cfg.CategoryServiceURL)
	userService := service.NewUserService(cfg.UserServiceURL)
	authService := service.NewAuthService(cfg.AuthServiceURL)

	// Inisialisasi handler
	authorHandler := handler.NewAuthorHandler(authorService)
	bookHandler := handler.NewBookHandler(bookService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	// Inisialisasi Echo
	e := echo.New()

	// Inisialisasi routes
	router.InitRoutes(e, authorHandler, bookHandler, categoryHandler, userHandler, authHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
