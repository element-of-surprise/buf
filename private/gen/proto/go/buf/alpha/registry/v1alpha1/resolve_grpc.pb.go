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
// source: buf/alpha/registry/v1alpha1/resolve.proto

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

// ResolveServiceClient is the client API for ResolveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResolveServiceClient interface {
	// GetModulePins finds all the latest digests and respective dependencies of
	// the provided module references and picks a set of distinct modules pins.
	//
	// Note that module references with commits should still be passed to this function
	// to make sure this function can do dependency resolution.
	//
	// This function also deals with tiebreaking what ModulePin wins for the same repository.
	GetModulePins(ctx context.Context, in *GetModulePinsRequest, opts ...grpc.CallOption) (*GetModulePinsResponse, error)
}

type resolveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResolveServiceClient(cc grpc.ClientConnInterface) ResolveServiceClient {
	return &resolveServiceClient{cc}
}

func (c *resolveServiceClient) GetModulePins(ctx context.Context, in *GetModulePinsRequest, opts ...grpc.CallOption) (*GetModulePinsResponse, error) {
	out := new(GetModulePinsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.ResolveService/GetModulePins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolveServiceServer is the server API for ResolveService service.
// All implementations should embed UnimplementedResolveServiceServer
// for forward compatibility
type ResolveServiceServer interface {
	// GetModulePins finds all the latest digests and respective dependencies of
	// the provided module references and picks a set of distinct modules pins.
	//
	// Note that module references with commits should still be passed to this function
	// to make sure this function can do dependency resolution.
	//
	// This function also deals with tiebreaking what ModulePin wins for the same repository.
	GetModulePins(context.Context, *GetModulePinsRequest) (*GetModulePinsResponse, error)
}

// UnimplementedResolveServiceServer should be embedded to have forward compatible implementations.
type UnimplementedResolveServiceServer struct {
}

func (UnimplementedResolveServiceServer) GetModulePins(context.Context, *GetModulePinsRequest) (*GetModulePinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModulePins not implemented")
}

// UnsafeResolveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResolveServiceServer will
// result in compilation errors.
type UnsafeResolveServiceServer interface {
	mustEmbedUnimplementedResolveServiceServer()
}

func RegisterResolveServiceServer(s grpc.ServiceRegistrar, srv ResolveServiceServer) {
	s.RegisterService(&ResolveService_ServiceDesc, srv)
}

func _ResolveService_GetModulePins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModulePinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolveServiceServer).GetModulePins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.ResolveService/GetModulePins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolveServiceServer).GetModulePins(ctx, req.(*GetModulePinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResolveService_ServiceDesc is the grpc.ServiceDesc for ResolveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResolveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.ResolveService",
	HandlerType: (*ResolveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModulePins",
			Handler:    _ResolveService_GetModulePins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/resolve.proto",
}

// LocalResolveServiceClient is the client API for LocalResolveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalResolveServiceClient interface {
	// GetLocalModulePins gets the latest pins for the specified local module references.
	// It also includes all of the modules transitive dependencies for the specified references.
	//
	// We want this for two reasons:
	//
	// 1. It makes it easy to say "we know we're looking for owner/repo on this specific remote".
	//    While we could just do this in GetModulePins by being aware of what our remote is
	//    (something we probably still need to know, DNS problems aside, which are more
	//    theoretical), this helps.
	// 2. Having a separate method makes us able to say "do not make decisions about what
	//    wins between competing pins for the same repo". This should only be done in
	//    GetModulePins, not in this function, i.e. only done at the top level.
	GetLocalModulePins(ctx context.Context, in *GetLocalModulePinsRequest, opts ...grpc.CallOption) (*GetLocalModulePinsResponse, error)
}

type localResolveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalResolveServiceClient(cc grpc.ClientConnInterface) LocalResolveServiceClient {
	return &localResolveServiceClient{cc}
}

func (c *localResolveServiceClient) GetLocalModulePins(ctx context.Context, in *GetLocalModulePinsRequest, opts ...grpc.CallOption) (*GetLocalModulePinsResponse, error) {
	out := new(GetLocalModulePinsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.LocalResolveService/GetLocalModulePins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalResolveServiceServer is the server API for LocalResolveService service.
// All implementations should embed UnimplementedLocalResolveServiceServer
// for forward compatibility
type LocalResolveServiceServer interface {
	// GetLocalModulePins gets the latest pins for the specified local module references.
	// It also includes all of the modules transitive dependencies for the specified references.
	//
	// We want this for two reasons:
	//
	// 1. It makes it easy to say "we know we're looking for owner/repo on this specific remote".
	//    While we could just do this in GetModulePins by being aware of what our remote is
	//    (something we probably still need to know, DNS problems aside, which are more
	//    theoretical), this helps.
	// 2. Having a separate method makes us able to say "do not make decisions about what
	//    wins between competing pins for the same repo". This should only be done in
	//    GetModulePins, not in this function, i.e. only done at the top level.
	GetLocalModulePins(context.Context, *GetLocalModulePinsRequest) (*GetLocalModulePinsResponse, error)
}

// UnimplementedLocalResolveServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLocalResolveServiceServer struct {
}

func (UnimplementedLocalResolveServiceServer) GetLocalModulePins(context.Context, *GetLocalModulePinsRequest) (*GetLocalModulePinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalModulePins not implemented")
}

// UnsafeLocalResolveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalResolveServiceServer will
// result in compilation errors.
type UnsafeLocalResolveServiceServer interface {
	mustEmbedUnimplementedLocalResolveServiceServer()
}

func RegisterLocalResolveServiceServer(s grpc.ServiceRegistrar, srv LocalResolveServiceServer) {
	s.RegisterService(&LocalResolveService_ServiceDesc, srv)
}

func _LocalResolveService_GetLocalModulePins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalModulePinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalResolveServiceServer).GetLocalModulePins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.LocalResolveService/GetLocalModulePins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalResolveServiceServer).GetLocalModulePins(ctx, req.(*GetLocalModulePinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalResolveService_ServiceDesc is the grpc.ServiceDesc for LocalResolveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalResolveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.LocalResolveService",
	HandlerType: (*LocalResolveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocalModulePins",
			Handler:    _LocalResolveService_GetLocalModulePins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/resolve.proto",
}
