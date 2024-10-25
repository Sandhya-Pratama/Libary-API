package repository

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/book-service/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]*entity.Book, error) {
	var books []*entity.Book
	err := r.db.WithContext(ctx).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) CreateBook(ctx context.Context, book *entity.Book) error {
	err := r.db.WithContext(ctx).Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int64) (*entity.Book, error) {
	var book entity.Book
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, book *entity.Book) error {
	err := r.db.WithContext(ctx).Model(&entity.Book{}).Where("id = ?", book.ID).Updates(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
