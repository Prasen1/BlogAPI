package main

import (
	"context"
	"testing"
	"time"

	pb "BlogAPI/pb"
	"BlogAPI/storage"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	storage := storage.NewInMemoryStorage()
	srv := &server{storage: storage}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CreatePostRequest{
		Title:           "Test Title",
		Content:         "Test Content",
		Author:          "Test Author",
		PublicationDate: "2024-05-15",
		Tags:            []string{"test", "post"},
	}

	resp, err := srv.CreatePost(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Post)
	assert.Equal(t, req.Title, resp.Post.Title)
	assert.Equal(t, req.Content, resp.Post.Content)
	assert.Equal(t, req.Author, resp.Post.Author)
	assert.Equal(t, req.PublicationDate, resp.Post.PublicationDate)
	assert.ElementsMatch(t, req.Tags, resp.Post.Tags)
}

func TestReadPost(t *testing.T) {
	storage := storage.NewInMemoryStorage()
	srv := &server{storage: storage}

	post := &pb.Post{
		PostId:          "post_1",
		Title:           "Test Title",
		Content:         "Test Content",
		Author:          "Test Author",
		PublicationDate: "2024-05-15",
		Tags:            []string{"test", "post"},
	}
	storage.CreatePost(post)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.ReadPostRequest{PostId: "post_1"}

	resp, err := srv.ReadPost(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Post)
	assert.Equal(t, post.PostId, resp.Post.PostId)
	assert.Equal(t, post.Title, resp.Post.Title)
	assert.Equal(t, post.Content, resp.Post.Content)
	assert.Equal(t, post.Author, resp.Post.Author)
	assert.Equal(t, post.PublicationDate, resp.Post.PublicationDate)
	assert.ElementsMatch(t, post.Tags, resp.Post.Tags)
}

func TestUpdatePost(t *testing.T) {
	storage := storage.NewInMemoryStorage()
	srv := &server{storage: storage}

	post := &pb.Post{
		PostId:          "post_1",
		Title:           "Test Title",
		Content:         "Test Content",
		Author:          "Test Author",
		PublicationDate: "2024-05-15",
		Tags:            []string{"test", "post"},
	}
	storage.CreatePost(post)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UpdatePostRequest{
		PostId:  "post_1",
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  "Updated Author",
		Tags:    []string{"updated", "post"},
	}

	resp, err := srv.UpdatePost(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Post)
	assert.Equal(t, req.PostId, resp.Post.PostId)
	assert.Equal(t, req.Title, resp.Post.Title)
	assert.Equal(t, req.Content, resp.Post.Content)
	assert.Equal(t, req.Author, resp.Post.Author)
	assert.ElementsMatch(t, req.Tags, resp.Post.Tags)
}

func TestDeletePost(t *testing.T) {
	storage := storage.NewInMemoryStorage()
	srv := &server{storage: storage}

	post := &pb.Post{
		PostId:          "post_1",
		Title:           "Test Title",
		Content:         "Test Content",
		Author:          "Test Author",
		PublicationDate: "2024-05-15",
		Tags:            []string{"test", "post"},
	}
	storage.CreatePost(post)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.DeletePostRequest{PostId: "post_1"}

	resp, err := srv.DeletePost(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, "Post deleted successfully", resp.Message)
}
