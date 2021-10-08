package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcproject/greet/greetpb"
	"log"
	"net"
)

type sumserver struct {
	greetpb.UnimplementedDoSumServiceServer
}

func (s *sumserver) Sum(c context.Context, req *greetpb.DoSumRequest) (*greetpb.DoSumResponse, error) {
	fmt.Println("sum is called")
	a := req.GetDosum().GetA()
	b := req.GetDosum().GetB()
	sum := a + b
	res := &greetpb.DoSumResponse{Sum: sum}
	return res, nil
}

func smain() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("sum server failed %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterDoSumServiceServer(s, &sumserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serv %v", err)
	}
}
