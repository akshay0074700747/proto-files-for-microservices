// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: cart.proto

package pb

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
	CartService_CreateCart_FullMethodName      = "/cart.CartService/CreateCart"
	CartService_GetCart_FullMethodName         = "/cart.CartService/GetCart"
	CartService_AddtoCart_FullMethodName       = "/cart.CartService/AddtoCart"
	CartService_DeleteCartItem_FullMethodName  = "/cart.CartService/DeleteCartItem"
	CartService_TruncateCart_FullMethodName    = "/cart.CartService/TruncateCart"
	CartService_ChangeQty_FullMethodName       = "/cart.CartService/ChangeQty"
	CartService_TrasferWishlist_FullMethodName = "/cart.CartService/TrasferWishlist"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	CreateCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponce, error)
	GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*GetCartResponce, error)
	AddtoCart(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error)
	DeleteCartItem(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (*GetCartResponce, error)
	TruncateCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponce, error)
	ChangeQty(ctx context.Context, in *ChangeQtyRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error)
	TrasferWishlist(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*GetCartResponce, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponce, error) {
	out := new(CartResponce)
	err := c.cc.Invoke(ctx, CartService_CreateCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*GetCartResponce, error) {
	out := new(GetCartResponce)
	err := c.cc.Invoke(ctx, CartService_GetCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddtoCart(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error) {
	out := new(AddtoCartResponce)
	err := c.cc.Invoke(ctx, CartService_AddtoCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCartItem(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (*GetCartResponce, error) {
	out := new(GetCartResponce)
	err := c.cc.Invoke(ctx, CartService_DeleteCartItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) TruncateCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponce, error) {
	out := new(CartResponce)
	err := c.cc.Invoke(ctx, CartService_TruncateCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ChangeQty(ctx context.Context, in *ChangeQtyRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error) {
	out := new(AddtoCartResponce)
	err := c.cc.Invoke(ctx, CartService_ChangeQty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) TrasferWishlist(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*GetCartResponce, error) {
	out := new(GetCartResponce)
	err := c.cc.Invoke(ctx, CartService_TrasferWishlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	CreateCart(context.Context, *CartRequest) (*CartResponce, error)
	GetCart(context.Context, *CartRequest) (*GetCartResponce, error)
	AddtoCart(context.Context, *AddtoCartRequest) (*AddtoCartResponce, error)
	DeleteCartItem(context.Context, *AddtoCartRequest) (*GetCartResponce, error)
	TruncateCart(context.Context, *CartRequest) (*CartResponce, error)
	ChangeQty(context.Context, *ChangeQtyRequest) (*AddtoCartResponce, error)
	TrasferWishlist(context.Context, *CartRequest) (*GetCartResponce, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) CreateCart(context.Context, *CartRequest) (*CartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServiceServer) GetCart(context.Context, *CartRequest) (*GetCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) AddtoCart(context.Context, *AddtoCartRequest) (*AddtoCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddtoCart not implemented")
}
func (UnimplementedCartServiceServer) DeleteCartItem(context.Context, *AddtoCartRequest) (*GetCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
func (UnimplementedCartServiceServer) TruncateCart(context.Context, *CartRequest) (*CartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TruncateCart not implemented")
}
func (UnimplementedCartServiceServer) ChangeQty(context.Context, *ChangeQtyRequest) (*AddtoCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeQty not implemented")
}
func (UnimplementedCartServiceServer) TrasferWishlist(context.Context, *CartRequest) (*GetCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrasferWishlist not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_CreateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddtoCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddtoCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddtoCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddtoCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddtoCart(ctx, req.(*AddtoCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddtoCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_DeleteCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteCartItem(ctx, req.(*AddtoCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_TruncateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).TruncateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_TruncateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).TruncateCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ChangeQty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeQtyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ChangeQty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_ChangeQty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ChangeQty(ctx, req.(*ChangeQtyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_TrasferWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).TrasferWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_TrasferWishlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).TrasferWishlist(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "AddtoCart",
			Handler:    _CartService_AddtoCart_Handler,
		},
		{
			MethodName: "DeleteCartItem",
			Handler:    _CartService_DeleteCartItem_Handler,
		},
		{
			MethodName: "TruncateCart",
			Handler:    _CartService_TruncateCart_Handler,
		},
		{
			MethodName: "ChangeQty",
			Handler:    _CartService_ChangeQty_Handler,
		},
		{
			MethodName: "TrasferWishlist",
			Handler:    _CartService_TrasferWishlist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
