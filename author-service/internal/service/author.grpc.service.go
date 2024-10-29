package service

import (
	"context"

	pb "github.com/Sandhya-Pratama/Libary-API/author-service/author-service/common/proto"
	"github.com/Sandhya-Pratama/Libary-API/author-service/entity"
)

type AuthorGRPCServer struct {
	authorUseCase AuthorUseCase
	pb.UnimplementedAuthorServiceServer
}

func NewAuthorGRPCServer(authorUseCase AuthorUseCase) *AuthorGRPCServer {
	return &AuthorGRPCServer{authorUseCase: authorUseCase}
}

// Implementasi RPC GetAllAuthors
func (s *AuthorGRPCServer) GetAllAuthors(ctx context.Context, req *pb.Empty) (*pb.AuthorListResponse, error) {
	authors, err := s.authorUseCase.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}

	var authorList []*pb.AuthorResponse
	for _, author := range authors {
		authorList = append(authorList, &pb.AuthorResponse{
			Id:   author.ID,
			Name: author.Name,
			Bio:  author.Bio,
		})
	}

	return &pb.AuthorListResponse{Authors: authorList}, nil
}

// Implementasi RPC CreateAuthor
func (s *AuthorGRPCServer) CreateAuthor(ctx context.Context, req *pb.AuthorRequest) (*pb.AuthorResponse, error) {
	author := entity.NewAuthor(req.Name, req.Bio)
	err := s.authorUseCase.CreateAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	return &pb.AuthorResponse{
		Id:   author.ID,
		Name: author.Name,
		Bio:  author.Bio,
	}, nil
}

// Implementasi RPC UpdateAuthor
func (s *AuthorGRPCServer) UpdateAuthor(ctx context.Context, req *pb.UpdateAuthorRequest) (*pb.AuthorResponse, error) {
	author := entity.NewAuthor(req.Name, req.Bio)
	author.ID = req.Id
	err := s.authorUseCase.UpdateAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	return &pb.AuthorResponse{
		Id:   author.ID,
		Name: author.Name,
		Bio:  author.Bio,
	}, nil
}

// Implementasi RPC DeleteAuthor
func (s *AuthorGRPCServer) DeleteAuthor(ctx context.Context, req *pb.DeleteAuthorRequest) (*pb.Empty, error) {
	err := s.authorUseCase.DeleteAuthor(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
