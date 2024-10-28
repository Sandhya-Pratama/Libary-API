package service

import (
	"context"

	"github.com/Sandhya-Pratama/Libary-API/author-service/entity"
)

type AuthorUseCase interface {
	GetAllAuthors(ctx context.Context) ([]*entity.Author, error)
	CreateAuthor(ctx context.Context, Author *entity.Author) error
	UpdateAuthor(ctx context.Context, Author *entity.Author) error
	DeleteAuthor(ctx context.Context, id int64) error
}

type AuthorRepository interface {
	GetAllAuthors(ctx context.Context) ([]*entity.Author, error)
	CreateAuthor(ctx context.Context, Author *entity.Author) error
	UpdateAuthor(ctx context.Context, Author *entity.Author) error
	DeleteAuthor(ctx context.Context, id int64) error
}

type AuthorService struct {
	repository AuthorRepository
}

func NewAuthorService(repository AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repository,
	}
}

func (s *AuthorService) GetAllAuthors(ctx context.Context) ([]*entity.Author, error) {
	return s.repository.GetAllAuthors(ctx)
}

func (s *AuthorService) CreateAuthor(ctx context.Context, Author *entity.Author) error {
	return s.repository.CreateAuthor(ctx, Author)
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, Author *entity.Author) error {
	return s.repository.UpdateAuthor(ctx, Author)
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, id int64) error {
	return s.repository.DeleteAuthor(ctx, id)
}
