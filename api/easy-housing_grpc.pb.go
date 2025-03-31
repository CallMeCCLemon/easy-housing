// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: api/easy-housing.proto

package api

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
	EasyHousing_Echo_FullMethodName = "/easy_housing.EasyHousing/Echo"
)

// EasyHousingClient is the client API for EasyHousing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EasyHousingClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type easyHousingClient struct {
	cc grpc.ClientConnInterface
}

func NewEasyHousingClient(cc grpc.ClientConnInterface) EasyHousingClient {
	return &easyHousingClient{cc}
}

func (c *easyHousingClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, EasyHousing_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EasyHousingServer is the server API for EasyHousing service.
// All implementations must embed UnimplementedEasyHousingServer
// for forward compatibility.
type EasyHousingServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedEasyHousingServer()
}

// UnimplementedEasyHousingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEasyHousingServer struct{}

func (UnimplementedEasyHousingServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEasyHousingServer) mustEmbedUnimplementedEasyHousingServer() {}
func (UnimplementedEasyHousingServer) testEmbeddedByValue()                     {}

// UnsafeEasyHousingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EasyHousingServer will
// result in compilation errors.
type UnsafeEasyHousingServer interface {
	mustEmbedUnimplementedEasyHousingServer()
}

func RegisterEasyHousingServer(s grpc.ServiceRegistrar, srv EasyHousingServer) {
	// If the following call pancis, it indicates UnimplementedEasyHousingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EasyHousing_ServiceDesc, srv)
}

func _EasyHousing_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EasyHousingServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EasyHousing_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EasyHousingServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EasyHousing_ServiceDesc is the grpc.ServiceDesc for EasyHousing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EasyHousing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "easy_housing.EasyHousing",
	HandlerType: (*EasyHousingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EasyHousing_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/easy-housing.proto",
}
