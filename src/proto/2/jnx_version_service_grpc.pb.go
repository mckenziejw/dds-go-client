// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// VersionDiscoveryClient is the client API for VersionDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionDiscoveryClient interface {
	// [detail]: Get the Junos version
	// [brief]: Get the Junos version
	JunosVersionGet(ctx context.Context, in *JunosVersionRequest, opts ...grpc.CallOption) (*JunosVersionResponse, error)
	// [detail]: Get List of Services
	// [brief]: Get List of Services
	ServiceListGet(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListResponse, error)
	// [detail]: Get List of all RPCs given a Service name
	// [brief]: Get List of all RPCs given a Service name
	RpcListGet(ctx context.Context, in *RpcListRequest, opts ...grpc.CallOption) (*RpcListResponse, error)
	// [detail]: Get the Service version of the given RPC.
	// [brief]: Get the Service version of the given RPC.
	ServiceVersionGet(ctx context.Context, in *ServiceVersionRequest, opts ...grpc.CallOption) (*ServiceVersionResponse, error)
}

type versionDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionDiscoveryClient(cc grpc.ClientConnInterface) VersionDiscoveryClient {
	return &versionDiscoveryClient{cc}
}

func (c *versionDiscoveryClient) JunosVersionGet(ctx context.Context, in *JunosVersionRequest, opts ...grpc.CallOption) (*JunosVersionResponse, error) {
	out := new(JunosVersionResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.version.VersionDiscovery/JunosVersionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionDiscoveryClient) ServiceListGet(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListResponse, error) {
	out := new(ServiceListResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.version.VersionDiscovery/ServiceListGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionDiscoveryClient) RpcListGet(ctx context.Context, in *RpcListRequest, opts ...grpc.CallOption) (*RpcListResponse, error) {
	out := new(RpcListResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.version.VersionDiscovery/RpcListGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionDiscoveryClient) ServiceVersionGet(ctx context.Context, in *ServiceVersionRequest, opts ...grpc.CallOption) (*ServiceVersionResponse, error) {
	out := new(ServiceVersionResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.version.VersionDiscovery/ServiceVersionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionDiscoveryServer is the server API for VersionDiscovery service.
// All implementations must embed UnimplementedVersionDiscoveryServer
// for forward compatibility
type VersionDiscoveryServer interface {
	// [detail]: Get the Junos version
	// [brief]: Get the Junos version
	JunosVersionGet(context.Context, *JunosVersionRequest) (*JunosVersionResponse, error)
	// [detail]: Get List of Services
	// [brief]: Get List of Services
	ServiceListGet(context.Context, *ServiceListRequest) (*ServiceListResponse, error)
	// [detail]: Get List of all RPCs given a Service name
	// [brief]: Get List of all RPCs given a Service name
	RpcListGet(context.Context, *RpcListRequest) (*RpcListResponse, error)
	// [detail]: Get the Service version of the given RPC.
	// [brief]: Get the Service version of the given RPC.
	ServiceVersionGet(context.Context, *ServiceVersionRequest) (*ServiceVersionResponse, error)
	mustEmbedUnimplementedVersionDiscoveryServer()
}

// UnimplementedVersionDiscoveryServer must be embedded to have forward compatible implementations.
type UnimplementedVersionDiscoveryServer struct {
}

func (UnimplementedVersionDiscoveryServer) JunosVersionGet(context.Context, *JunosVersionRequest) (*JunosVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JunosVersionGet not implemented")
}
func (UnimplementedVersionDiscoveryServer) ServiceListGet(context.Context, *ServiceListRequest) (*ServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceListGet not implemented")
}
func (UnimplementedVersionDiscoveryServer) RpcListGet(context.Context, *RpcListRequest) (*RpcListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcListGet not implemented")
}
func (UnimplementedVersionDiscoveryServer) ServiceVersionGet(context.Context, *ServiceVersionRequest) (*ServiceVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceVersionGet not implemented")
}
func (UnimplementedVersionDiscoveryServer) mustEmbedUnimplementedVersionDiscoveryServer() {}

// UnsafeVersionDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionDiscoveryServer will
// result in compilation errors.
type UnsafeVersionDiscoveryServer interface {
	mustEmbedUnimplementedVersionDiscoveryServer()
}

func RegisterVersionDiscoveryServer(s grpc.ServiceRegistrar, srv VersionDiscoveryServer) {
	s.RegisterService(&VersionDiscovery_ServiceDesc, srv)
}

func _VersionDiscovery_JunosVersionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JunosVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionDiscoveryServer).JunosVersionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.version.VersionDiscovery/JunosVersionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionDiscoveryServer).JunosVersionGet(ctx, req.(*JunosVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionDiscovery_ServiceListGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionDiscoveryServer).ServiceListGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.version.VersionDiscovery/ServiceListGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionDiscoveryServer).ServiceListGet(ctx, req.(*ServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionDiscovery_RpcListGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionDiscoveryServer).RpcListGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.version.VersionDiscovery/RpcListGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionDiscoveryServer).RpcListGet(ctx, req.(*RpcListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionDiscovery_ServiceVersionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionDiscoveryServer).ServiceVersionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.version.VersionDiscovery/ServiceVersionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionDiscoveryServer).ServiceVersionGet(ctx, req.(*ServiceVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionDiscovery_ServiceDesc is the grpc.ServiceDesc for VersionDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jnx.jet.version.VersionDiscovery",
	HandlerType: (*VersionDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JunosVersionGet",
			Handler:    _VersionDiscovery_JunosVersionGet_Handler,
		},
		{
			MethodName: "ServiceListGet",
			Handler:    _VersionDiscovery_ServiceListGet_Handler,
		},
		{
			MethodName: "RpcListGet",
			Handler:    _VersionDiscovery_RpcListGet_Handler,
		},
		{
			MethodName: "ServiceVersionGet",
			Handler:    _VersionDiscovery_ServiceVersionGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jnx_version_service.proto",
}