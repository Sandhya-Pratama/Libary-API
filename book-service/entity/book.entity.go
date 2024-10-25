package entity

import "time"

type Book struct {
	ID         int64
	Title      string
	Author     string
	Year       int
	Stock      int
	AuthorID   int64
	CategoryID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewBook(title, author string, year, stock int, authorID, categoryID int64) *Book {
	return &Book{
		Title:      title,
		Author:     author,
		Year:       year,
		Stock:      stock,
		AuthorID:   authorID,
		CategoryID: categoryID,
		CreatedAt:  time.Now(),
	}
}

func UpdateBook(id int64, title, author string, year, stock int, authorID, categoryID int64) *Book {
	return &Book{
		ID:         id,
		Title:      title,
		Author:     author,
		Year:       year,
		Stock:      stock,
		AuthorID:   authorID,
		CategoryID: categoryID,
		UpdatedAt:  time.Now(),
	}

}
