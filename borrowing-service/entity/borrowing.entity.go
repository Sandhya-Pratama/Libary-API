package entity

import "time"

type BorrowingBook struct {
	ID         int64
	UserID     int64
	BookID     int64
	BorrowDate time.Time
	ReturnDate *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewBorrowingBook(userID, bookID int64, borrowDate time.Time, returnDate *time.Time) *BorrowingBook {
	return &BorrowingBook{
		UserID:     userID,
		BookID:     bookID,
		BorrowDate: borrowDate,
		ReturnDate: returnDate,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (b *BorrowingBook) UpdateBorrowingBook(userID, bookID int64, borrowDate, returnDate time.Time) {
	b.UserID = userID
	b.BookID = bookID
	b.BorrowDate = borrowDate
	b.ReturnDate = &returnDate
	b.UpdatedAt = time.Now()
}
