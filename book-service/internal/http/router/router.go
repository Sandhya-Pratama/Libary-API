package router

import (
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/handler"
	"github.com/labstack/echo/v4"
)

const (
	Admin = "Admin"
	User  = "User"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Role    []string
}

func PublicRoutes(bookHandler *handler.BookHandler) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/books",
			Handler: bookHandler.GetAllBooks,
		},
	}
}

func PrivateRoutes(bookHandler *handler.BookHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: bookHandler.CreateBook,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/books/:id",
			Handler: bookHandler.UpdateBook,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.DELETE,
			Path:    "/books/:id",
			Handler: bookHandler.DeleteBook,
			Role:    onlyAdmin,
		},
	}
}
