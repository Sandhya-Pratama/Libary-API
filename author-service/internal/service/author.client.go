package service

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Sandhya-Pratama/Libary-API/author-service/author-service/common/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthorClient struct {
	client pb.AuthorServiceClient
}

func NewAuthorClient() *AuthorClient {
	authorServiceURL := os.Getenv("AUTHOR_SERVICE_URL")
	conn, err := grpc.Dial(authorServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Author Service: %v", err)
	}
	client := pb.NewAuthorServiceClient(conn)
	return &AuthorClient{client: client}
}

func (c *AuthorClient) GetAllAuthors(ctx context.Context) (*pb.AuthorListResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return c.client.GetAllAuthors(ctx, &pb.Empty{})
}
