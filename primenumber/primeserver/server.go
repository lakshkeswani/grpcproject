package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpcproject/primenumber/primepb"
	"log"
	"net"
	"time"
)

type server struct {
	primepb.UnimplementedPrimeServiceServer
}

func (*server) Prime(req *primepb.PrimeRequest, stream primepb.PrimeService_PrimeServer) error {
	var k int32 = 2
	num := req.GetNumber().GetNum()
	for num > 1 {
		if num%k == 0 {
			res := &primepb.PrimemultipleResponse{Result: k}
			stream.Send(res)
			num = num / k
			fmt.Printf("%d", k)
			time.Sleep(3000 * time.Millisecond)
		} else {
			k = k + 1
		}
	}
	return nil
}
func main() {
	fmt.Printf("server staring ")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("prime sever failed %v", err.Error())
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("can't listent the port %v", err)
	}
	fmt.Printf("server connected")

}
