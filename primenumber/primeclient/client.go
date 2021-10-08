package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcproject/primenumber/primepb"
	"io"
	"log"
)

var port string = "localhost:50051"

func main() {
	fmt.Println("client starting")
	conn, err := doconnection()
	if err != nil {
		log.Fatal(" client could not connect %v from main ", err)
	}
	defer conn.Close()
	c := primepb.NewPrimeServiceClient(conn)
	doServerStreaming(c)

}
func doconnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect %v", err)
		return nil, err
	}
	return conn, nil
}
func doServerStreaming(c primepb.PrimeServiceClient) {
	fmt.Println()
	req := &primepb.PrimeRequest{Number: &primepb.Number{Num: 120}}
	resultStream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatal("eror while calling prime")
	}
	for {
		msg, err := resultStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error while recieve the message")
		}
		log.Printf("%v", msg.GetResult())
	}
}
