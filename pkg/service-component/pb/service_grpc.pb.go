// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: service.proto

package pb

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

// ServiceComponentClient is the client API for ServiceComponent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceComponentClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceComponentClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceComponentClient(cc grpc.ClientConnInterface) ServiceComponentClient {
	return &serviceComponentClient{cc}
}

func (c *serviceComponentClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/service_component.ServiceComponent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceComponentClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service_component.ServiceComponent/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceComponentServer is the server API for ServiceComponent service.
// All implementations must embed UnimplementedServiceComponentServer
// for forward compatibility
type ServiceComponentServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedServiceComponentServer()
}

// UnimplementedServiceComponentServer must be embedded to have forward compatible implementations.
type UnimplementedServiceComponentServer struct {
}

func (UnimplementedServiceComponentServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServiceComponentServer) Set(context.Context, *SetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedServiceComponentServer) mustEmbedUnimplementedServiceComponentServer() {}

// UnsafeServiceComponentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceComponentServer will
// result in compilation errors.
type UnsafeServiceComponentServer interface {
	mustEmbedUnimplementedServiceComponentServer()
}

func RegisterServiceComponentServer(s grpc.ServiceRegistrar, srv ServiceComponentServer) {
	s.RegisterService(&ServiceComponent_ServiceDesc, srv)
}

func _ServiceComponent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceComponentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_component.ServiceComponent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceComponentServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceComponent_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceComponentServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_component.ServiceComponent/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceComponentServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceComponent_ServiceDesc is the grpc.ServiceDesc for ServiceComponent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceComponent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_component.ServiceComponent",
	HandlerType: (*ServiceComponentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ServiceComponent_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _ServiceComponent_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}