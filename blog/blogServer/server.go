package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcproject"
	"grpcproject/blog/blogPB"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var Collection *mongo.Collection

type server struct {
	blogPB.UnimplementedBlogServiceServer
}
type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, req *blogPB.CreateBlogRequest) (*blogPB.CreateBlogResponse, error) {

	blog := req.GetBlog()
	fmt.Println("recieved request %v ", blog)
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}
	res, err := Collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal("error while inserting to data base")
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal erorr %v", err))
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("error while parsing to data objid")
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal erorr "))
	}
	fmt.Println("sucessfully inserted")
	return &blogPB.CreateBlogResponse{Blog: &blogPB.Blog{
		Id:       oid.Hex(),
		AuthorId: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}}, nil
}
func (*server) UpdateBlog(ctx context.Context, req *blogPB.UpdateBlogRequest) (*blogPB.UpdateBlogResponse, error) {
	blog := req.GetBlog()
	//md,ok :=metadata.FromIncomingContext(ctx)
	fmt.Println("update blog request recieved")
	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		log.Fatal("error while parsing ")
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("can not parse id"))
	}
	data := &blogItem{}
	data.ID = oid
	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()
	filter := bson.M{"_id": oid}
	res, err := Collection.ReplaceOne(context.Background(), filter, data)
	if err != nil {
		log.Fatal("erorr while updating")
	}
	fmt.Println("result = %v", res)
	return &blogPB.UpdateBlogResponse{Blog: dataToBlog(data)}, nil
}
func (*server) ReadBlog(ctx context.Context, req *blogPB.ReadBlogRequest) (*blogPB.ReadBlogResponse, error) {
	fmt.Println("read blog request")
	id := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("error while parsing ")
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("can not parse id"))
	}
	data := &blogItem{}
	filter := bson.M{"_id": oid}
	res := Collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}
	return &blogPB.ReadBlogResponse{Blog: dataToBlog(data)}, nil
}
func (*server) DeleteBlog(ctx context.Context, req *blogPB.DeleteBlogRequest) (*blogPB.DeleteBlogResponse, error) {
	fmt.Println("delete blog request")
	id := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("error while parsing ")
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("can not parse id"))
	}
	filter := bson.M{"_id": oid}
	res, err := Collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}
	fmt.Println("delete Reaponse %v", res)
	return &blogPB.DeleteBlogResponse{Result: "sucessfull"}, nil
}
func (*server) ListBlog(req *blogPB.ListBlogRequest, stream blogPB.BlogService_ListBlogServer) error {
	curs, err := Collection.Find(context.Background(), nil)

	defer curs.Close(context.Background())
	if err != nil {
		log.Println("error while parsing ")
		return status.Errorf(
			codes.Internal, fmt.Sprintf("internal server error"))
	}

	for curs.Next(context.Background()) {
		data := &blogItem{}
		curs.Decode(data)
	}

	return err
}

func dataToBlog(data *blogItem) *blogPB.Blog {
	return &blogPB.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
func main() {
	go func() {
		mux := runtime.NewServeMux()
		blogPB.RegisterBlogServiceHandlerServer(context.Background(), mux, &server{})
		log.Fatal(http.ListenAndServe(grpcproject.Restlocalhost, mux))
	}()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Hello i am server")
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(grpcproject.MongodbAddress))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database("mydb").Collection("blog")
	lis, err := net.Listen("tcp", grpcproject.Address)
	if err != nil {
		log.Fatal("Failed to  %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogPB.RegisterBlogServiceServer(s, &server{})

	fmt.Println("every thing working fine")
	go func() {
		fmt.Println("starting server")
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serv %v", err)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	fmt.Println("stopping the server ")
	s.Stop()
	fmt.Println("stopping the lisner")
	lis.Close()
	client.Disconnect(context.TODO())
	fmt.Println("mongodb disconnected")
	fmt.Println("end of program")

}
