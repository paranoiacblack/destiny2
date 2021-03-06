// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package applications

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	GetApplicationAPIUsage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*APIUsage, error)
	GetBungieApplications(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApplicationsResponse, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) GetApplicationAPIUsage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*APIUsage, error) {
	out := new(APIUsage)
	err := c.cc.Invoke(ctx, "/applications.App/GetApplicationAPIUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetBungieApplications(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApplicationsResponse, error) {
	out := new(ApplicationsResponse)
	err := c.cc.Invoke(ctx, "/applications.App/GetBungieApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	GetApplicationAPIUsage(context.Context, *UsageRequest) (*APIUsage, error)
	GetBungieApplications(context.Context, *emptypb.Empty) (*ApplicationsResponse, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) GetApplicationAPIUsage(context.Context, *UsageRequest) (*APIUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationAPIUsage not implemented")
}
func (UnimplementedAppServer) GetBungieApplications(context.Context, *emptypb.Empty) (*ApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBungieApplications not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_GetApplicationAPIUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetApplicationAPIUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications.App/GetApplicationAPIUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetApplicationAPIUsage(ctx, req.(*UsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetBungieApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetBungieApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications.App/GetBungieApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetBungieApplications(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "applications.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApplicationAPIUsage",
			Handler:    _App_GetApplicationAPIUsage_Handler,
		},
		{
			MethodName: "GetBungieApplications",
			Handler:    _App_GetBungieApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/applications/applications.proto",
}
