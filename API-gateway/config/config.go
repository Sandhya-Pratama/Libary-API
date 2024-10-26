package config

import (
	"os"
)

type Config struct {
	Port         string
	BorrowingSvc string
	BookSvc      string
	AuthorSvc    string
	CategorySvc  string
	UserSvc      string
	JwtSecret    string
}

func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("API_GATEWAY_PORT", "8080"),
		BorrowingSvc: getEnv("BORROWING_SERVICE", "localhost:8081"),
		BookSvc:      getEnv("BOOK_SERVICE", "localhost:8082"),
		AuthorSvc:    getEnv("AUTHOR_SERVICE", "localhost:8083"),
		CategorySvc:  getEnv("CATEGORY_SERVICE", "localhost:8084"),
		UserSvc:      getEnv("USER_SERVICE", "localhost:8085"),
		JwtSecret:    getEnv("JWT_SECRET", "your_jwt_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
