package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/entity"
)

type BorrowingBookUseCase interface {
	BorrowBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error
	ReturnBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error
	GetBorrowingsByUser(ctx context.Context, userID int64) ([]*entity.BorrowingBook, error)
}

type BorrowingBookRepository interface {
	BorrowBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error
	ReturnBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error
	GetBorrowingsByUser(ctx context.Context, userID int64) ([]*entity.BorrowingBook, error)
}

type BorrowingBookService struct {
	repository BorrowingBookRepository
}

func NewBorrowingBookService(repository BorrowingBookRepository) *BorrowingBookService {
	return &BorrowingBookService{
		repository: repository,
	}
}

func (s *BorrowingBookService) BorrowBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error {
	return s.repository.BorrowBook(ctx, borrowingBook)
}

func (s *BorrowingBookService) ReturnBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error {
	return s.repository.ReturnBook(ctx, borrowingBook)
}

func (s *BorrowingBookService) GetBorrowingsByUser(ctx context.Context, userID int64) ([]*entity.BorrowingBook, error) {
	return s.repository.GetBorrowingsByUser(ctx, userID)
}
