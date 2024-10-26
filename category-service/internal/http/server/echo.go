package server

import (
	"net/http"

	common "github.com/Sandhya-Pratama/Libary-API/category-service/common/middleware"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/http/binder"
	"github.com/Sandhya-Pratama/Libary-API/category-service/internal/http/router"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(
	cfg *config.Config,
	binder *binder.Binder,
	publicRoutes, privateRoutes []*router.Route) *Server {
	e := echo.New()

	e.HideBanner = true
	e.Binder = binder

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORS(),
		// session.Middleware(session.NewCookieStore([]byte(cfg.SessionSecret))),
	)

	//untuk deklarasi versi dari api
	v1 := e.Group("/v1")

	for _, public := range publicRoutes {
		//e.add = untuk menambahkan route baru.
		v1.Add(public.Method, public.Path, public.Handler)
	}

	//ketika sudah ingin menggunakan middleware, maka menambahkan private.Middleware.
	for _, private := range privateRoutes {
		v1.Add(private.Method, private.Path, private.Handler, JWTProtected(cfg.JWT.SecretKey), RBACMiddleware(private.Role...))
	}
	return &Server{e}
}

// func untuk pendeklarasian JWT Middleware
func JWTProtected(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &common.JwtCustomClaims{}
		},
		SigningKey: []byte(secretKey),
	})
}

// func untuk pendeklarasian RBAC Middleware
func RBACMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "silahkan login terlebih dahulu"})
			}

			claims := user.Claims.(*common.JwtCustomClaims)

			// Check if the user has the required role
			if !contains(roles, claims.Role) {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "anda tidak diperbolehkan untuk mengakses resource ini"})
			}

			return next(c)
		}
	}
}

// Helper function to check if a string is in a slice of strings
func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}
