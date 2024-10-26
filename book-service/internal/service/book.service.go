package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/book-service/entity"
)

type BookUseCase interface {
	GetAllBooks(ctx context.Context) ([]*entity.Book, error)
	CreateBook(ctx context.Context, book *entity.Book) error
	GetBookByID(ctx context.Context, id int64) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) error
	DeleteBook(ctx context.Context, id int64) error
}

type BookRepository interface {
	GetAllBooks(ctx context.Context) ([]*entity.Book, error)
	CreateBook(ctx context.Context, book *entity.Book) error
	GetBookByID(ctx context.Context, id int64) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) error
	DeleteBook(ctx context.Context, id int64) error
}

type BookService struct {
	repository BookRepository
}

func NewBookService(repository BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*entity.Book, error) {
	return s.repository.GetAllBooks(ctx)
}

func (s *BookService) CreateBook(ctx context.Context, book *entity.Book) error {
	return s.repository.CreateBook(ctx, book)
}

func (s *BookService) GetBookByID(ctx context.Context, id int64) (*entity.Book, error) {
	return s.repository.GetBookByID(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, book *entity.Book) error {
	return s.repository.UpdateBook(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id int64) error {
	return s.repository.DeleteBook(ctx, id)
}
