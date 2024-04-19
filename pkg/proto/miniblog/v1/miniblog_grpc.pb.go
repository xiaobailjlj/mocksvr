// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://miniblog.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: miniblog/v1/miniblog.proto

package v1

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

// MiniBlogClient is the client API for MiniBlog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiniBlogClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type miniBlogClient struct {
	cc grpc.ClientConnInterface
}

func NewMiniBlogClient(cc grpc.ClientConnInterface) MiniBlogClient {
	return &miniBlogClient{cc}
}

func (c *miniBlogClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/v1.MiniBlog/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiniBlogServer is the server API for MiniBlog service.
// All implementations must embed UnimplementedMiniBlogServer
// for forward compatibility
type MiniBlogServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedMiniBlogServer()
}

// UnimplementedMiniBlogServer must be embedded to have forward compatible implementations.
type UnimplementedMiniBlogServer struct {
}

func (UnimplementedMiniBlogServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedMiniBlogServer) mustEmbedUnimplementedMiniBlogServer() {}

// UnsafeMiniBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiniBlogServer will
// result in compilation errors.
type UnsafeMiniBlogServer interface {
	mustEmbedUnimplementedMiniBlogServer()
}

func RegisterMiniBlogServer(s grpc.ServiceRegistrar, srv MiniBlogServer) {
	s.RegisterService(&MiniBlog_ServiceDesc, srv)
}

func _MiniBlog_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniBlogServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MiniBlog/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniBlogServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MiniBlog_ServiceDesc is the grpc.ServiceDesc for MiniBlog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MiniBlog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MiniBlog",
	HandlerType: (*MiniBlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _MiniBlog_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "miniblog/v1/miniblog.proto",
}
