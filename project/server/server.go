package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/controller"
	"grpcproject/project/server/db"
	"net"
	"net/http"
)

func startRestServer() {
	mux := runtime.NewServeMux()
	projectpb.RegisterUserServiceHandlerServer(context.Background(), mux, &controller.Server{})

	resources.Log.Fatal(http.ListenAndServe(resources.Restlocalhost, mux))
}
func startGrpcServer() {

	log.Println("server started")
	lis, err := net.Listen(resources.Tcp, resources.Address)
	if err != nil {
		resources.Log.Fatal("unable to listen to server", err)
	}
	s := grpc.NewServer()
	projectpb.RegisterUserServiceServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		resources.Log.Fatal("server started %v", err)
	}
}

func main() {
	resources.Log.Initalizer()
	db.InitialMigration()
	go startRestServer()
	startGrpcServer()
}
