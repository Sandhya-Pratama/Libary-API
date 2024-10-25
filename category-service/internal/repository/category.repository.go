package repository

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/category-service/entity"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategorys(ctx context.Context) ([]*entity.Category, error) {
	var Categorys []*entity.Category
	err := r.db.WithContext(ctx).Find(&Categorys).Error
	if err != nil {
		return nil, err
	}
	return Categorys, nil
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, Category *entity.Category) error {
	err := r.db.WithContext(ctx).Create(&Category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int64) (*entity.Category, error) {
	var Category entity.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&Category).Error
	if err != nil {
		return nil, err
	}
	return &Category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, Category *entity.Category) error {
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", Category.ID).Updates(&Category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
