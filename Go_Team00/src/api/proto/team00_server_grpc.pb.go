// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: team00_server.proto

package serv

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

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	DeviceInfo(ctx context.Context, in *RequestEmpty, opts ...grpc.CallOption) (DeviceService_DeviceInfoClient, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) DeviceInfo(ctx context.Context, in *RequestEmpty, opts ...grpc.CallOption) (DeviceService_DeviceInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &DeviceService_ServiceDesc.Streams[0], "/pb.DeviceService/DeviceInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceServiceDeviceInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeviceService_DeviceInfoClient interface {
	Recv() (*ResponseMess, error)
	grpc.ClientStream
}

type deviceServiceDeviceInfoClient struct {
	grpc.ClientStream
}

func (x *deviceServiceDeviceInfoClient) Recv() (*ResponseMess, error) {
	m := new(ResponseMess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	DeviceInfo(*RequestEmpty, DeviceService_DeviceInfoServer) error
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) DeviceInfo(*RequestEmpty, DeviceService_DeviceInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method DeviceInfo not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_DeviceInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeviceServiceServer).DeviceInfo(m, &deviceServiceDeviceInfoServer{stream})
}

type DeviceService_DeviceInfoServer interface {
	Send(*ResponseMess) error
	grpc.ServerStream
}

type deviceServiceDeviceInfoServer struct {
	grpc.ServerStream
}

func (x *deviceServiceDeviceInfoServer) Send(m *ResponseMess) error {
	return x.ServerStream.SendMsg(m)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeviceInfo",
			Handler:       _DeviceService_DeviceInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "team00_server.proto",
}