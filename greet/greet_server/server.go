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
	"net"
	"strconv"
	"time"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GeetResponse, error) {
	fmt.Println("greet called %v", req)
	firstname := req.GetGreetiung().FirstName
	result := "hello" + firstname
	res := &greetpb.GeetResponse{Result: result}

	return res, nil
}
func (*server) GreetWithDeadlines(ctx context.Context, req *greetpb.GreetWithDeadlinesRequest) (*greetpb.GreetWithDeadlinesResponse, error) {
	fmt.Println("greet with deadline called %v", req)
	firstname := req.GetGreeting().FirstName
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("the client has cancelled the request!")
			return nil, status.Errorf(codes.Canceled, "client has canceled the request from server")
		}
		time.Sleep(time.Second)

	}
	result := "hello" + firstname
	res := &greetpb.GreetWithDeadlinesResponse{Result: result}

	return res, nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("started")
	result := "hello"
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{Result: result})
		}
		if err != nil {
			log.Fatal("Error while reading client stream %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result = result + "  " + firstName + "! "
	}

}
func (*server) GreetManytimes(req *greetpb.GreetManytimesRequest, stream greetpb.GreetService_GreetManytimesServer) error {
	firstname := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "hello" + firstname + "number" + strconv.Itoa(i)
		res := &greetpb.GreetManytimesResponse{Result: result}
		stream.Send(res)
		fmt.Println("result is %v", res.GetResult())
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("started")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("error while reading the client streaming", err)
			return err
		}
		firstname := req.GetGreeting().GetFirstName()

		result := "hello " + firstname
		err = stream.Send(&greetpb.GreetEveryoneResponse{Result: result})
		if err != nil {
			log.Fatal("err while  sending response", err)
			return err
		}
	}
}

func main() {
	fmt.Println("Hello i am server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to  %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serv %v", err)
	}
}
