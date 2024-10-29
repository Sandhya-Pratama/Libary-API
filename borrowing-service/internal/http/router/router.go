package router

import (
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/http/handler"
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

func PublicRoutes(borrowingBookHandler *handler.BorrowingBookHandler) []*Route {
	return []*Route{
		// {
		// 	Method:  echo.GET,
		// 	Path:    "/BorrowingBooks",
		// 	Handler: borrowingBookHandler.GetAllBorrowingBooks,
		// },
	}
}

func PrivateRoutes(borrowingBookHandler *handler.BorrowingBookHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/borrow",
			Handler: borrowingBookHandler.BorrowBook,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/borrowing/:id",
			Handler: borrowingBookHandler.GetBorrowingsByUser,
			Role:    onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/return/:id",
			Handler: borrowingBookHandler.ReturnBook,
			Role:    onlyAdmin,
		},

		// {
		// 	Method:  echo.PUT,
		// 	Path:    "/BorrowingBooks/:id",
		// 	Handler: borrowingBookHandler.UpdateBorrowingBook,
		// 	Role:    onlyAdmin,
		// },

		// {
		// 	Method:  echo.GET,
		// 	Path:    "/BorrowingBooks/:id",
		// 	Handler: borrowingBookHandler.GetBorrowingBookByID,
		// 	Role:    onlyAdmin,
		// },

		// {
		// 	Method:  echo.DELETE,
		// 	Path:    "/BorrowingBooks/:id",
		// 	Handler: borrowingBookHandler.DeleteBorrowingBook,
		// 	Role:    onlyAdmin,
		// },
	}
}
