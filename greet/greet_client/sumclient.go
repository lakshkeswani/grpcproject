package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcproject/greet/greetpb"
	"log"
)

func cmain() {
	fmt.Println("sum client started")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client not able to connect with server")
	}
	defer conn.Close()
	c := greetpb.NewDoSumServiceClient(conn)
	dosum(c)
}
func dosum(c greetpb.DoSumServiceClient) {
	req := &greetpb.DoSumRequest{Dosum: &greetpb.DoSum{
		A: 12,
		B: 23,
	}}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling dosum rpc ")
	}
	log.Println("response from sum rpc %v", res.Sum)

}
