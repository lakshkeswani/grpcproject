package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcproject/average/averagepb"
	"log"
	"time"
)

func main() {
	fmt.Println("hello i am client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect %v", err)
	}
	defer conn.Close()
	c := averagepb.NewAverageServiceClient(conn)
	doclientStreaming(c)
	fmt.Println("client closing connnection")
}
func doclientStreaming(c averagepb.AverageServiceClient) {
	fmt.Println("client streaming started")
	stream, err := c.AverageSer(context.Background())
	if err != nil {
		log.Fatal("error while getting stream")
	}
	reqs := []*averagepb.AverageRequestt{
		&averagepb.AverageRequestt{Average: &averagepb.Average{Num: float64(5)}},
		&averagepb.AverageRequestt{Average: &averagepb.Average{Num: float64(1)}},
		&averagepb.AverageRequestt{Average: &averagepb.Average{Num: float64(2)}},
		&averagepb.AverageRequestt{Average: &averagepb.Average{Num: float64(3)}},
		&averagepb.AverageRequestt{Average: &averagepb.Average{Num: float64(4)}},
	}
	reqs = append(reqs, &averagepb.AverageRequestt{Average: &averagepb.Average{Num: float64(5)}})

	for _, req := range reqs {
		err = stream.Send(req)
		if err != nil {
			log.Fatal("error while sending msg")
		}
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while recieving ")
	}
	fmt.Println("response %v", res)

}
