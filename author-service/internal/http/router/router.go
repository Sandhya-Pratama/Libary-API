package router

import (
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/http/handler"
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

func PublicRoutes(authorHandler *handler.AuthorHandler) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/authors",
			Handler: authorHandler.GetAllAuthors,
		},
	}
}

func PrivateRoutes(authorHandler *handler.AuthorHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/authors",
			Handler: authorHandler.CreateAuthor,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/authors/:id",
			Handler: authorHandler.UpdateAuthor,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.DELETE,
			Path:    "/authors/:id",
			Handler: authorHandler.DeleteAuthor,
			Role:    onlyAdmin,
		},
	}
}
