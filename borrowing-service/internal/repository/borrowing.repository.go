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

func (r *BorrowingBookRepository) CreateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error {
	err := r.db.WithContext(ctx).Create(&BorrowingBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BorrowingBookRepository) GetBorrowingBookByID(ctx context.Context, id int64) (*entity.BorrowingBook, error) {
	var BorrowingBook entity.BorrowingBook
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&BorrowingBook).Error
	if err != nil {
		return nil, err
	}
	return &BorrowingBook, nil
}

func (r *BorrowingBookRepository) UpdateBorrowingBook(ctx context.Context, BorrowingBook *entity.BorrowingBook) error {
	err := r.db.WithContext(ctx).Model(&entity.BorrowingBook{}).Where("id = ?", BorrowingBook.ID).Updates(&BorrowingBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BorrowingBookRepository) DeleteBorrowingBook(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.BorrowingBook{}, id).Error; err != nil {
		return err
	}
	return nil
}
