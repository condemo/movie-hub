// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: data_handler.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DataHandler_GetLastUpdates_FullMethodName = "/data_handler.DataHandler/GetLastUpdates"
)

// DataHandlerClient is the client API for DataHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataHandlerClient interface {
	GetLastUpdates(ctx context.Context, in *LastUpdatesRequest, opts ...grpc.CallOption) (*MediaResponse, error)
}

type dataHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataHandlerClient(cc grpc.ClientConnInterface) DataHandlerClient {
	return &dataHandlerClient{cc}
}

func (c *dataHandlerClient) GetLastUpdates(ctx context.Context, in *LastUpdatesRequest, opts ...grpc.CallOption) (*MediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MediaResponse)
	err := c.cc.Invoke(ctx, DataHandler_GetLastUpdates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataHandlerServer is the server API for DataHandler service.
// All implementations must embed UnimplementedDataHandlerServer
// for forward compatibility.
type DataHandlerServer interface {
	GetLastUpdates(context.Context, *LastUpdatesRequest) (*MediaResponse, error)
	mustEmbedUnimplementedDataHandlerServer()
}

// UnimplementedDataHandlerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataHandlerServer struct{}

func (UnimplementedDataHandlerServer) GetLastUpdates(context.Context, *LastUpdatesRequest) (*MediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastUpdates not implemented")
}
func (UnimplementedDataHandlerServer) mustEmbedUnimplementedDataHandlerServer() {}
func (UnimplementedDataHandlerServer) testEmbeddedByValue()                     {}

// UnsafeDataHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataHandlerServer will
// result in compilation errors.
type UnsafeDataHandlerServer interface {
	mustEmbedUnimplementedDataHandlerServer()
}

func RegisterDataHandlerServer(s grpc.ServiceRegistrar, srv DataHandlerServer) {
	// If the following call pancis, it indicates UnimplementedDataHandlerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataHandler_ServiceDesc, srv)
}

func _DataHandler_GetLastUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataHandlerServer).GetLastUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataHandler_GetLastUpdates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataHandlerServer).GetLastUpdates(ctx, req.(*LastUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataHandler_ServiceDesc is the grpc.ServiceDesc for DataHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data_handler.DataHandler",
	HandlerType: (*DataHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastUpdates",
			Handler:    _DataHandler_GetLastUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_handler.proto",
}
