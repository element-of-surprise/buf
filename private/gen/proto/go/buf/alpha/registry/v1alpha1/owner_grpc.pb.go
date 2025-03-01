// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             (unknown)
// source: buf/alpha/registry/v1alpha1/owner.proto

package registryv1alpha1

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

// OwnerServiceClient is the client API for OwnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnerServiceClient interface {
	// GetOwnerByName takes an owner name and returns the owner as
	// either a user or organization.
	GetOwnerByName(ctx context.Context, in *GetOwnerByNameRequest, opts ...grpc.CallOption) (*GetOwnerByNameResponse, error)
}

type ownerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnerServiceClient(cc grpc.ClientConnInterface) OwnerServiceClient {
	return &ownerServiceClient{cc}
}

func (c *ownerServiceClient) GetOwnerByName(ctx context.Context, in *GetOwnerByNameRequest, opts ...grpc.CallOption) (*GetOwnerByNameResponse, error) {
	out := new(GetOwnerByNameResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.OwnerService/GetOwnerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerServiceServer is the server API for OwnerService service.
// All implementations should embed UnimplementedOwnerServiceServer
// for forward compatibility
type OwnerServiceServer interface {
	// GetOwnerByName takes an owner name and returns the owner as
	// either a user or organization.
	GetOwnerByName(context.Context, *GetOwnerByNameRequest) (*GetOwnerByNameResponse, error)
}

// UnimplementedOwnerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOwnerServiceServer struct {
}

func (UnimplementedOwnerServiceServer) GetOwnerByName(context.Context, *GetOwnerByNameRequest) (*GetOwnerByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnerByName not implemented")
}

// UnsafeOwnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnerServiceServer will
// result in compilation errors.
type UnsafeOwnerServiceServer interface {
	mustEmbedUnimplementedOwnerServiceServer()
}

func RegisterOwnerServiceServer(s grpc.ServiceRegistrar, srv OwnerServiceServer) {
	s.RegisterService(&OwnerService_ServiceDesc, srv)
}

func _OwnerService_GetOwnerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).GetOwnerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.OwnerService/GetOwnerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).GetOwnerByName(ctx, req.(*GetOwnerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OwnerService_ServiceDesc is the grpc.ServiceDesc for OwnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OwnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.OwnerService",
	HandlerType: (*OwnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOwnerByName",
			Handler:    _OwnerService_GetOwnerByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/owner.proto",
}
