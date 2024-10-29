package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Sandhya-Pratama/Libary-API/borrowing-service/borrowing-service/common/proto"
	"github.com/Sandhya-Pratama/Libary-API/borrowing-service/entity"
	"gorm.io/gorm"
)

type BorrowingService struct {
	db *gorm.DB
	pb.UnimplementedBorrowingServiceServer
}

func NewBorrowingService(db *gorm.DB) *BorrowingService {
	return &BorrowingService{db: db}
}

// BorrowBook handles book borrowing
func (s *BorrowingService) BorrowBook(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowResponse, error) {
	borrowing := entity.BorrowingBook{
		UserID:     int64(req.UserId), // Konversi eksplisit dari int32 ke int64
		BookID:     int64(req.BookId), // Konversi eksplisit dari int32 ke int64
		BorrowDate: time.Now(),
	}

	// Save borrowing to the database
	if err := s.db.Create(&borrowing).Error; err != nil {
		return nil, fmt.Errorf("failed to create borrowing record: %w", err)
	}

	return &pb.BorrowResponse{
		BorrowingId: int32(borrowing.ID),
		Message:     "Book borrowed successfully",
	}, nil
}

// ReturnBook handles book returning
func (s *BorrowingService) ReturnBook(ctx context.Context, req *pb.ReturnRequest) (*pb.ReturnResponse, error) {
	var borrowing entity.BorrowingBook

	// Find the borrowing record
	if err := s.db.First(&borrowing, req.BorrowingId).Error; err != nil {
		return nil, fmt.Errorf("borrowing record not found: %w", err)
	}

	// Update the return date
	now := time.Now()           // Simpan waktu sekarang ke dalam variabel
	borrowing.ReturnDate = &now // Assign pointer ke waktu sekarang
	if err := s.db.Save(&borrowing).Error; err != nil {
		return nil, fmt.Errorf("failed to update return date: %w", err)
	}

	return &pb.ReturnResponse{
		BorrowingId: int32(borrowing.ID),
		Message:     "Book returned successfully",
	}, nil
}

// GetBorrowingsByUser fetches all borrowings by a user
func (s *BorrowingService) GetBorrowingsByUser(ctx context.Context, req *pb.UserBorrowingsRequest) (*pb.BorrowingListResponse, error) {
	var borrowings []entity.BorrowingBook
	if err := s.db.Where("user_id = ?", req.UserId).Find(&borrowings).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve borrowings: %w", err)
	}

	var borrowingList []*pb.BorrowingResponse
	for _, b := range borrowings {
		borrowingList = append(borrowingList, &pb.BorrowingResponse{
			BorrowingId: int32(b.ID),
			UserId:      int32(b.UserID),
			BookId:      int32(b.BookID),
			BorrowDate:  b.BorrowDate.Format("2006-01-02"),
			ReturnDate:  b.ReturnDate.Format("2006-01-02"), // Konversi time.Time ke Timestamp
		})
	}

	return &pb.BorrowingListResponse{Borrowings: borrowingList}, nil
}
