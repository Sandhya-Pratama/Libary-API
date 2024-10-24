package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/Libary-API/entity"
	"github.com/Sandhya-Pratama/Libary-API/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUseCase
}

// melakukan instace dari user handler
func NewUserHandler(userService service.UserUseCase) *UserHandler {
	return &UserHandler{userService}
}

// func untuk melakukan getAll User
func (h *UserHandler) GetAllUsers(ctx echo.Context) error {
	users, err := h.userService.GetAllUsers(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Username string `json:"usename" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password"`
		Role     string `json:"role" validate:"oneof=Admin User"`
	}
	//ini func untuk error checking
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user := entity.NewUser(input.Username, input.Email, input.Password, input.Role)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	//kalau retrun nya kaya gini akan tampil pesan "User created successfully"
	return ctx.JSON(http.StatusCreated, "User created successfully")
}

// func untuk melakukan updateUser by id
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `json:"id"`
		Username string `json:"username"` // Diperbaiki dari "usename" ke "username"
		Email    string `json:"email" validate:"email"`
		Password string `json:"password"`
		Role     string `json:"role" validate:"oneof=Admin User"`
	}

	// Bind JSON input
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Username, input.Email, input.Password, input.Role)

	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update user"})
}

// func untuk melakukan deleteUser by id
func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `json:"id"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	err := h.userService.DeleteUser(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully delete user"})
}
