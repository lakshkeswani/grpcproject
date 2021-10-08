package controller

import "grpcproject/project/projectpb"

type Server struct {
	projectpb.UnimplementedUserServiceServer
}
