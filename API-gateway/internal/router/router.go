package router

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, authorHandler *handler.AuthorHandler, bookHandler *handler.BookHandler, categoryHandler *handler.CategoryHandler, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {
	api := e.Group("/api")

	// Routes untuk Author Service
	api.GET("/authors", authorHandler.GetAllAuthors)
	api.POST("/authors", authorHandler.CreateAuthor)
	api.PUT("/authors/:id", authorHandler.UpdateAuthor)
	api.DELETE("/authors/:id", authorHandler.DeleteAuthor)

	// Routes untuk Book Service
	api.GET("/books", bookHandler.GetAllBooks)
	api.POST("/books", bookHandler.CreateBook)
	api.PUT("/books/:id", bookHandler.UpdateBook)
	api.DELETE("/books/:id", bookHandler.DeleteBook)

	// Routes untuk Category Service
	api.GET("/categories", categoryHandler.GetAllCategories)
	api.POST("/categories", categoryHandler.CreateCategory)
	api.PUT("/categories/:id", categoryHandler.UpdateCategory)
	api.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Routes untuk User Service
	api.GET("/users", userHandler.GetAllUsers)
	api.POST("/users", userHandler.CreateUser)
	api.PUT("/users/:id", userHandler.UpdateUser)
	api.DELETE("/users/:id", userHandler.DeleteUser)

	// Routes untuk Auth Service
	api.POST("/auth/login", authHandler.Login)
}
