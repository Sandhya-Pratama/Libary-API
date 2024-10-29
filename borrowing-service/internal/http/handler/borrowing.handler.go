package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/entity"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/internal/service"
	"github.com/labstack/echo/v4"
)

type BorrowingBookHandler struct {
	BorrowingBookService service.BorrowingBookUseCase
}

// NewBorrowingBookHandler creates a new BorrowingBookHandler with the provided BorrowingBookUseCase.
func NewBorrowingBookHandler(BorrowingBookService service.BorrowingBookUseCase) *BorrowingBookHandler {
	return &BorrowingBookHandler{BorrowingBookService}
}

// BorrowBorrowingBook handles the HTTP request to borrow a BorrowingBook.
func (h *BorrowingBookHandler) BorrowBook(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.FormValue("user_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
	}
	bookID, err := strconv.Atoi(ctx.FormValue("BorrowingBook_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid BorrowingBook_id"})
	}

	// Membuat instance BorrowingBook dan mengisi data
	borrowingBook := &entity.BorrowingBook{
		UserID: int64(userID),
		BookID: int64(bookID),
	}

	// Memanggil service untuk meminjam buku
	if err := h.BorrowingBookService.BorrowBook(context.Background(), borrowingBook); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to borrow book"})
	}

	// Membuat respons jika peminjaman berhasil
	res := map[string]interface{}{
		"message":     "Book borrowed successfully",
		"borrowingId": borrowingBook.ID, // asumsikan ID dihasilkan oleh database saat peminjaman
	}
	return ctx.JSON(http.StatusOK, res)
}

// ReturnBorrowingBook handles the HTTP request to return a BorrowingBook.
func (h *BorrowingBookHandler) ReturnBook(ctx echo.Context) error {
	returningID, err := strconv.Atoi(ctx.FormValue("returning_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid returning_id"})
	}

	// Membuat instance BorrowingBook dan mengisi data
	borrowingBook := &entity.BorrowingBook{
		ID: int64(returningID),
	}

	// Memanggil service untuk mengembalikan buku
	if err := h.BorrowingBookService.ReturnBook(context.Background(), borrowingBook); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to return book"})
	}

	// Membuat respons jika pengembalian buku berhasil
	res := map[string]interface{}{
		"message": "Book returned successfully",
	}
	return ctx.JSON(http.StatusOK, res)
}

// GetBorrowingsByUser handles the HTTP request to get all borrowings for a user.
func (h *BorrowingBookHandler) GetBorrowingsByUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
	}

	// Call the service to get borrowings by user
	res, err := h.BorrowingBookService.GetBorrowingsByUser(context.Background(), int64(userID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch borrowings"})
	}

	return ctx.JSON(http.StatusOK, res)
}
