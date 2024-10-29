package service

import (
	"context"
	"errors"

	pb "github.com/Sandhya-Pratama/Libary-API/book-service/book-service/common/proto"
	"github.com/Sandhya-Pratama/Libary-API/book-service/entity"
)

type BookGRPCServer struct {
	bookUseCase BookUseCase
	pb.UnimplementedBookServiceServer
}

func NewBookGRPCServer(bookUseCase BookUseCase) *BookGRPCServer {
	return &BookGRPCServer{bookUseCase: bookUseCase}
}

func (s *BookGRPCServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := entity.NewBook(req.Title, req.Isbn, req.AuthorId, req.CategoryId, req.Stock)
	err := s.bookUseCase.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBookResponse{
		Book: &pb.Book{
			Title:      book.Title,
			Isbn:       book.ISBN,
			AuthorId:   book.AuthorID,
			CategoryId: book.CategoryID,
			Stock:      book.Stock,
		},
	}, nil
}

func (s *BookGRPCServer) GetBookByID(ctx context.Context, req *pb.GetBookByIDRequest) (*pb.GetBookByIDResponse, error) {
	book, err := s.bookUseCase.GetBookByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	return &pb.GetBookByIDResponse{
		Book: &pb.Book{
			Title:      book.Title,
			Isbn:       book.ISBN,
			AuthorId:   book.AuthorID,
			CategoryId: book.CategoryID,
			Stock:      book.Stock,
		},
	}, nil
}

// Implementasi RPC UpdateBook
func (s *BookGRPCServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := entity.NewBook(req.Title, req.Isbn, req.AuthorId, req.CategoryId, req.Stock)
	book.ID = req.Id
	err := s.bookUseCase.UpdateBook(ctx, book)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBookResponse{Book: &pb.Book{
		Title:      book.Title,
		Isbn:       book.ISBN,
		AuthorId:   book.AuthorID,
		CategoryId: book.CategoryID,
		Stock:      book.Stock,
	}}, nil
}

// Implementasi RPC DeleteBook
func (s *BookGRPCServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	err := s.bookUseCase.DeleteBook(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookResponse{Message: "successfully deleted book"}, nil
}
