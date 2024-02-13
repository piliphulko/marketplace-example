// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: service-acct-auth.proto

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
	AccountAut_CreateAccount_FullMethodName          = "/service_acct_aut.AccountAut/CreateAccount"
	AccountAut_CreateAccountWarehouse_FullMethodName = "/service_acct_aut.AccountAut/CreateAccountWarehouse"
	AccountAut_CreateAccountVendor_FullMethodName    = "/service_acct_aut.AccountAut/CreateAccountVendor"
	AccountAut_AutAccount_FullMethodName             = "/service_acct_aut.AccountAut/AutAccount"
	AccountAut_AutAccountWarehouse_FullMethodName    = "/service_acct_aut.AccountAut/AutAccountWarehouse"
	AccountAut_AutAccountVendor_FullMethodName       = "/service_acct_aut.AccountAut/AutAccountVendor"
	AccountAut_ChangeAccount_FullMethodName          = "/service_acct_aut.AccountAut/ChangeAccount"
	AccountAut_ChangeAccountWarehouse_FullMethodName = "/service_acct_aut.AccountAut/ChangeAccountWarehouse"
	AccountAut_ChangeAccountVendor_FullMethodName    = "/service_acct_aut.AccountAut/ChangeAccountVendor"
	AccountAut_CheckJWT_FullMethodName               = "/service_acct_aut.AccountAut/CheckJWT"
	AccountAut_GetCountryCity_FullMethodName         = "/service_acct_aut.AccountAut/GetCountryCity"
)

// AccountAutClient is the client API for AccountAut service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountAutClient interface {
	CreateAccount(ctx context.Context, in *basic.CustomerNew, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateAccountWarehouse(ctx context.Context, in *basic.WarehouseNew, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateAccountVendor(ctx context.Context, in *basic.VendorNew, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AutAccount(ctx context.Context, in *basic.CustomerAut, opts ...grpc.CallOption) (*basic.StringJWT, error)
	AutAccountWarehouse(ctx context.Context, in *basic.WarehouseAut, opts ...grpc.CallOption) (*basic.StringJWT, error)
	AutAccountVendor(ctx context.Context, in *basic.VendorAut, opts ...grpc.CallOption) (*basic.StringJWT, error)
	ChangeAccount(ctx context.Context, in *basic.CustomerChange, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangeAccountWarehouse(ctx context.Context, in *basic.WarehouseChange, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangeAccountVendor(ctx context.Context, in *basic.VendorChange, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckJWT(ctx context.Context, in *basic.StringJWT, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCountryCity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.CountryCityPairs, error)
}

type accountAutClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountAutClient(cc grpc.ClientConnInterface) AccountAutClient {
	return &accountAutClient{cc}
}

func (c *accountAutClient) CreateAccount(ctx context.Context, in *basic.CustomerNew, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) CreateAccountWarehouse(ctx context.Context, in *basic.WarehouseNew, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_CreateAccountWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) CreateAccountVendor(ctx context.Context, in *basic.VendorNew, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_CreateAccountVendor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) AutAccount(ctx context.Context, in *basic.CustomerAut, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	out := new(basic.StringJWT)
	err := c.cc.Invoke(ctx, AccountAut_AutAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) AutAccountWarehouse(ctx context.Context, in *basic.WarehouseAut, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	out := new(basic.StringJWT)
	err := c.cc.Invoke(ctx, AccountAut_AutAccountWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) AutAccountVendor(ctx context.Context, in *basic.VendorAut, opts ...grpc.CallOption) (*basic.StringJWT, error) {
	out := new(basic.StringJWT)
	err := c.cc.Invoke(ctx, AccountAut_AutAccountVendor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) ChangeAccount(ctx context.Context, in *basic.CustomerChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_ChangeAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) ChangeAccountWarehouse(ctx context.Context, in *basic.WarehouseChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_ChangeAccountWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) ChangeAccountVendor(ctx context.Context, in *basic.VendorChange, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_ChangeAccountVendor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) CheckJWT(ctx context.Context, in *basic.StringJWT, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccountAut_CheckJWT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAutClient) GetCountryCity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*basic.CountryCityPairs, error) {
	out := new(basic.CountryCityPairs)
	err := c.cc.Invoke(ctx, AccountAut_GetCountryCity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountAutServer is the server API for AccountAut service.
// All implementations must embed UnimplementedAccountAutServer
// for forward compatibility
type AccountAutServer interface {
	CreateAccount(context.Context, *basic.CustomerNew) (*emptypb.Empty, error)
	CreateAccountWarehouse(context.Context, *basic.WarehouseNew) (*emptypb.Empty, error)
	CreateAccountVendor(context.Context, *basic.VendorNew) (*emptypb.Empty, error)
	AutAccount(context.Context, *basic.CustomerAut) (*basic.StringJWT, error)
	AutAccountWarehouse(context.Context, *basic.WarehouseAut) (*basic.StringJWT, error)
	AutAccountVendor(context.Context, *basic.VendorAut) (*basic.StringJWT, error)
	ChangeAccount(context.Context, *basic.CustomerChange) (*emptypb.Empty, error)
	ChangeAccountWarehouse(context.Context, *basic.WarehouseChange) (*emptypb.Empty, error)
	ChangeAccountVendor(context.Context, *basic.VendorChange) (*emptypb.Empty, error)
	CheckJWT(context.Context, *basic.StringJWT) (*emptypb.Empty, error)
	GetCountryCity(context.Context, *emptypb.Empty) (*basic.CountryCityPairs, error)
	mustEmbedUnimplementedAccountAutServer()
}

// UnimplementedAccountAutServer must be embedded to have forward compatible implementations.
type UnimplementedAccountAutServer struct {
}

func (UnimplementedAccountAutServer) CreateAccount(context.Context, *basic.CustomerNew) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountAutServer) CreateAccountWarehouse(context.Context, *basic.WarehouseNew) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountWarehouse not implemented")
}
func (UnimplementedAccountAutServer) CreateAccountVendor(context.Context, *basic.VendorNew) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountVendor not implemented")
}
func (UnimplementedAccountAutServer) AutAccount(context.Context, *basic.CustomerAut) (*basic.StringJWT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutAccount not implemented")
}
func (UnimplementedAccountAutServer) AutAccountWarehouse(context.Context, *basic.WarehouseAut) (*basic.StringJWT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutAccountWarehouse not implemented")
}
func (UnimplementedAccountAutServer) AutAccountVendor(context.Context, *basic.VendorAut) (*basic.StringJWT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutAccountVendor not implemented")
}
func (UnimplementedAccountAutServer) ChangeAccount(context.Context, *basic.CustomerChange) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAccount not implemented")
}
func (UnimplementedAccountAutServer) ChangeAccountWarehouse(context.Context, *basic.WarehouseChange) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAccountWarehouse not implemented")
}
func (UnimplementedAccountAutServer) ChangeAccountVendor(context.Context, *basic.VendorChange) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAccountVendor not implemented")
}
func (UnimplementedAccountAutServer) CheckJWT(context.Context, *basic.StringJWT) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckJWT not implemented")
}
func (UnimplementedAccountAutServer) GetCountryCity(context.Context, *emptypb.Empty) (*basic.CountryCityPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountryCity not implemented")
}
func (UnimplementedAccountAutServer) mustEmbedUnimplementedAccountAutServer() {}

// UnsafeAccountAutServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountAutServer will
// result in compilation errors.
type UnsafeAccountAutServer interface {
	mustEmbedUnimplementedAccountAutServer()
}

func RegisterAccountAutServer(s grpc.ServiceRegistrar, srv AccountAutServer) {
	s.RegisterService(&AccountAut_ServiceDesc, srv)
}

func _AccountAut_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.CustomerNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).CreateAccount(ctx, req.(*basic.CustomerNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_CreateAccountWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.WarehouseNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).CreateAccountWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_CreateAccountWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).CreateAccountWarehouse(ctx, req.(*basic.WarehouseNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_CreateAccountVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.VendorNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).CreateAccountVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_CreateAccountVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).CreateAccountVendor(ctx, req.(*basic.VendorNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_AutAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.CustomerAut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).AutAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_AutAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).AutAccount(ctx, req.(*basic.CustomerAut))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_AutAccountWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.WarehouseAut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).AutAccountWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_AutAccountWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).AutAccountWarehouse(ctx, req.(*basic.WarehouseAut))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_AutAccountVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.VendorAut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).AutAccountVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_AutAccountVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).AutAccountVendor(ctx, req.(*basic.VendorAut))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_ChangeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.CustomerChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).ChangeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_ChangeAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).ChangeAccount(ctx, req.(*basic.CustomerChange))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_ChangeAccountWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.WarehouseChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).ChangeAccountWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_ChangeAccountWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).ChangeAccountWarehouse(ctx, req.(*basic.WarehouseChange))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_ChangeAccountVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.VendorChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).ChangeAccountVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_ChangeAccountVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).ChangeAccountVendor(ctx, req.(*basic.VendorChange))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_CheckJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(basic.StringJWT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).CheckJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_CheckJWT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).CheckJWT(ctx, req.(*basic.StringJWT))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAut_GetCountryCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAutServer).GetCountryCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAut_GetCountryCity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAutServer).GetCountryCity(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountAut_ServiceDesc is the grpc.ServiceDesc for AccountAut service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountAut_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_acct_aut.AccountAut",
	HandlerType: (*AccountAutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountAut_CreateAccount_Handler,
		},
		{
			MethodName: "CreateAccountWarehouse",
			Handler:    _AccountAut_CreateAccountWarehouse_Handler,
		},
		{
			MethodName: "CreateAccountVendor",
			Handler:    _AccountAut_CreateAccountVendor_Handler,
		},
		{
			MethodName: "AutAccount",
			Handler:    _AccountAut_AutAccount_Handler,
		},
		{
			MethodName: "AutAccountWarehouse",
			Handler:    _AccountAut_AutAccountWarehouse_Handler,
		},
		{
			MethodName: "AutAccountVendor",
			Handler:    _AccountAut_AutAccountVendor_Handler,
		},
		{
			MethodName: "ChangeAccount",
			Handler:    _AccountAut_ChangeAccount_Handler,
		},
		{
			MethodName: "ChangeAccountWarehouse",
			Handler:    _AccountAut_ChangeAccountWarehouse_Handler,
		},
		{
			MethodName: "ChangeAccountVendor",
			Handler:    _AccountAut_ChangeAccountVendor_Handler,
		},
		{
			MethodName: "CheckJWT",
			Handler:    _AccountAut_CheckJWT_Handler,
		},
		{
			MethodName: "GetCountryCity",
			Handler:    _AccountAut_GetCountryCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-acct-auth.proto",
}
