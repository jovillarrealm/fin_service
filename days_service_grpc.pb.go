// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/days_service.proto

package days

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

// DaysServiceClient is the client API for DaysService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaysServiceClient interface {
	GetDays(ctx context.Context, in *CondicionesIniciales, opts ...grpc.CallOption) (*Dias, error)
}

type daysServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDaysServiceClient(cc grpc.ClientConnInterface) DaysServiceClient {
	return &daysServiceClient{cc}
}

func (c *daysServiceClient) GetDays(ctx context.Context, in *CondicionesIniciales, opts ...grpc.CallOption) (*Dias, error) {
	out := new(Dias)
	err := c.cc.Invoke(ctx, "/days_service.days_service/GetDays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaysServiceServer is the server API for DaysService service.
// All implementations must embed UnimplementedDaysServiceServer
// for forward compatibility
type DaysServiceServer interface {
	GetDays(context.Context, *CondicionesIniciales) (*Dias, error)
	mustEmbedUnimplementedDaysServiceServer()
}

// UnimplementedDaysServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDaysServiceServer struct {
}

func (UnimplementedDaysServiceServer) GetDays(context.Context, *CondicionesIniciales) (*Dias, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDays not implemented")
}
func (UnimplementedDaysServiceServer) mustEmbedUnimplementedDaysServiceServer() {}

// UnsafeDaysServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaysServiceServer will
// result in compilation errors.
type UnsafeDaysServiceServer interface {
	mustEmbedUnimplementedDaysServiceServer()
}

func RegisterDaysServiceServer(s grpc.ServiceRegistrar, srv DaysServiceServer) {
	s.RegisterService(&DaysService_ServiceDesc, srv)
}

func _DaysService_GetDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CondicionesIniciales)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaysServiceServer).GetDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/days_service.days_service/GetDays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaysServiceServer).GetDays(ctx, req.(*CondicionesIniciales))
	}
	return interceptor(ctx, in, info, handler)
}

// DaysService_ServiceDesc is the grpc.ServiceDesc for DaysService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DaysService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "days_service.days_service",
	HandlerType: (*DaysServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDays",
			Handler:    _DaysService_GetDays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/days_service.proto",
}
