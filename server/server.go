package main

import (
	pb "BlogAPI/pb"
	storage "BlogAPI/storage"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBlogServiceServer
	storage storage.Storage
}

func (s *server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	postID := fmt.Sprintf("post_%d", time.Now().UnixNano())
	post := &pb.Post{
		PostId:          postID,
		Title:           req.GetTitle(),
		Content:         req.GetContent(),
		Author:          req.GetAuthor(),
		PublicationDate: req.GetPublicationDate(),
		Tags:            req.GetTags(),
	}

	if err := s.storage.CreatePost(post); err != nil {
		return nil, err
	}
	log.Printf("Created post with Id: %s", post.PostId)
	return &pb.PostResponse{Post: post}, nil
}

func (s *server) ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.PostResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	post, err := s.storage.ReadPost(req.GetPostId())
	if err != nil {
		return &pb.PostResponse{Error: err.Error()}, err
	}
	log.Printf("Returned read request for post id: %s", post.PostId)
	return &pb.PostResponse{Post: post}, nil
}

func (s *server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	post := &pb.Post{
		PostId:  req.GetPostId(),
		Title:   req.GetTitle(),
		Content: req.GetContent(),
		Author:  req.GetAuthor(),
		Tags:    req.GetTags(),
	}

	if err := s.storage.UpdatePost(post); err != nil {
		return &pb.PostResponse{Error: err.Error()}, err
	}
	log.Printf("Updated post with Id: %s", post.PostId)
	return &pb.PostResponse{Post: post}, nil
}

func (s *server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	log.Printf("Got delete request for post with Id: %s", req.GetPostId())
	if err := s.storage.DeletePost(req.GetPostId()); err != nil {
		return &pb.DeletePostResponse{Error: err.Error()}, err
	}
	return &pb.DeletePostResponse{Message: "Post deleted successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	storage := storage.NewInMemoryStorage()
	pb.RegisterBlogServiceServer(s, &server{storage: storage})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
