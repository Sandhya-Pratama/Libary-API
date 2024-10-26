package handler

import (
	"net/http"
	"time"

	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/entity"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/service"
	"github.com/labstack/echo/v4"
)

type BorrowingBookHandler struct {
	BorrowingBookService service.BorrowingBookUseCase
}

func NewBorrowingBookHandler(BorrowingBookService service.BorrowingBookUseCase) *BorrowingBookHandler {
	return &BorrowingBookHandler{BorrowingBookService}
}

// CreateBorrowingBook membuat peminjaman baru
func (h *BorrowingBookHandler) CreateBorrowingBook(ctx echo.Context) error {
	// Mendefinisikan struct input untuk menerima data JSON
	var input struct {
		UserID     int64      `json:"user_id" validate:"required"`
		BookID     int64      `json:"book_id" validate:"required"`
		BorrowDate time.Time  `json:"borrow_date" validate:"required"` // Tanggal Peminjaman
		ReturnDate *time.Time `json:"return_date" validate:"required"` // Tanggal Pengembalian
	}

	// Binding input JSON ke struct
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Memvalidasi input setelah proses binding
	if err := ctx.Validate(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Validation failed", "details": validator.ValidatorErrors(err)})
	}

	// Membuat objek BorrowingBook baru dengan data yang sudah tervalidasi
	borrowingBook := entity.NewBorrowingBook(input.UserID, input.BookID, input.BorrowDate, input.ReturnDate)

	// Menggunakan service untuk membuat BorrowingBook baru di database
	if err := h.BorrowingBookService.CreateBorrowingBook(ctx.Request().Context(), borrowingBook); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	// Mengembalikan respon JSON saat berhasil
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message":       "Borrowing book created successfully",
		"borrowingBook": borrowingBook,
	})
}
