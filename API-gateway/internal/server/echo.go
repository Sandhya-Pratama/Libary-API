package server

import (
	"api-gateway/common/middleware"
	"api-gateway/internal/config"
	"api-gateway/internal/http/handler"
	"api-gateway/internal/http/router"

	"github.com/labstack/echo/v4"
)

func InitServer(cfg *config.Config, handler *handler.Handler) *echo.Echo {
	e := echo.New()

	// Set up JWT middleware
	e.Use(middleware.JWTMiddleware(cfg.JWTSecret))

	// Initialize routes
	router.InitRoutes(e, handler)

	return e
}
