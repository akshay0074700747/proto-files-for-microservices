// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_SignupUser_FullMethodName           = "/user.UserService/SignupUser"
	UserService_LoginUser_FullMethodName            = "/user.UserService/LoginUser"
	UserService_GetUser_FullMethodName              = "/user.UserService/GetUser"
	UserService_GetAdmin_FullMethodName             = "/user.UserService/GetAdmin"
	UserService_GetSuAdmin_FullMethodName           = "/user.UserService/GetSuAdmin"
	UserService_GetAllUsersResponce_FullMethodName  = "/user.UserService/GetAllUsersResponce"
	UserService_GetAllAdminsResponce_FullMethodName = "/user.UserService/GetAllAdminsResponce"
	UserService_AddAdmin_FullMethodName             = "/user.UserService/AddAdmin"
	UserService_Check_FullMethodName                = "/user.UserService/Check"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	SignupUser(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*UserResponce, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponce, error)
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponce, error)
	GetAdmin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponce, error)
	GetSuAdmin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponce, error)
	GetAllUsersResponce(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllUsersResponce, error)
	GetAllAdminsResponce(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllUsersResponce, error)
	AddAdmin(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*UserResponce, error)
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SignupUser(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_SignupUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAdmin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_GetAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSuAdmin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_GetSuAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsersResponce(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllUsersResponce, error) {
	out := new(AllUsersResponce)
	err := c.cc.Invoke(ctx, UserService_GetAllUsersResponce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllAdminsResponce(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllUsersResponce, error) {
	out := new(AllUsersResponce)
	err := c.cc.Invoke(ctx, UserService_GetAllAdminsResponce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAdmin(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_AddAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, UserService_Check_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	SignupUser(context.Context, *SignupUserRequest) (*UserResponce, error)
	LoginUser(context.Context, *LoginRequest) (*UserResponce, error)
	GetUser(context.Context, *UserRequest) (*UserResponce, error)
	GetAdmin(context.Context, *UserRequest) (*UserResponce, error)
	GetSuAdmin(context.Context, *UserRequest) (*UserResponce, error)
	GetAllUsersResponce(context.Context, *empty.Empty) (*AllUsersResponce, error)
	GetAllAdminsResponce(context.Context, *empty.Empty) (*AllUsersResponce, error)
	AddAdmin(context.Context, *SignupUserRequest) (*UserResponce, error)
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) SignupUser(context.Context, *SignupUserRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupUser not implemented")
}
func (UnimplementedUserServiceServer) LoginUser(context.Context, *LoginRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *UserRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetAdmin(context.Context, *UserRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdmin not implemented")
}
func (UnimplementedUserServiceServer) GetSuAdmin(context.Context, *UserRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuAdmin not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsersResponce(context.Context, *empty.Empty) (*AllUsersResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsersResponce not implemented")
}
func (UnimplementedUserServiceServer) GetAllAdminsResponce(context.Context, *empty.Empty) (*AllUsersResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAdminsResponce not implemented")
}
func (UnimplementedUserServiceServer) AddAdmin(context.Context, *SignupUserRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedUserServiceServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_SignupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SignupUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignupUser(ctx, req.(*SignupUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAdmin(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSuAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSuAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetSuAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSuAdmin(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsersResponce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsersResponce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsersResponce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsersResponce(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllAdminsResponce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllAdminsResponce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllAdminsResponce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllAdminsResponce(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAdmin(ctx, req.(*SignupUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignupUser",
			Handler:    _UserService_SignupUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserService_LoginUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetAdmin",
			Handler:    _UserService_GetAdmin_Handler,
		},
		{
			MethodName: "GetSuAdmin",
			Handler:    _UserService_GetSuAdmin_Handler,
		},
		{
			MethodName: "GetAllUsersResponce",
			Handler:    _UserService_GetAllUsersResponce_Handler,
		},
		{
			MethodName: "GetAllAdminsResponce",
			Handler:    _UserService_GetAllAdminsResponce_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _UserService_AddAdmin_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _UserService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
