package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcproject/squareroot/sqarerootPB"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect %v", err)
	}
	defer conn.Close()
	c := sqarerootPB.NewSquareRootServiceClient(conn)
	doUranry(c)
}
func doUranry(c sqarerootPB.SquareRootServiceClient) {
	fmt.Println("starting")
	req := &sqarerootPB.SqaureRootRequest{Number: -2}

	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			fmt.Println(resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("we probably sent a negative number")
			}
		}
		log.Fatal("error while calling greet rpc %v", err)
	}
	log.Println("response from Sqaureroot rpc %v", res.NumberRoot)

}
