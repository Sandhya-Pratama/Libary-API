package router

import (
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/http/handler"
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

func PublicRoutes(CategoryHandler *handler.CategoryHandler) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/Categorys",
			Handler: CategoryHandler.GetAllCategorys,
		},
	}
}

func PrivateRoutes(CategoryHandler *handler.CategoryHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/Categorys",
			Handler: CategoryHandler.CreateCategory,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/Categorys/:id",
			Handler: CategoryHandler.UpdateCategory,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.DELETE,
			Path:    "/Categorys/:id",
			Handler: CategoryHandler.DeleteCategory,
			Role:    onlyAdmin,
		},
	}
}
