// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserManagerClient is the client API for UserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagerClient interface {
	//    登录，对用户的操作
	LoginWithUserLoginInfo(ctx context.Context, in *UserLoginInfo, opts ...grpc.CallOption) (*OperationResponse, error)
}

type userManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagerClient(cc grpc.ClientConnInterface) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) LoginWithUserLoginInfo(ctx context.Context, in *UserLoginInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.UserManager/LoginWithUserLoginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServer is the server API for UserManager service.
// All implementations must embed UnimplementedUserManagerServer
// for forward compatibility
type UserManagerServer interface {
	//    登录，对用户的操作
	LoginWithUserLoginInfo(context.Context, *UserLoginInfo) (*OperationResponse, error)
	mustEmbedUnimplementedUserManagerServer()
}

// UnimplementedUserManagerServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagerServer struct {
}

func (UnimplementedUserManagerServer) LoginWithUserLoginInfo(context.Context, *UserLoginInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithUserLoginInfo not implemented")
}
func (UnimplementedUserManagerServer) mustEmbedUnimplementedUserManagerServer() {}

// UnsafeUserManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagerServer will
// result in compilation errors.
type UnsafeUserManagerServer interface {
	mustEmbedUnimplementedUserManagerServer()
}

func RegisterUserManagerServer(s grpc.ServiceRegistrar, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_LoginWithUserLoginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).LoginWithUserLoginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserManager/LoginWithUserLoginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).LoginWithUserLoginInfo(ctx, req.(*UserLoginInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginWithUserLoginInfo",
			Handler:    _UserManager_LoginWithUserLoginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userManager.proto",
}