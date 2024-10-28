package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	AuthorService service.AuthorUseCase
}

func NewAuthorHandler(authorService service.AuthorUseCase) *AuthorHandler {
	return &AuthorHandler{authorService}
}

func (h *AuthorHandler) GetAllAuthors(ctx echo.Context) error {
	authors, err := h.AuthorService.GetAllAuthors(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, authors)
}
