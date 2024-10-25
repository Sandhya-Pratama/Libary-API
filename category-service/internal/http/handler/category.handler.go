package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/Libary-API/category-service/entity"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/service"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryService service.CategoryUseCase
}

func NewCategoryHandler(CategoryService service.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{CategoryService}
}

func (h *CategoryHandler) GetAllCategorys(ctx echo.Context) error {
	Categorys, err := h.CategoryService.GetAllCategorys(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, Categorys)
}

func (h *CategoryHandler) CreateCategory(ctx echo.Context) error {
	var input struct {
		Name string `json:"name" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	Category := entity.NewCategory(input.Name)
	err := h.CategoryService.CreateCategory(ctx.Request().Context(), Category)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusCreated, Category)
}

func (h *CategoryHandler) UpdateCategory(ctx echo.Context) error {
	var input struct {
		id   int64  `json:"id" validate:"required"`
		name string `json:"name" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	Category := entity.NewCategory(input.name)
	Category.ID = input.id
	err := h.CategoryService.UpdateCategory(ctx.Request().Context(), Category)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update Category"})
}

func (h *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	err := h.CategoryService.DeleteCategory(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully delete Category"})
}
