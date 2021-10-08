package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcproject/greet/greetpb"
	"io"
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
	c := greetpb.NewGreetServiceClient(conn)
	//doUranry(c)
	//doServerStreaming(c)
	//	doClientstreaming(c)
	//dobidirectional(c)

	doUranryWithDeadline(c, 5*time.Second)
	doUranryWithDeadline(c, 1*time.Second)

	fmt.Println("client completed")
}
func doUranry(c greetpb.GreetServiceClient) {
	fmt.Println("starting")
	req := &greetpb.GreetRequest{Greetiung: &greetpb.Greeting{
		FirstName: "Laksh",
		LastName:  "keswani",
	}}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling greet rpc %v", err)
	}
	log.Println("response from greet rpc %v", res.Result)

}
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("streaming client ok")
	req := &greetpb.GreetManytimesRequest{Greeting: &greetpb.Greeting{
		FirstName: "Laksh",
		LastName:  "Keswani",
	},
	}
	streamResult, err := c.GreetManytimes(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greetmanytimes  %v", err.Error())
	}
	for {
		msg, err := streamResult.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error recieving msg from stream")
		}
		log.Println("Response from GreetManyTimes : %v ", msg.GetResult())
	}
}
func doClientstreaming(c greetpb.GreetServiceClient) {
	fmt.Println("client streaming started")
	reqs := []*greetpb.LongGreetrequest{
		&greetpb.LongGreetrequest{Greeting: &greetpb.Greeting{
			FirstName: "laksh",
			LastName:  "keswani",
		}},
		&greetpb.LongGreetrequest{Greeting: &greetpb.Greeting{
			FirstName: "kantesh",
			LastName:  "keswani",
		}},
		&greetpb.LongGreetrequest{Greeting: &greetpb.Greeting{
			FirstName: "sonia",
			LastName:  "keswani",
		}},
		&greetpb.LongGreetrequest{Greeting: &greetpb.Greeting{
			FirstName: "ajnali",
			LastName:  "keswani",
		}}, &greetpb.LongGreetrequest{Greeting: &greetpb.Greeting{
			FirstName: "ajay",
			LastName:  "keswani",
		}},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal("error while doing client streaming %v", err)
	}
	for _, req := range reqs {
		fmt.Println("sernding requests ", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatal("erorr while sending msg %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while recieving ")
	}
	fmt.Println("response %v", res)
}
func dobidirectional(c greetpb.GreetServiceClient) {
	fmt.Println("bidirectional started")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal("erorr while creating the stream")
	}
	reqs := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{
			FirstName: "laksh",
			LastName:  "keswani",
		}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{
			FirstName: "kantesh",
			LastName:  "keswani",
		}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{
			FirstName: "sonia",
			LastName:  "keswani",
		}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{
			FirstName: "ajnali",
			LastName:  "keswani",
		}}, &greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{
			FirstName: "ajay",
			LastName:  "keswani",
		}},
	}
	waitc := make(chan struct{})
	go func() {
		//send the msgs
		for _, req := range reqs {
			fmt.Println("sending msg %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		//recieving stream
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Error while recieving msg")
				break
			}
			fmt.Println("recieved %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}

func doUranryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("starting uranry Greet with deadlines")
	req := &greetpb.GreetWithDeadlinesRequest{Greeting: &greetpb.Greeting{
		FirstName: "Laksh",
		LastName:  "keswani",
	}}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	res, err := c.GreetWithDeadlines(ctx, req)
	defer cancel()
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded %v", err)
			} else {
				log.Fatal("unexpeted erorr %v", statusErr)
			}
		} else {
			log.Fatal("error while calling greet rpc %v", err)
		}
		return
	}
	log.Println("response from greet with deadlines rpc %v", res.Result)

}
