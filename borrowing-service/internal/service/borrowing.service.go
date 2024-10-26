package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/entity"
)

type BorrowingBookUseCase interface {
	// GetAllBorrowingBooks(ctx context.Context) ([]*entity.BorrowingBook, error)
	CreateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error
	// GetBorrowingBookByID(ctx context.Context, id int64) (*entity.BorrowingBook, error)
	// UpdateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error
	// DeleteBorrowingBook(ctx context.Context, id int64) error
}

type BorrowingBookRepository interface {
	// GetAllBorrowingBooks(ctx context.Context) ([]*entity.BorrowingBook, error)
	CreateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error
	// GetBorrowingBookByID(ctx context.Context, id int64) (*entity.BorrowingBook, error)
	// UpdateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error
	// DeleteBorrowingBook(ctx context.Context, id int64) error
}

type BorrowingBookService struct {
	repository BorrowingBookRepository
}

func NewBorrowingBookService(repository BorrowingBookRepository) *BorrowingBookService {
	return &BorrowingBookService{
		repository: repository,
	}
}

func (s *BorrowingBookService) CreateBorrowingBook(ctx context.Context, borrowingBook *entity.BorrowingBook) error {
	// Lakukan penyimpanan borrowingBook melalui repository
	if err := s.repository.CreateBorrowingBook(ctx, borrowingBook); err != nil {
		return err // Mengembalikan error jika terjadi
	}
	// Jika berhasil, kembalikan nil
	return nil
}

// func (s *BorrowingBookService) GetBorrowingBookByID(ctx context.Context, id int64) (*entity.BorrowingBook, error) {
// 	return s.repository.GetBorrowingBookByID(ctx, id)
// }

// func (s *BorrowingBookService) UpdateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error {
// 	return s.repository.UpdateBorrowingBook(ctx, BorrowingBook)
// }

// func (s *BorrowingBookService) DeleteBorrowingBook(ctx context.Context, id int64) error {
// 	return s.repository.DeleteBorrowingBook(ctx, id)
// }
