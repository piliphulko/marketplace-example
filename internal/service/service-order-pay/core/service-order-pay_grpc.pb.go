// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: service-order-pay.proto

package core

import (
	context "context"
	basic "github.com/piliphulko/marketplace-example/api/basic"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrderPay_GetMarketplace_FullMethodName = "/OrderPay/GetMarketplace"
	OrderPay_CreateOrder_FullMethodName    = "/OrderPay/CreateOrder"
	OrderPay_AddToOrder_FullMethodName     = "/OrderPay/AddToOrder"
	OrderPay_ConfirmOrder_FullMethodName   = "/OrderPay/ConfirmOrder"
	OrderPay_CancelOrder_FullMethodName    = "/OrderPay/CancelOrder"
	OrderPay_CompleteOrder_FullMethodName  = "/OrderPay/CompleteOrder"
)

// OrderPayClient is the client API for OrderPay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderPayClient interface {
	GetMarketplace(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.GoodsArray, error)
	CreateOrder(ctx context.Context, in *basic.NewOrderARRAY, opts ...grpc.CallOption) (*basic.OrderUuid, error)
	AddToOrder(ctx context.Context, in *basic.AddToOrderARRAY, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConfirmOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CompleteOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderPayClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderPayClient(cc grpc.ClientConnInterface) OrderPayClient {
	return &orderPayClient{cc}
}

func (c *orderPayClient) GetMarketplace(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.GoodsArray, error) {
	out := new(basic.GoodsArray)
	err := c.cc.Invoke(ctx, OrderPay_GetMarketplace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPayClient) CreateOrder(ctx context.Context, in *basic.NewOrderARRAY, opts ...grpc.CallOption) (*basic.OrderUuid, error) {
	out := new(basic.OrderUuid)
	err := c.cc.Invoke(ctx, OrderPay_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPayClient) AddToOrder(ctx context.Context, in *basic.AddToOrderARRAY, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OrderPay_AddToOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPayClient) ConfirmOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OrderPay_ConfirmOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPayClient) CancelOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OrderPay_CancelOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPayClient) CompleteOrder(ctx context.Context, in *basic.OrderUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OrderPay_CompleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderPayServer is the server API for OrderPay service.
// All implementations must embed UnimplementedOrderPayServer
// for forward compatibility
type OrderPayServer interface {
	GetMarketplace(context.Context, *emptypb.Empty) (*basic.GoodsArray, error)
	CreateOrder(context.Context, *basic.NewOrderARRAY) (*basic.OrderUuid, error)
	AddToOrder(context.Context, *basic.AddToOrderARRAY) (*emptypb.Empty, error)
	ConfirmOrder(context.Context, *basic.OrderUuid) (*emptypb.Empty, error)
	CancelOrder(context.Context, *basic.OrderUuid) (*emptypb.Empty, error)
	CompleteOrder(context.Context, *basic.OrderUuid) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderPayServer()
}

// UnimplementedOrderPayServer must be embedded to have forward compatible implementations.
type UnimplementedOrderPayServer struct {
}

func (UnimplementedOrderPayServer) GetMarketplace(context.Context, *emptypb.Empty) (*basic.GoodsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketplace not implemented")
}
func (UnimplementedOrderPayServer) CreateOrder(context.Context, *basic.NewOrderARRAY) (*basic.OrderUuid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderPayServer) AddToOrder(context.Context, *basic.AddToOrderARRAY) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToOrder not implemented")
}
func (UnimplementedOrderPayServer) ConfirmOrder(context.Context, *basic.OrderUuid) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOrder not implemented")
}
func (UnimplementedOrderPayServer) CancelOrder(context.Context, *basic.OrderUuid) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderPayServer) CompleteOrder(context.Context, *basic.OrderUuid) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (UnimplementedOrderPayServer) mustEmbedUnimplementedOrderPayServer() {}

// UnsafeOrderPayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderPayServer will
// result in compilation errors.
type UnsafeOrderPayServer interface {
	mustEmbedUnimplementedOrderPayServer()
}

func RegisterOrderPayServer(s grpc.ServiceRegistrar, srv OrderPayServer) {
	s.RegisterService(&OrderPay_ServiceDesc, srv)
}

func _OrderPay_GetMarketplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPayServer).GetMarketplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderPay_GetMarketplace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPayServer).GetMarketplace(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPay_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.NewOrderARRAY)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPayServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderPay_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPayServer).CreateOrder(ctx, req.(*basic.NewOrderARRAY))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPay_AddToOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.AddToOrderARRAY)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPayServer).AddToOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderPay_AddToOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPayServer).AddToOrder(ctx, req.(*basic.AddToOrderARRAY))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPay_ConfirmOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.OrderUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPayServer).ConfirmOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderPay_ConfirmOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPayServer).ConfirmOrder(ctx, req.(*basic.OrderUuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPay_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.OrderUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPayServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderPay_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPayServer).CancelOrder(ctx, req.(*basic.OrderUuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPay_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.OrderUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPayServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderPay_CompleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPayServer).CompleteOrder(ctx, req.(*basic.OrderUuid))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderPay_ServiceDesc is the grpc.ServiceDesc for OrderPay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderPay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderPay",
	HandlerType: (*OrderPayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMarketplace",
			Handler:    _OrderPay_GetMarketplace_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderPay_CreateOrder_Handler,
		},
		{
			MethodName: "AddToOrder",
			Handler:    _OrderPay_AddToOrder_Handler,
		},
		{
			MethodName: "ConfirmOrder",
			Handler:    _OrderPay_ConfirmOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderPay_CancelOrder_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _OrderPay_CompleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-order-pay.proto",
}