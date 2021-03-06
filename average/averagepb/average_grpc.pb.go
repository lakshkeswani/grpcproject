// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package averagepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AverageServiceClient is the client API for AverageService controller.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AverageServiceClient interface {
	AverageSer(ctx context.Context, opts ...grpc.CallOption) (AverageService_AverageSerClient, error)
}

type averageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAverageServiceClient(cc grpc.ClientConnInterface) AverageServiceClient {
	return &averageServiceClient{cc}
}

func (c *averageServiceClient) AverageSer(ctx context.Context, opts ...grpc.CallOption) (AverageService_AverageSerClient, error) {
	stream, err := c.cc.NewStream(ctx, &AverageService_ServiceDesc.Streams[0], "/average.AverageService/AverageSer", opts...)
	if err != nil {
		return nil, err
	}
	x := &averageServiceAverageSerClient{stream}
	return x, nil
}

type AverageService_AverageSerClient interface {
	Send(*AverageRequestt) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type averageServiceAverageSerClient struct {
	grpc.ClientStream
}

func (x *averageServiceAverageSerClient) Send(m *AverageRequestt) error {
	return x.ClientStream.SendMsg(m)
}

func (x *averageServiceAverageSerClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AverageServiceServer is the server API for AverageService controller.
// All implementations must embed UnimplementedAverageServiceServer
// for forward compatibility
type AverageServiceServer interface {
	AverageSer(AverageService_AverageSerServer) error
	mustEmbedUnimplementedAverageServiceServer()
}

// UnimplementedAverageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAverageServiceServer struct {
}

func (UnimplementedAverageServiceServer) AverageSer(AverageService_AverageSerServer) error {
	return status.Errorf(codes.Unimplemented, "method AverageSer not implemented")
}
func (UnimplementedAverageServiceServer) mustEmbedUnimplementedAverageServiceServer() {}

// UnsafeAverageServiceServer may be embedded to opt out of forward compatibility for this controller.
// Use of this interface is not recommended, as added methods to AverageServiceServer will
// result in compilation errors.
type UnsafeAverageServiceServer interface {
	mustEmbedUnimplementedAverageServiceServer()
}

func RegisterAverageServiceServer(s grpc.ServiceRegistrar, srv AverageServiceServer) {
	s.RegisterService(&AverageService_ServiceDesc, srv)
}

func _AverageService_AverageSer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AverageServiceServer).AverageSer(&averageServiceAverageSerServer{stream})
}

type AverageService_AverageSerServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequestt, error)
	grpc.ServerStream
}

type averageServiceAverageSerServer struct {
	grpc.ServerStream
}

func (x *averageServiceAverageSerServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *averageServiceAverageSerServer) Recv() (*AverageRequestt, error) {
	m := new(AverageRequestt)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AverageService_ServiceDesc is the grpc.ServiceDesc for AverageService controller.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AverageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "average.AverageService",
	HandlerType: (*AverageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AverageSer",
			Handler:       _AverageService_AverageSer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "average/averagepb/average.proto",
}
