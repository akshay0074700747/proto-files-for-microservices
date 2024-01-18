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
	GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (CartService_GetCartClient, error)
	AddtoCart(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error)
	DeleteCartItem(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (CartService_DeleteCartItemClient, error)
	TruncateCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponce, error)
	ChangeQty(ctx context.Context, in *ChangeQtyRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error)
	TrasferWishlist(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (CartService_TrasferWishlistClient, error)
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

func (c *cartServiceClient) GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (CartService_GetCartClient, error) {
	stream, err := c.cc.NewStream(ctx, &CartService_ServiceDesc.Streams[0], CartService_GetCart_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cartServiceGetCartClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CartService_GetCartClient interface {
	Recv() (*AddtoCartResponce, error)
	grpc.ClientStream
}

type cartServiceGetCartClient struct {
	grpc.ClientStream
}

func (x *cartServiceGetCartClient) Recv() (*AddtoCartResponce, error) {
	m := new(AddtoCartResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cartServiceClient) AddtoCart(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (*AddtoCartResponce, error) {
	out := new(AddtoCartResponce)
	err := c.cc.Invoke(ctx, CartService_AddtoCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCartItem(ctx context.Context, in *AddtoCartRequest, opts ...grpc.CallOption) (CartService_DeleteCartItemClient, error) {
	stream, err := c.cc.NewStream(ctx, &CartService_ServiceDesc.Streams[1], CartService_DeleteCartItem_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cartServiceDeleteCartItemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CartService_DeleteCartItemClient interface {
	Recv() (*AddtoCartResponce, error)
	grpc.ClientStream
}

type cartServiceDeleteCartItemClient struct {
	grpc.ClientStream
}

func (x *cartServiceDeleteCartItemClient) Recv() (*AddtoCartResponce, error) {
	m := new(AddtoCartResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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

func (c *cartServiceClient) TrasferWishlist(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (CartService_TrasferWishlistClient, error) {
	stream, err := c.cc.NewStream(ctx, &CartService_ServiceDesc.Streams[2], CartService_TrasferWishlist_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cartServiceTrasferWishlistClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CartService_TrasferWishlistClient interface {
	Recv() (*AddtoCartResponce, error)
	grpc.ClientStream
}

type cartServiceTrasferWishlistClient struct {
	grpc.ClientStream
}

func (x *cartServiceTrasferWishlistClient) Recv() (*AddtoCartResponce, error) {
	m := new(AddtoCartResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	CreateCart(context.Context, *CartRequest) (*CartResponce, error)
	GetCart(*CartRequest, CartService_GetCartServer) error
	AddtoCart(context.Context, *AddtoCartRequest) (*AddtoCartResponce, error)
	DeleteCartItem(*AddtoCartRequest, CartService_DeleteCartItemServer) error
	TruncateCart(context.Context, *CartRequest) (*CartResponce, error)
	ChangeQty(context.Context, *ChangeQtyRequest) (*AddtoCartResponce, error)
	TrasferWishlist(*CartRequest, CartService_TrasferWishlistServer) error
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) CreateCart(context.Context, *CartRequest) (*CartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServiceServer) GetCart(*CartRequest, CartService_GetCartServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) AddtoCart(context.Context, *AddtoCartRequest) (*AddtoCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddtoCart not implemented")
}
func (UnimplementedCartServiceServer) DeleteCartItem(*AddtoCartRequest, CartService_DeleteCartItemServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
func (UnimplementedCartServiceServer) TruncateCart(context.Context, *CartRequest) (*CartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TruncateCart not implemented")
}
func (UnimplementedCartServiceServer) ChangeQty(context.Context, *ChangeQtyRequest) (*AddtoCartResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeQty not implemented")
}
func (UnimplementedCartServiceServer) TrasferWishlist(*CartRequest, CartService_TrasferWishlistServer) error {
	return status.Errorf(codes.Unimplemented, "method TrasferWishlist not implemented")
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

func _CartService_GetCart_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CartServiceServer).GetCart(m, &cartServiceGetCartServer{stream})
}

type CartService_GetCartServer interface {
	Send(*AddtoCartResponce) error
	grpc.ServerStream
}

type cartServiceGetCartServer struct {
	grpc.ServerStream
}

func (x *cartServiceGetCartServer) Send(m *AddtoCartResponce) error {
	return x.ServerStream.SendMsg(m)
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

func _CartService_DeleteCartItem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddtoCartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CartServiceServer).DeleteCartItem(m, &cartServiceDeleteCartItemServer{stream})
}

type CartService_DeleteCartItemServer interface {
	Send(*AddtoCartResponce) error
	grpc.ServerStream
}

type cartServiceDeleteCartItemServer struct {
	grpc.ServerStream
}

func (x *cartServiceDeleteCartItemServer) Send(m *AddtoCartResponce) error {
	return x.ServerStream.SendMsg(m)
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

func _CartService_TrasferWishlist_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CartServiceServer).TrasferWishlist(m, &cartServiceTrasferWishlistServer{stream})
}

type CartService_TrasferWishlistServer interface {
	Send(*AddtoCartResponce) error
	grpc.ServerStream
}

type cartServiceTrasferWishlistServer struct {
	grpc.ServerStream
}

func (x *cartServiceTrasferWishlistServer) Send(m *AddtoCartResponce) error {
	return x.ServerStream.SendMsg(m)
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
			MethodName: "AddtoCart",
			Handler:    _CartService_AddtoCart_Handler,
		},
		{
			MethodName: "TruncateCart",
			Handler:    _CartService_TruncateCart_Handler,
		},
		{
			MethodName: "ChangeQty",
			Handler:    _CartService_ChangeQty_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCart",
			Handler:       _CartService_GetCart_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeleteCartItem",
			Handler:       _CartService_DeleteCartItem_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TrasferWishlist",
			Handler:       _CartService_TrasferWishlist_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cart.proto",
}
