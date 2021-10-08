package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"grpcproject"
	"grpcproject/squareroot/sqarerootPB"
	"log"
	"math"
	"net"
)

type server struct {
	sqarerootPB.UnimplementedSquareRootServiceServer
}

func (*server) SquareRoot(ctx context.Context, req *sqarerootPB.SqaureRootRequest) (*sqarerootPB.SquareRootResponse, error) {
	fmt.Println("Average function working")
	num := req.GetNumber()
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("recieved a negative number :%v", num),
		)
	}
	result := math.Sqrt(float64(num))
	res := &sqarerootPB.SquareRootResponse{NumberRoot: result}

	return res, nil
}
func main() {
	fmt.Println("server started")
	lis, err := net.Listen(grpcproject.Tcp, grpcproject.Address)
	if err != nil {
		log.Fatal("error while listing to server")
	}
	s := grpc.NewServer()
	reflection.Register(s)

	sqarerootPB.RegisterSquareRootServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve %v", err)
	}

}
