package config

import (
	"os"
)

type Config struct {
	JWTSecret            string
	AuthorService        string
	BookService          string
	CategoryService      string
	UserService          string
	BorrowingBookService string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret:            os.Getenv("JWT_SECRET"),
		AuthorService:        os.Getenv("AUTHOR_SERVICE"),
		BookService:          os.Getenv("BOOK_SERVICE"),
		CategoryService:      os.Getenv("CATEGORY_SERVICE"),
		UserService:          os.Getenv("USER_SERVICE"),
		BorrowingBookService: os.Getenv("BORROWING_BOOK_SERVICE"),
	}
}
