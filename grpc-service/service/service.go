package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kaznowski-brothers/blog-backend/grpc-service/db"
	pb "github.com/kaznowski-brothers/blog-backend/proto"
	"log"
)

type BlogService struct {
	db *db.DB
}

func New() *BlogService {
	s := &BlogService{
		db: db.New(),
	}
	return s
}

func (s *BlogService) Ping(ctx context.Context, req *empty.Empty) (*pb.Pong, error) {
	log.Printf("Ping")
	numberOfQueries := s.db.DummyCall()
	return &pb.Pong{Answer: fmt.Sprintf("Pong answer %v", numberOfQueries)}, nil
}
