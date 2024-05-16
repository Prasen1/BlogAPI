package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "BlogAPI/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nSelect an operation:")
		fmt.Println("1. Create Post")
		fmt.Println("2. Read Post")
		fmt.Println("3. Update Post")
		fmt.Println("4. Delete Post")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			createPost(client, reader)
		case "2":
			readPost(client, reader)
		case "3":
			updatePost(client, reader)
		case "4":
			deletePost(client, reader)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func createPost(client pb.BlogServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter content: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	fmt.Print("Enter author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter publication date (YYYY-MM-DD): ")
	pubDate, _ := reader.ReadString('\n')
	pubDate = strings.TrimSpace(pubDate)

	fmt.Print("Enter tags (comma-separated): ")
	tagsInput, _ := reader.ReadString('\n')
	tags := strings.Split(strings.TrimSpace(tagsInput), ",")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CreatePostRequest{
		Title:           title,
		Content:         content,
		Author:          author,
		PublicationDate: pubDate,
		Tags:            tags,
	}

	resp, err := client.CreatePost(ctx, req)
	if err != nil {
		log.Printf("CreatePost failed: %v", err)
		return
	}
	fmt.Printf("Created Post: %v\n", resp.Post)
}

func readPost(client pb.BlogServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter post ID: ")
	postID, _ := reader.ReadString('\n')
	postID = strings.TrimSpace(postID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.ReadPostRequest{PostId: postID}
	resp, err := client.ReadPost(ctx, req)
	if err != nil {
		log.Printf("ReadPost failed: %v", err)
		return
	}
	if resp.Error != "" {
		fmt.Printf("Error: %s\n", resp.Error)
	} else {
		fmt.Printf("Post: %v\n", resp.Post)
	}
}

func updatePost(client pb.BlogServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter post ID: ")
	postID, _ := reader.ReadString('\n')
	postID = strings.TrimSpace(postID)

	fmt.Print("Enter new title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter new content: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	fmt.Print("Enter new author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter new tags (comma-separated): ")
	tagsInput, _ := reader.ReadString('\n')
	tags := strings.Split(strings.TrimSpace(tagsInput), ",")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UpdatePostRequest{
		PostId:  postID,
		Title:   title,
		Content: content,
		Author:  author,
		Tags:    tags,
	}

	resp, err := client.UpdatePost(ctx, req)
	if err != nil {
		log.Printf("UpdatePost failed: %v", err)
		return
	}
	if resp.Error != "" {
		fmt.Printf("Error: %s\n", resp.Error)
	} else {
		fmt.Printf("Updated Post: %v\n", resp.Post)
	}
}

func deletePost(client pb.BlogServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter post ID: ")
	postID, _ := reader.ReadString('\n')
	postID = strings.TrimSpace(postID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.DeletePostRequest{PostId: postID}
	resp, err := client.DeletePost(ctx, req)
	if err != nil {
		log.Printf("DeletePost failed: %v", err)
		return
	}
	if resp.Error != "" {
		fmt.Printf("Error: %s\n", resp.Error)
	} else {
		fmt.Printf("Message: %s\n", resp.Message)
	}
}
