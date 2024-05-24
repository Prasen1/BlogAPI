package storage

import (
	pb "BlogAPI/pb"
	"errors"
	"sync"
)

type Storage interface {
	CreatePost(post *pb.Post) error
	ReadPost(postID string) (*pb.Post, error)
	UpdatePost(post *pb.Post) error
	DeletePost(postID string) error
}

type InMemoryStorage struct {
	posts map[string]*pb.Post
	mu    sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{posts: make(map[string]*pb.Post)}
}

func (s *InMemoryStorage) CreatePost(post *pb.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts[post.PostId] = post
	return nil
}

func (s *InMemoryStorage) ReadPost(postID string) (*pb.Post, error) {
	//s.mu.Lock()
	//defer s.mu.Unlock()
	post, exists := s.posts[postID]
	if !exists {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (s *InMemoryStorage) UpdatePost(post *pb.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts[post.PostId] = post
	return nil
}

func (s *InMemoryStorage) DeletePost(postID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.posts, postID)
	return nil
}
