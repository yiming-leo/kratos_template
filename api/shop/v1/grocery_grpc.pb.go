// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: api/shop/v1/grocery.proto

package v1

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

const (
	Grocery_CreateGrocery_FullMethodName = "/api.shop.v1.Grocery/CreateGrocery"
	Grocery_UpdateGrocery_FullMethodName = "/api.shop.v1.Grocery/UpdateGrocery"
	Grocery_DeleteGrocery_FullMethodName = "/api.shop.v1.Grocery/DeleteGrocery"
	Grocery_GetGrocery_FullMethodName    = "/api.shop.v1.Grocery/GetGrocery"
	Grocery_ListGrocery_FullMethodName   = "/api.shop.v1.Grocery/ListGrocery"
)

// GroceryClient is the client API for Grocery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroceryClient interface {
	CreateGrocery(ctx context.Context, in *CreateGroceryRequest, opts ...grpc.CallOption) (*CreateGroceryReply, error)
	UpdateGrocery(ctx context.Context, in *UpdateGroceryRequest, opts ...grpc.CallOption) (*UpdateGroceryReply, error)
	DeleteGrocery(ctx context.Context, in *DeleteGroceryRequest, opts ...grpc.CallOption) (*DeleteGroceryReply, error)
	GetGrocery(ctx context.Context, in *GetGroceryRequest, opts ...grpc.CallOption) (*GetGroceryReply, error)
	ListGrocery(ctx context.Context, in *ListGroceryRequest, opts ...grpc.CallOption) (*ListGroceryReply, error)
}

type groceryClient struct {
	cc grpc.ClientConnInterface
}

func NewGroceryClient(cc grpc.ClientConnInterface) GroceryClient {
	return &groceryClient{cc}
}

func (c *groceryClient) CreateGrocery(ctx context.Context, in *CreateGroceryRequest, opts ...grpc.CallOption) (*CreateGroceryReply, error) {
	out := new(CreateGroceryReply)
	err := c.cc.Invoke(ctx, Grocery_CreateGrocery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groceryClient) UpdateGrocery(ctx context.Context, in *UpdateGroceryRequest, opts ...grpc.CallOption) (*UpdateGroceryReply, error) {
	out := new(UpdateGroceryReply)
	err := c.cc.Invoke(ctx, Grocery_UpdateGrocery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groceryClient) DeleteGrocery(ctx context.Context, in *DeleteGroceryRequest, opts ...grpc.CallOption) (*DeleteGroceryReply, error) {
	out := new(DeleteGroceryReply)
	err := c.cc.Invoke(ctx, Grocery_DeleteGrocery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groceryClient) GetGrocery(ctx context.Context, in *GetGroceryRequest, opts ...grpc.CallOption) (*GetGroceryReply, error) {
	out := new(GetGroceryReply)
	err := c.cc.Invoke(ctx, Grocery_GetGrocery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groceryClient) ListGrocery(ctx context.Context, in *ListGroceryRequest, opts ...grpc.CallOption) (*ListGroceryReply, error) {
	out := new(ListGroceryReply)
	err := c.cc.Invoke(ctx, Grocery_ListGrocery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroceryServer is the server API for Grocery service.
// All implementations must embed UnimplementedGroceryServer
// for forward compatibility
type GroceryServer interface {
	CreateGrocery(context.Context, *CreateGroceryRequest) (*CreateGroceryReply, error)
	UpdateGrocery(context.Context, *UpdateGroceryRequest) (*UpdateGroceryReply, error)
	DeleteGrocery(context.Context, *DeleteGroceryRequest) (*DeleteGroceryReply, error)
	GetGrocery(context.Context, *GetGroceryRequest) (*GetGroceryReply, error)
	ListGrocery(context.Context, *ListGroceryRequest) (*ListGroceryReply, error)
	mustEmbedUnimplementedGroceryServer()
}

// UnimplementedGroceryServer must be embedded to have forward compatible implementations.
type UnimplementedGroceryServer struct {
}

func (UnimplementedGroceryServer) CreateGrocery(context.Context, *CreateGroceryRequest) (*CreateGroceryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGrocery not implemented")
}
func (UnimplementedGroceryServer) UpdateGrocery(context.Context, *UpdateGroceryRequest) (*UpdateGroceryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGrocery not implemented")
}
func (UnimplementedGroceryServer) DeleteGrocery(context.Context, *DeleteGroceryRequest) (*DeleteGroceryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGrocery not implemented")
}
func (UnimplementedGroceryServer) GetGrocery(context.Context, *GetGroceryRequest) (*GetGroceryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGrocery not implemented")
}
func (UnimplementedGroceryServer) ListGrocery(context.Context, *ListGroceryRequest) (*ListGroceryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGrocery not implemented")
}
func (UnimplementedGroceryServer) mustEmbedUnimplementedGroceryServer() {}

// UnsafeGroceryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroceryServer will
// result in compilation errors.
type UnsafeGroceryServer interface {
	mustEmbedUnimplementedGroceryServer()
}

func RegisterGroceryServer(s grpc.ServiceRegistrar, srv GroceryServer) {
	s.RegisterService(&Grocery_ServiceDesc, srv)
}

func _Grocery_CreateGrocery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroceryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroceryServer).CreateGrocery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grocery_CreateGrocery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroceryServer).CreateGrocery(ctx, req.(*CreateGroceryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grocery_UpdateGrocery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroceryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroceryServer).UpdateGrocery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grocery_UpdateGrocery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroceryServer).UpdateGrocery(ctx, req.(*UpdateGroceryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grocery_DeleteGrocery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroceryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroceryServer).DeleteGrocery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grocery_DeleteGrocery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroceryServer).DeleteGrocery(ctx, req.(*DeleteGroceryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grocery_GetGrocery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroceryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroceryServer).GetGrocery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grocery_GetGrocery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroceryServer).GetGrocery(ctx, req.(*GetGroceryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grocery_ListGrocery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroceryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroceryServer).ListGrocery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grocery_ListGrocery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroceryServer).ListGrocery(ctx, req.(*ListGroceryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Grocery_ServiceDesc is the grpc.ServiceDesc for Grocery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grocery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.shop.v1.Grocery",
	HandlerType: (*GroceryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGrocery",
			Handler:    _Grocery_CreateGrocery_Handler,
		},
		{
			MethodName: "UpdateGrocery",
			Handler:    _Grocery_UpdateGrocery_Handler,
		},
		{
			MethodName: "DeleteGrocery",
			Handler:    _Grocery_DeleteGrocery_Handler,
		},
		{
			MethodName: "GetGrocery",
			Handler:    _Grocery_GetGrocery_Handler,
		},
		{
			MethodName: "ListGrocery",
			Handler:    _Grocery_ListGrocery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/shop/v1/grocery.proto",
}
