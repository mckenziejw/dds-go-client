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

// EcmptracerFlowClient is the client API for EcmptracerFlow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcmptracerFlowClient interface {
	// [brief]: Add new flow
	// [detail]: This rpc call returns a success or a failure reason depending upon the status of the flow added
	EcmptracerFlowAdd(ctx context.Context, in *EcmptracerFlowAddRequest, opts ...grpc.CallOption) (*EcmptracerFlowMonitorResponse, error)
	// [brief]: Delete flow and corresponding filters
	// [detail]: This rpc call returns a success or a failure depending upon the status of the flows deleted
	EcmptracerFlowDelete(ctx context.Context, in *EcmptracerFlowInfoRequest, opts ...grpc.CallOption) (*EcmptracerFlowMonitorResponse, error)
	// [brief]: Get counter stats for a flow
	// [detail]: This rpc call returns the counter names, pkts and bytes count for the flow
	EcmptracerFlowGet(ctx context.Context, in *EcmptracerFlowInfoRequest, opts ...grpc.CallOption) (*EcmptracerFlowInfoResponse, error)
}

type ecmptracerFlowClient struct {
	cc grpc.ClientConnInterface
}

func NewEcmptracerFlowClient(cc grpc.ClientConnInterface) EcmptracerFlowClient {
	return &ecmptracerFlowClient{cc}
}

func (c *ecmptracerFlowClient) EcmptracerFlowAdd(ctx context.Context, in *EcmptracerFlowAddRequest, opts ...grpc.CallOption) (*EcmptracerFlowMonitorResponse, error) {
	out := new(EcmptracerFlowMonitorResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.ecmptracer.EcmptracerFlow/EcmptracerFlowAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecmptracerFlowClient) EcmptracerFlowDelete(ctx context.Context, in *EcmptracerFlowInfoRequest, opts ...grpc.CallOption) (*EcmptracerFlowMonitorResponse, error) {
	out := new(EcmptracerFlowMonitorResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.ecmptracer.EcmptracerFlow/EcmptracerFlowDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecmptracerFlowClient) EcmptracerFlowGet(ctx context.Context, in *EcmptracerFlowInfoRequest, opts ...grpc.CallOption) (*EcmptracerFlowInfoResponse, error) {
	out := new(EcmptracerFlowInfoResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.ecmptracer.EcmptracerFlow/EcmptracerFlowGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcmptracerFlowServer is the server API for EcmptracerFlow service.
// All implementations must embed UnimplementedEcmptracerFlowServer
// for forward compatibility
type EcmptracerFlowServer interface {
	// [brief]: Add new flow
	// [detail]: This rpc call returns a success or a failure reason depending upon the status of the flow added
	EcmptracerFlowAdd(context.Context, *EcmptracerFlowAddRequest) (*EcmptracerFlowMonitorResponse, error)
	// [brief]: Delete flow and corresponding filters
	// [detail]: This rpc call returns a success or a failure depending upon the status of the flows deleted
	EcmptracerFlowDelete(context.Context, *EcmptracerFlowInfoRequest) (*EcmptracerFlowMonitorResponse, error)
	// [brief]: Get counter stats for a flow
	// [detail]: This rpc call returns the counter names, pkts and bytes count for the flow
	EcmptracerFlowGet(context.Context, *EcmptracerFlowInfoRequest) (*EcmptracerFlowInfoResponse, error)
	mustEmbedUnimplementedEcmptracerFlowServer()
}

// UnimplementedEcmptracerFlowServer must be embedded to have forward compatible implementations.
type UnimplementedEcmptracerFlowServer struct {
}

func (UnimplementedEcmptracerFlowServer) EcmptracerFlowAdd(context.Context, *EcmptracerFlowAddRequest) (*EcmptracerFlowMonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EcmptracerFlowAdd not implemented")
}
func (UnimplementedEcmptracerFlowServer) EcmptracerFlowDelete(context.Context, *EcmptracerFlowInfoRequest) (*EcmptracerFlowMonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EcmptracerFlowDelete not implemented")
}
func (UnimplementedEcmptracerFlowServer) EcmptracerFlowGet(context.Context, *EcmptracerFlowInfoRequest) (*EcmptracerFlowInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EcmptracerFlowGet not implemented")
}
func (UnimplementedEcmptracerFlowServer) mustEmbedUnimplementedEcmptracerFlowServer() {}

// UnsafeEcmptracerFlowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcmptracerFlowServer will
// result in compilation errors.
type UnsafeEcmptracerFlowServer interface {
	mustEmbedUnimplementedEcmptracerFlowServer()
}

func RegisterEcmptracerFlowServer(s grpc.ServiceRegistrar, srv EcmptracerFlowServer) {
	s.RegisterService(&EcmptracerFlow_ServiceDesc, srv)
}

func _EcmptracerFlow_EcmptracerFlowAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EcmptracerFlowAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcmptracerFlowServer).EcmptracerFlowAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.ecmptracer.EcmptracerFlow/EcmptracerFlowAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcmptracerFlowServer).EcmptracerFlowAdd(ctx, req.(*EcmptracerFlowAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcmptracerFlow_EcmptracerFlowDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EcmptracerFlowInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcmptracerFlowServer).EcmptracerFlowDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.ecmptracer.EcmptracerFlow/EcmptracerFlowDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcmptracerFlowServer).EcmptracerFlowDelete(ctx, req.(*EcmptracerFlowInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcmptracerFlow_EcmptracerFlowGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EcmptracerFlowInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcmptracerFlowServer).EcmptracerFlowGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.ecmptracer.EcmptracerFlow/EcmptracerFlowGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcmptracerFlowServer).EcmptracerFlowGet(ctx, req.(*EcmptracerFlowInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EcmptracerFlow_ServiceDesc is the grpc.ServiceDesc for EcmptracerFlow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EcmptracerFlow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jnx.jet.ecmptracer.EcmptracerFlow",
	HandlerType: (*EcmptracerFlowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EcmptracerFlowAdd",
			Handler:    _EcmptracerFlow_EcmptracerFlowAdd_Handler,
		},
		{
			MethodName: "EcmptracerFlowDelete",
			Handler:    _EcmptracerFlow_EcmptracerFlowDelete_Handler,
		},
		{
			MethodName: "EcmptracerFlowGet",
			Handler:    _EcmptracerFlow_EcmptracerFlowGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jnx_ecmptracer_service.proto",
}
