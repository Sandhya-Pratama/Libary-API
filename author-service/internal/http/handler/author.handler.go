package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/Libary-API/author-service/entity"
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/author-service/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	AuthorService service.AuthorUseCase
}

func NewAuthorHandler(AuthorService service.AuthorUseCase) *AuthorHandler {
	return &AuthorHandler{AuthorService}
}

func (h *AuthorHandler) GetAllAuthors(ctx echo.Context) error {
	Authors, err := h.AuthorService.GetAllAuthors(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, Authors)
}

func (h *AuthorHandler) CreateAuthor(ctx echo.Context) error {
	var input struct {
		Name string `json:"name" validate:"required"`
		Bio  string `json:"bio" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	Author := entity.NewAuthor(input.Name, input.Bio)
	err := h.AuthorService.CreateAuthor(ctx.Request().Context(), Author)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusCreated, Author)
}

func (h *AuthorHandler) UpdateAuthor(ctx echo.Context) error {
	var input struct {
		id   int64  `json:"id" validate:"required"`
		name string `json:"name" validate:"required"`
		bio  string `json:"bio" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	Author := entity.NewAuthor(input.name, input.bio)
	Author.ID = input.id
	err := h.AuthorService.UpdateAuthor(ctx.Request().Context(), Author)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update Author"})
}

func (h *AuthorHandler) DeleteAuthor(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	err := h.AuthorService.DeleteAuthor(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully delete Author"})
}
