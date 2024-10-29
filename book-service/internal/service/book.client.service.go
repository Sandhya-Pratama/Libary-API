package service

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Sandhya-Pratama/Libary-API/book-service/book-service/common/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookClient struct {
	client pb.BookServiceClient
}

func NewBookClient() *BookClient {
	bookServiceURL := os.Getenv("BOOK_SERVICE_URL")

	// Membuat koneksi dengan gRPC server dengan timeout konteks
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		bookServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Gunakan transport credentials tanpa enkripsi
		grpc.WithBlock(), // Menunggu koneksi hingga berhasil
	)
	if err != nil {
		log.Fatalf("Failed to connect to Book Service: %v", err)
	}

	client := pb.NewBookServiceClient(conn)
	return &BookClient{client: client}
}

// Method to check book availability
func (c *BookClient) CheckBookAvailability(ctx context.Context, bookID int64) (bool, error) {
	req := &pb.GetBookByIDRequest{Id: bookID}
	res, err := c.client.GetBookByID(ctx, req)
	if err != nil {
		return false, err
	}
	return res.Book.Stock > 0, nil
}
