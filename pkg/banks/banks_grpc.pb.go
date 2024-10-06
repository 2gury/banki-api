// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: banks.proto

package banks

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
	Banks_GetBanks_FullMethodName   = "/worker.Banks/GetBanks"
	Banks_UpdateBank_FullMethodName = "/worker.Banks/UpdateBank"
)

// BanksClient is the client API for Banks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BanksClient interface {
	GetBanks(ctx context.Context, in *GetBanksRequest, opts ...grpc.CallOption) (*GetBanksResponse, error)
	UpdateBank(ctx context.Context, in *UpdateBankRequest, opts ...grpc.CallOption) (*UpdateBankResponse, error)
}

type banksClient struct {
	cc grpc.ClientConnInterface
}

func NewBanksClient(cc grpc.ClientConnInterface) BanksClient {
	return &banksClient{cc}
}

func (c *banksClient) GetBanks(ctx context.Context, in *GetBanksRequest, opts ...grpc.CallOption) (*GetBanksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBanksResponse)
	err := c.cc.Invoke(ctx, Banks_GetBanks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) UpdateBank(ctx context.Context, in *UpdateBankRequest, opts ...grpc.CallOption) (*UpdateBankResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBankResponse)
	err := c.cc.Invoke(ctx, Banks_UpdateBank_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BanksServer is the server API for Banks service.
// All implementations must embed UnimplementedBanksServer
// for forward compatibility.
type BanksServer interface {
	GetBanks(context.Context, *GetBanksRequest) (*GetBanksResponse, error)
	UpdateBank(context.Context, *UpdateBankRequest) (*UpdateBankResponse, error)
	mustEmbedUnimplementedBanksServer()
}

// UnimplementedBanksServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBanksServer struct{}

func (UnimplementedBanksServer) GetBanks(context.Context, *GetBanksRequest) (*GetBanksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanks not implemented")
}
func (UnimplementedBanksServer) UpdateBank(context.Context, *UpdateBankRequest) (*UpdateBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBank not implemented")
}
func (UnimplementedBanksServer) mustEmbedUnimplementedBanksServer() {}
func (UnimplementedBanksServer) testEmbeddedByValue()               {}

// UnsafeBanksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BanksServer will
// result in compilation errors.
type UnsafeBanksServer interface {
	mustEmbedUnimplementedBanksServer()
}

func RegisterBanksServer(s grpc.ServiceRegistrar, srv BanksServer) {
	// If the following call pancis, it indicates UnimplementedBanksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Banks_ServiceDesc, srv)
}

func _Banks_GetBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBanksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetBanks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetBanks(ctx, req.(*GetBanksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_UpdateBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).UpdateBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_UpdateBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).UpdateBank(ctx, req.(*UpdateBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Banks_ServiceDesc is the grpc.ServiceDesc for Banks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Banks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Banks",
	HandlerType: (*BanksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBanks",
			Handler:    _Banks_GetBanks_Handler,
		},
		{
			MethodName: "UpdateBank",
			Handler:    _Banks_UpdateBank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "banks.proto",
}