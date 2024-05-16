package main

import (
	"context"
	"log"
	"time"

	pb "BlogAPI/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// Create a new post
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	post, err := c.CreatePost(ctx, &pb.CreatePostRequest{
		Title:           "My First Blog Post",
		Content:         "This is the content of the first blog post.",
		Author:          "Author1",
		PublicationDate: "2024-05-15",
		Tags:            []string{"golang", "grpc", "tutorial"},
	})
	if err != nil {
		log.Fatalf("Could not create post: %v", err)
	}
	log.Printf("Post created: %v", post.GetPost())

	// Read the created post
	post, err = c.ReadPost(ctx, &pb.ReadPostRequest{PostId: post.GetPost().GetPostId()})
	if err != nil {
		log.Fatalf("Could not read post: %v", err)
	}
	log.Printf("Post read: %v", post.GetPost())

	// Update the post
	post, err = c.UpdatePost(ctx, &pb.UpdatePostRequest{
		PostId:  post.GetPost().GetPostId(),
		Title:   "Updated Blog Post Title",
		Content: "This is the updated content of the blog post.",
		Author:  "Author1",
		Tags:    []string{"golang", "grpc", "updated"},
	})
	if err != nil {
		log.Fatalf("Could not update post: %v", err)
	}
	log.Printf("Post updated: %v", post.GetPost())

	// Delete the post
	deleteRes, err := c.DeletePost(ctx, &pb.DeletePostRequest{PostId: post.GetPost().GetPostId()})
	if err != nil {
		log.Fatalf("Could not delete post: %v", err)
	}
	log.Printf("Delete response: %v", deleteRes.GetMessage())
}
