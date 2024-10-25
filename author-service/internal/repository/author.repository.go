package repository

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/author-service/entity"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) GetAllAuthors(ctx context.Context) ([]*entity.Author, error) {
	var Authors []*entity.Author
	err := r.db.WithContext(ctx).Find(&Authors).Error
	if err != nil {
		return nil, err
	}
	return Authors, nil
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, Author *entity.Author) error {
	err := r.db.WithContext(ctx).Create(&Author).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthorRepository) GetAuthorByID(ctx context.Context, id int64) (*entity.Author, error) {
	var Author entity.Author
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&Author).Error
	if err != nil {
		return nil, err
	}
	return &Author, nil
}

func (r *AuthorRepository) UpdateAuthor(ctx context.Context, Author *entity.Author) error {
	err := r.db.WithContext(ctx).Model(&entity.Author{}).Where("id = ?", Author.ID).Updates(&Author).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Author{}, id).Error; err != nil {
		return err
	}
	return nil
}
