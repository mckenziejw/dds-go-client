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

// DDSClient is the client API for DDS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DDSClient interface {
	// [brief]: API to get the object hierarchy model
	// [detail]: Returns the schema definition for all
	// the types of objects in the system. It includes
	// the attribute details and the inter object
	// relationships.
	SchemaModelGet(ctx context.Context, in *SchemaModelRequest, opts ...grpc.CallOption) (*SchemaModelResponse, error)
	// [brief]: API to get the object type documentation
	// [detail]: Returns the additional technical information
	// for a given type of object. The information returned
	// will be a high level, human readable description of
	// of the type.
	TypeDocumentationGet(ctx context.Context, in *TypeDocumentationRequest, opts ...grpc.CallOption) (*TypeDocumentationResponse, error)
	// [brief]: API to get the list of subscribe topics
	// [detail]: Topics are labels which are associated
	// with objects. Clients can subscribe to one more
	// topic(s) to recieve a stream of associated objects.
	TopicListGet(ctx context.Context, in *TopicListRequest, opts ...grpc.CallOption) (*TopicListResponse, error)
	// [brief]: API to subscribe for a set of topics
	// [detail]: Returns a stream of objects associated
	// for a given one or more topic(s). Clients can
	// build the system state of interest based on these
	// objects. The stream of objects are continously synced
	// based on the option given in the request.
	TopicSubscribe(ctx context.Context, in *TopicSubscribeRequest, opts ...grpc.CallOption) (DDS_TopicSubscribeClient, error)
	// [brief]: API to fetch application data
	// [detail]: API to fetch data from a specific application
	// and a node. This set of data includes certain stats,
	// telemetry and other CLI data.
	ApplicationDataGet(ctx context.Context, in *ApplicationDataRequest, opts ...grpc.CallOption) (DDS_ApplicationDataGetClient, error)
}

type dDSClient struct {
	cc grpc.ClientConnInterface
}

func NewDDSClient(cc grpc.ClientConnInterface) DDSClient {
	return &dDSClient{cc}
}

func (c *dDSClient) SchemaModelGet(ctx context.Context, in *SchemaModelRequest, opts ...grpc.CallOption) (*SchemaModelResponse, error) {
	out := new(SchemaModelResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.dds.DDS/SchemaModelGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDSClient) TypeDocumentationGet(ctx context.Context, in *TypeDocumentationRequest, opts ...grpc.CallOption) (*TypeDocumentationResponse, error) {
	out := new(TypeDocumentationResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.dds.DDS/TypeDocumentationGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDSClient) TopicListGet(ctx context.Context, in *TopicListRequest, opts ...grpc.CallOption) (*TopicListResponse, error) {
	out := new(TopicListResponse)
	err := c.cc.Invoke(ctx, "/jnx.jet.dds.DDS/TopicListGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDSClient) TopicSubscribe(ctx context.Context, in *TopicSubscribeRequest, opts ...grpc.CallOption) (DDS_TopicSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &DDS_ServiceDesc.Streams[0], "/jnx.jet.dds.DDS/TopicSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dDSTopicSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DDS_TopicSubscribeClient interface {
	Recv() (*ObjectStreamResponse, error)
	grpc.ClientStream
}

type dDSTopicSubscribeClient struct {
	grpc.ClientStream
}

func (x *dDSTopicSubscribeClient) Recv() (*ObjectStreamResponse, error) {
	m := new(ObjectStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dDSClient) ApplicationDataGet(ctx context.Context, in *ApplicationDataRequest, opts ...grpc.CallOption) (DDS_ApplicationDataGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &DDS_ServiceDesc.Streams[1], "/jnx.jet.dds.DDS/ApplicationDataGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &dDSApplicationDataGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DDS_ApplicationDataGetClient interface {
	Recv() (*DataStreamResponse, error)
	grpc.ClientStream
}

type dDSApplicationDataGetClient struct {
	grpc.ClientStream
}

func (x *dDSApplicationDataGetClient) Recv() (*DataStreamResponse, error) {
	m := new(DataStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DDSServer is the server API for DDS service.
// All implementations must embed UnimplementedDDSServer
// for forward compatibility
type DDSServer interface {
	// [brief]: API to get the object hierarchy model
	// [detail]: Returns the schema definition for all
	// the types of objects in the system. It includes
	// the attribute details and the inter object
	// relationships.
	SchemaModelGet(context.Context, *SchemaModelRequest) (*SchemaModelResponse, error)
	// [brief]: API to get the object type documentation
	// [detail]: Returns the additional technical information
	// for a given type of object. The information returned
	// will be a high level, human readable description of
	// of the type.
	TypeDocumentationGet(context.Context, *TypeDocumentationRequest) (*TypeDocumentationResponse, error)
	// [brief]: API to get the list of subscribe topics
	// [detail]: Topics are labels which are associated
	// with objects. Clients can subscribe to one more
	// topic(s) to recieve a stream of associated objects.
	TopicListGet(context.Context, *TopicListRequest) (*TopicListResponse, error)
	// [brief]: API to subscribe for a set of topics
	// [detail]: Returns a stream of objects associated
	// for a given one or more topic(s). Clients can
	// build the system state of interest based on these
	// objects. The stream of objects are continously synced
	// based on the option given in the request.
	TopicSubscribe(*TopicSubscribeRequest, DDS_TopicSubscribeServer) error
	// [brief]: API to fetch application data
	// [detail]: API to fetch data from a specific application
	// and a node. This set of data includes certain stats,
	// telemetry and other CLI data.
	ApplicationDataGet(*ApplicationDataRequest, DDS_ApplicationDataGetServer) error
	mustEmbedUnimplementedDDSServer()
}

// UnimplementedDDSServer must be embedded to have forward compatible implementations.
type UnimplementedDDSServer struct {
}

func (UnimplementedDDSServer) SchemaModelGet(context.Context, *SchemaModelRequest) (*SchemaModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaModelGet not implemented")
}
func (UnimplementedDDSServer) TypeDocumentationGet(context.Context, *TypeDocumentationRequest) (*TypeDocumentationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TypeDocumentationGet not implemented")
}
func (UnimplementedDDSServer) TopicListGet(context.Context, *TopicListRequest) (*TopicListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicListGet not implemented")
}
func (UnimplementedDDSServer) TopicSubscribe(*TopicSubscribeRequest, DDS_TopicSubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method TopicSubscribe not implemented")
}
func (UnimplementedDDSServer) ApplicationDataGet(*ApplicationDataRequest, DDS_ApplicationDataGetServer) error {
	return status.Errorf(codes.Unimplemented, "method ApplicationDataGet not implemented")
}
func (UnimplementedDDSServer) mustEmbedUnimplementedDDSServer() {}

// UnsafeDDSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DDSServer will
// result in compilation errors.
type UnsafeDDSServer interface {
	mustEmbedUnimplementedDDSServer()
}

func RegisterDDSServer(s grpc.ServiceRegistrar, srv DDSServer) {
	s.RegisterService(&DDS_ServiceDesc, srv)
}

func _DDS_SchemaModelGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDSServer).SchemaModelGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.dds.DDS/SchemaModelGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDSServer).SchemaModelGet(ctx, req.(*SchemaModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDS_TypeDocumentationGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeDocumentationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDSServer).TypeDocumentationGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.dds.DDS/TypeDocumentationGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDSServer).TypeDocumentationGet(ctx, req.(*TypeDocumentationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDS_TopicListGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDSServer).TopicListGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jnx.jet.dds.DDS/TopicListGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDSServer).TopicListGet(ctx, req.(*TopicListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDS_TopicSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopicSubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DDSServer).TopicSubscribe(m, &dDSTopicSubscribeServer{stream})
}

type DDS_TopicSubscribeServer interface {
	Send(*ObjectStreamResponse) error
	grpc.ServerStream
}

type dDSTopicSubscribeServer struct {
	grpc.ServerStream
}

func (x *dDSTopicSubscribeServer) Send(m *ObjectStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DDS_ApplicationDataGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplicationDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DDSServer).ApplicationDataGet(m, &dDSApplicationDataGetServer{stream})
}

type DDS_ApplicationDataGetServer interface {
	Send(*DataStreamResponse) error
	grpc.ServerStream
}

type dDSApplicationDataGetServer struct {
	grpc.ServerStream
}

func (x *dDSApplicationDataGetServer) Send(m *DataStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DDS_ServiceDesc is the grpc.ServiceDesc for DDS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DDS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jnx.jet.dds.DDS",
	HandlerType: (*DDSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SchemaModelGet",
			Handler:    _DDS_SchemaModelGet_Handler,
		},
		{
			MethodName: "TypeDocumentationGet",
			Handler:    _DDS_TypeDocumentationGet_Handler,
		},
		{
			MethodName: "TopicListGet",
			Handler:    _DDS_TopicListGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TopicSubscribe",
			Handler:       _DDS_TopicSubscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ApplicationDataGet",
			Handler:       _DDS_ApplicationDataGet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "jnx_dds_service.proto",
}
