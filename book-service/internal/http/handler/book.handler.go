package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/Libary-API/book-service/entity"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/service"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService service.BookUseCase
}

func NewBookHandler(bookService service.BookUseCase) *BookHandler {
	return &BookHandler{bookService}
}

func (h *BookHandler) GetAllBooks(ctx echo.Context) error {
	books, err := h.bookService.GetAllBooks(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, books)
}

func (h *BookHandler) CreateBook(ctx echo.Context) error {
	var input struct {
		Title      string `json:"title" validate:"required"`
		ISBN       int64  `json:"isbn" validate:"required"`
		AuthorID   int64  `json:"author_id" validate:"required"`
		CategoryID int64  `json:"category_id" validate:"required"`
		Stock      int64  `json:"stock" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	book := entity.NewBook(input.Title, input.ISBN, input.AuthorID, input.CategoryID, input.Stock)
	err := h.bookService.CreateBook(ctx.Request().Context(), book)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBookByID(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	book, err := h.bookService.GetBookByID(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(ctx echo.Context) error {
	var input struct {
		ID         int64  `json:"id" validate:"required"`
		Title      string `json:"title" validate:"required"`
		ISBN       int64  `json:"isbn" validate:"required"`
		AuthorID   int64  `json:"author_id" validate:"required"`
		CategoryID int64  `json:"category_id" validate:"required"`
		Stock      int64  `json:"stock" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	book := entity.NewBook(input.Title, input.ISBN, input.AuthorID, input.CategoryID, input.Stock)
	book.ID = input.ID
	err := h.bookService.UpdateBook(ctx.Request().Context(), book)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update book"})
}

func (h *BookHandler) DeleteBook(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	err := h.bookService.DeleteBook(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully delete book"})
}
