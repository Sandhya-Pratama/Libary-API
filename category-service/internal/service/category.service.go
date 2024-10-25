package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/category-service/entity"
)

type CategoryUseCase interface {
	GetAllCategorys(ctx context.Context) ([]*entity.Category, error)
	CreateCategory(ctx context.Context, Category *entity.Category) error
	UpdateCategory(ctx context.Context, Category *entity.Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryRepository interface {
	GetAllCategorys(ctx context.Context) ([]*entity.Category, error)
	CreateCategory(ctx context.Context, Category *entity.Category) error
	UpdateCategory(ctx context.Context, Category *entity.Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryService struct {
	repository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) GetAllCategorys(ctx context.Context) ([]*entity.Category, error) {
	return s.repository.GetAllCategorys(ctx)
}

func (s *CategoryService) CreateCategory(ctx context.Context, Category *entity.Category) error {
	return s.repository.CreateCategory(ctx, Category)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, Category *entity.Category) error {
	return s.repository.UpdateCategory(ctx, Category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int64) error {
	return s.repository.DeleteCategory(ctx, id)
}
