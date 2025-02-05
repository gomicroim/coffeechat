// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: api/pushjob/pushjob.proto

package pushjob

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

// PushJobClient is the client API for PushJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushJobClient interface {
	// 通过ws发送给单个用户
	PushToSingleUser(ctx context.Context, in *SingleUserRequest, opts ...grpc.CallOption) (*SingleUserReply, error)
	// 通过ws发送给多个用户
	PushToBatchUser(ctx context.Context, in *BatchUserRequest, opts ...grpc.CallOption) (*BatchUserReply, error)
	// 通过ws发送给所有在线用户
	PushToAllOnlineUser(ctx context.Context, in *AllOnlineUserReq, opts ...grpc.CallOption) (*AllOnlineUserReply, error)
}

type pushJobClient struct {
	cc grpc.ClientConnInterface
}

func NewPushJobClient(cc grpc.ClientConnInterface) PushJobClient {
	return &pushJobClient{cc}
}

func (c *pushJobClient) PushToSingleUser(ctx context.Context, in *SingleUserRequest, opts ...grpc.CallOption) (*SingleUserReply, error) {
	out := new(SingleUserReply)
	err := c.cc.Invoke(ctx, "/pushjob.PushJob/PushToSingleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushJobClient) PushToBatchUser(ctx context.Context, in *BatchUserRequest, opts ...grpc.CallOption) (*BatchUserReply, error) {
	out := new(BatchUserReply)
	err := c.cc.Invoke(ctx, "/pushjob.PushJob/PushToBatchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushJobClient) PushToAllOnlineUser(ctx context.Context, in *AllOnlineUserReq, opts ...grpc.CallOption) (*AllOnlineUserReply, error) {
	out := new(AllOnlineUserReply)
	err := c.cc.Invoke(ctx, "/pushjob.PushJob/PushToAllOnlineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushJobServer is the server API for PushJob service.
// All implementations must embed UnimplementedPushJobServer
// for forward compatibility
type PushJobServer interface {
	// 通过ws发送给单个用户
	PushToSingleUser(context.Context, *SingleUserRequest) (*SingleUserReply, error)
	// 通过ws发送给多个用户
	PushToBatchUser(context.Context, *BatchUserRequest) (*BatchUserReply, error)
	// 通过ws发送给所有在线用户
	PushToAllOnlineUser(context.Context, *AllOnlineUserReq) (*AllOnlineUserReply, error)
	mustEmbedUnimplementedPushJobServer()
}

// UnimplementedPushJobServer must be embedded to have forward compatible implementations.
type UnimplementedPushJobServer struct {
}

func (UnimplementedPushJobServer) PushToSingleUser(context.Context, *SingleUserRequest) (*SingleUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToSingleUser not implemented")
}
func (UnimplementedPushJobServer) PushToBatchUser(context.Context, *BatchUserRequest) (*BatchUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToBatchUser not implemented")
}
func (UnimplementedPushJobServer) PushToAllOnlineUser(context.Context, *AllOnlineUserReq) (*AllOnlineUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToAllOnlineUser not implemented")
}
func (UnimplementedPushJobServer) mustEmbedUnimplementedPushJobServer() {}

// UnsafePushJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushJobServer will
// result in compilation errors.
type UnsafePushJobServer interface {
	mustEmbedUnimplementedPushJobServer()
}

func RegisterPushJobServer(s grpc.ServiceRegistrar, srv PushJobServer) {
	s.RegisterService(&PushJob_ServiceDesc, srv)
}

func _PushJob_PushToSingleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushJobServer).PushToSingleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pushjob.PushJob/PushToSingleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushJobServer).PushToSingleUser(ctx, req.(*SingleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushJob_PushToBatchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushJobServer).PushToBatchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pushjob.PushJob/PushToBatchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushJobServer).PushToBatchUser(ctx, req.(*BatchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushJob_PushToAllOnlineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllOnlineUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushJobServer).PushToAllOnlineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pushjob.PushJob/PushToAllOnlineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushJobServer).PushToAllOnlineUser(ctx, req.(*AllOnlineUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PushJob_ServiceDesc is the grpc.ServiceDesc for PushJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pushjob.PushJob",
	HandlerType: (*PushJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushToSingleUser",
			Handler:    _PushJob_PushToSingleUser_Handler,
		},
		{
			MethodName: "PushToBatchUser",
			Handler:    _PushJob_PushToBatchUser_Handler,
		},
		{
			MethodName: "PushToAllOnlineUser",
			Handler:    _PushJob_PushToAllOnlineUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pushjob/pushjob.proto",
}
