// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.5
// source: proto/user.proto

package proto

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

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	Userinfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) Userinfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/Business.UserInfo/Userinfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	Userinfo(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (UnimplementedUserInfoServer) Userinfo(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Userinfo not implemented")
}
func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_Userinfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).Userinfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Business.UserInfo/Userinfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).Userinfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Business.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Userinfo",
			Handler:    _UserInfo_Userinfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
