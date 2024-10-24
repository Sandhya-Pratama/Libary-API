package handler

import (
	"log"
	"net/http"

	"github.com/Sandhya-Pratama/Libary-API/entity"
	"github.com/Sandhya-Pratama/Libary-API/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	registrationService service.RegistrationUseCase
	loginService        service.LoginUseCase
	tokenService        service.TokenUseCase
}

func NewAuthHandler(
	registrationService service.RegistrationUseCase,
	loginService service.LoginUseCase,
	tokenService service.TokenUseCase) *AuthHandler {
	return &AuthHandler{
		registrationService: registrationService,
		loginService:        loginService,
		tokenService:        tokenService,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var input struct {
		Username string `json:"username"  validate:"required" `
		Password string `json:"password" validate:"required"`
	}

	// Bind data dari request ke struct input
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Login service dipanggil dengan urutan yang benar: username dulu, lalu password
	user, err := h.loginService.Login(ctx.Request().Context(), input.Username, input.Password)
	if err != nil {
		log.Println("Login failed for user:", input.Username, "Error:", err)
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": err.Error(),
		})
	}

	// Generate token jika user ditemukan dan login berhasil
	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	data := map[string]interface{}{
		"access_token": accessToken,
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h *AuthHandler) Registration(ctx echo.Context) error {
	//pengecekan request
	var input struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Role     string `json:"role" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//untuk manggil registration service di folder service
	user := entity.Register(input.Username, input.Email, input.Password, input.Role)
	err := h.registrationService.Registration(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":      "User registration successfully",
		"access_token": accessToken,
	})

}
