// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/grpc/user_v1/service.proto

package user_v1

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

// UserV1Client is the client API for UserV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserV1Client interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	InputStream(ctx context.Context, opts ...grpc.CallOption) (UserV1_InputStreamClient, error)
	OutputStream(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (UserV1_OutputStreamClient, error)
	BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (UserV1_BidirectionalStreamClient, error)
}

type userV1Client struct {
	cc grpc.ClientConnInterface
}

func NewUserV1Client(cc grpc.ClientConnInterface) UserV1Client {
	return &userV1Client{cc}
}

func (c *userV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) InputStream(ctx context.Context, opts ...grpc.CallOption) (UserV1_InputStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserV1_ServiceDesc.Streams[0], "/user_v1.UserV1/InputStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userV1InputStreamClient{stream}
	return x, nil
}

type UserV1_InputStreamClient interface {
	Send(*GetRequest) error
	CloseAndRecv() (*GetResponse, error)
	grpc.ClientStream
}

type userV1InputStreamClient struct {
	grpc.ClientStream
}

func (x *userV1InputStreamClient) Send(m *GetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userV1InputStreamClient) CloseAndRecv() (*GetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userV1Client) OutputStream(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (UserV1_OutputStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserV1_ServiceDesc.Streams[1], "/user_v1.UserV1/OutputStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userV1OutputStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserV1_OutputStreamClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type userV1OutputStreamClient struct {
	grpc.ClientStream
}

func (x *userV1OutputStreamClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userV1Client) BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (UserV1_BidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserV1_ServiceDesc.Streams[2], "/user_v1.UserV1/BidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userV1BidirectionalStreamClient{stream}
	return x, nil
}

type UserV1_BidirectionalStreamClient interface {
	Send(*GetRequest) error
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type userV1BidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *userV1BidirectionalStreamClient) Send(m *GetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userV1BidirectionalStreamClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserV1Server is the server API for UserV1 service.
// All implementations must embed UnimplementedUserV1Server
// for forward compatibility
type UserV1Server interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	InputStream(UserV1_InputStreamServer) error
	OutputStream(*GetRequest, UserV1_OutputStreamServer) error
	BidirectionalStream(UserV1_BidirectionalStreamServer) error
	mustEmbedUnimplementedUserV1Server()
}

// UnimplementedUserV1Server must be embedded to have forward compatible implementations.
type UnimplementedUserV1Server struct {
}

func (UnimplementedUserV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserV1Server) InputStream(UserV1_InputStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method InputStream not implemented")
}
func (UnimplementedUserV1Server) OutputStream(*GetRequest, UserV1_OutputStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OutputStream not implemented")
}
func (UnimplementedUserV1Server) BidirectionalStream(UserV1_BidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStream not implemented")
}
func (UnimplementedUserV1Server) mustEmbedUnimplementedUserV1Server() {}

// UnsafeUserV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserV1Server will
// result in compilation errors.
type UnsafeUserV1Server interface {
	mustEmbedUnimplementedUserV1Server()
}

func RegisterUserV1Server(s grpc.ServiceRegistrar, srv UserV1Server) {
	s.RegisterService(&UserV1_ServiceDesc, srv)
}

func _UserV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_InputStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserV1Server).InputStream(&userV1InputStreamServer{stream})
}

type UserV1_InputStreamServer interface {
	SendAndClose(*GetResponse) error
	Recv() (*GetRequest, error)
	grpc.ServerStream
}

type userV1InputStreamServer struct {
	grpc.ServerStream
}

func (x *userV1InputStreamServer) SendAndClose(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userV1InputStreamServer) Recv() (*GetRequest, error) {
	m := new(GetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserV1_OutputStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserV1Server).OutputStream(m, &userV1OutputStreamServer{stream})
}

type UserV1_OutputStreamServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type userV1OutputStreamServer struct {
	grpc.ServerStream
}

func (x *userV1OutputStreamServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserV1_BidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserV1Server).BidirectionalStream(&userV1BidirectionalStreamServer{stream})
}

type UserV1_BidirectionalStreamServer interface {
	Send(*GetResponse) error
	Recv() (*GetRequest, error)
	grpc.ServerStream
}

type userV1BidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *userV1BidirectionalStreamServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userV1BidirectionalStreamServer) Recv() (*GetRequest, error) {
	m := new(GetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserV1_ServiceDesc is the grpc.ServiceDesc for UserV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.UserV1",
	HandlerType: (*UserV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserV1_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InputStream",
			Handler:       _UserV1_InputStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OutputStream",
			Handler:       _UserV1_OutputStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStream",
			Handler:       _UserV1_BidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/grpc/user_v1/service.proto",
}