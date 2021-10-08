package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpcproject/average/averagepb"
	"io"
	"log"
	"net"
)

var tcp string = "tcp"
var address string = "0.0.0.0:50051"

type server struct {
	averagepb.UnimplementedAverageServiceServer
}

func (*server) AverageSer(stream averagepb.AverageService_AverageSerServer) error {
	fmt.Println("Average function working")
	var sum float64 = 0
	for i := 1; ; i++ {
		req, err := stream.Recv()
		sum += req.GetAverage().GetNum()
		if err == io.EOF {
			result := sum / float64(i)
			return stream.SendAndClose(&averagepb.AverageResponse{Result: result})
		}
	}
}
func main() {
	fmt.Println("server started")
	lis, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatal("unable to listen to server")
	}
	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serv %v", err)
	}
}
