package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"grpcproject/blog/blogPB"
	"log"
)

func main() {
	fmt.Println("hello i am blog client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect %v", err)
	}
	defer conn.Close()
	c := blogPB.NewBlogServiceClient(conn)
	blog := &blogPB.Blog{
		AuthorId: "kantesh",
		Title:    "myfirst blog ",
		Content:  "content of the first blog",
	}
	fmt.Println("creating the blog")
	res, err := c.CreateBlog(context.Background(), &blogPB.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatal("error while creating blog %v", err)
	}
	blogid := res.GetBlog().GetId()
	fmt.Println("id =", blogid)
	fmt.Println("client completed response = %v", res)
	//updateBlog(c)
	DeleteBlog(c)
	fmt.Println("client completed", res)

}

func readBlog(c blogPB.BlogServiceClient) {
	fmt.Println("reading the blog")
	var id string = "6145f56ef192960c7df6303e"
	req := &blogPB.ReadBlogRequest{
		BlogId: id,
	}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		ErrStaus, ok := status.FromError(err)
		if ok {
			log.Fatal("error while read request %v %v", ErrStaus.Code(), ErrStaus.Message())
		}
	}
	fmt.Println("response from server is %v", res)
}
func DeleteBlog(c blogPB.BlogServiceClient) {
	fmt.Println("reading the blog")
	var id string = "6145f56ef192960c7df6303e"
	req := &blogPB.DeleteBlogRequest{
		BlogId: id,
	}
	res, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		ErrStaus, ok := status.FromError(err)
		if ok {
			log.Fatal("error while delete request %v %v", ErrStaus.Code(), ErrStaus.Message())
		}
	}
	fmt.Println("response from server is %v", res)
}

func updateBlog(c blogPB.BlogServiceClient) {
	fmt.Println("reading the blog")
	blog := &blogPB.Blog{
		Id:       "6145f56ef192960c7df6303e",
		AuthorId: "updated laksh",
		Title:    "updated myfirst blog",
		Content:  "updated content of the first blog",
	}

	fmt.Println("creating the blog")
	res, err := c.UpdateBlog(context.Background(), &blogPB.UpdateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatal("error while creating blog %v", err)
	}
	blogid := res.GetBlog().GetId()
	fmt.Println("id =", blogid)
	fmt.Println("client completed response = %v", res)

}
