package repository

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/entity"
	"gorm.io/gorm"
)

type BorrowingBookRepository struct {
	db *gorm.DB
}

func NewBorrowingBookRepository(db *gorm.DB) *BorrowingBookRepository {
	return &BorrowingBookRepository{db: db}
}

func (r *BorrowingBookRepository) BorrowBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error {
	err := r.db.WithContext(ctx).Create(&BorrowingBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BorrowingBookRepository) ReturnBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error {
	err := r.db.WithContext(ctx).Model(&entity.BorrowingBook{}).Where("id = ?", BorrowingBook.ID).Updates(&BorrowingBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BorrowingBookRepository) GetBorrowingsByUser(ctx context.Context, userID int64) ([]*entity.BorrowingBook, error) {
	var borrowings []*entity.BorrowingBook
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&borrowings).Error
	if err != nil {
		return nil, err
	}
	return borrowings, nil
}
