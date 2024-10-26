package entity

import "time"

type Book struct {
	ID         int64
	Title      string
	ISBN       int64
	AuthorID   int64
	CategoryID int64
	Stock      int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewBook(title string, isbn, stock, authorID, categoryID int64) *Book {
	return &Book{
		Title:      title,
		ISBN:       isbn,
		Stock:      stock,
		AuthorID:   authorID,
		CategoryID: categoryID,
		CreatedAt:  time.Now(),
	}
}

func UpdateBook(id int64, title string, isbn, stock, authorID, categoryID int64) *Book {
	return &Book{
		ID:         id,
		Title:      title,
		ISBN:       isbn,
		Stock:      stock,
		AuthorID:   authorID,
		CategoryID: categoryID,
		UpdatedAt:  time.Now(),
	}

}
