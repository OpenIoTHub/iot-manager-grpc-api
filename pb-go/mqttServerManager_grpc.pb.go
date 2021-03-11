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

// MqttServerManagerClient is the client API for MqttServerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MqttServerManagerClient interface {
	//    对网关的操作
	AddMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error)
	DelMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error)
	UpdateMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error)
	//    rpc QueryMqttServer (MqttServerInfo) returns (OperationResponse) {}
	GetAllMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error)
}

type mqttServerManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewMqttServerManagerClient(cc grpc.ClientConnInterface) MqttServerManagerClient {
	return &mqttServerManagerClient{cc}
}

func (c *mqttServerManagerClient) AddMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.MqttServerManager/AddMqttServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttServerManagerClient) DelMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.MqttServerManager/DelMqttServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttServerManagerClient) UpdateMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.MqttServerManager/UpdateMqttServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttServerManagerClient) GetAllMqttServer(ctx context.Context, in *MqttServerInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.MqttServerManager/GetAllMqttServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqttServerManagerServer is the server API for MqttServerManager service.
// All implementations must embed UnimplementedMqttServerManagerServer
// for forward compatibility
type MqttServerManagerServer interface {
	//    对网关的操作
	AddMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error)
	DelMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error)
	UpdateMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error)
	//    rpc QueryMqttServer (MqttServerInfo) returns (OperationResponse) {}
	GetAllMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error)
	mustEmbedUnimplementedMqttServerManagerServer()
}

// UnimplementedMqttServerManagerServer must be embedded to have forward compatible implementations.
type UnimplementedMqttServerManagerServer struct {
}

func (UnimplementedMqttServerManagerServer) AddMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMqttServer not implemented")
}
func (UnimplementedMqttServerManagerServer) DelMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMqttServer not implemented")
}
func (UnimplementedMqttServerManagerServer) UpdateMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMqttServer not implemented")
}
func (UnimplementedMqttServerManagerServer) GetAllMqttServer(context.Context, *MqttServerInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMqttServer not implemented")
}
func (UnimplementedMqttServerManagerServer) mustEmbedUnimplementedMqttServerManagerServer() {}

// UnsafeMqttServerManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MqttServerManagerServer will
// result in compilation errors.
type UnsafeMqttServerManagerServer interface {
	mustEmbedUnimplementedMqttServerManagerServer()
}

func RegisterMqttServerManagerServer(s grpc.ServiceRegistrar, srv MqttServerManagerServer) {
	s.RegisterService(&_MqttServerManager_serviceDesc, srv)
}

func _MqttServerManager_AddMqttServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttServerManagerServer).AddMqttServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MqttServerManager/AddMqttServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttServerManagerServer).AddMqttServer(ctx, req.(*MqttServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttServerManager_DelMqttServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttServerManagerServer).DelMqttServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MqttServerManager/DelMqttServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttServerManagerServer).DelMqttServer(ctx, req.(*MqttServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttServerManager_UpdateMqttServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttServerManagerServer).UpdateMqttServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MqttServerManager/UpdateMqttServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttServerManagerServer).UpdateMqttServer(ctx, req.(*MqttServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttServerManager_GetAllMqttServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttServerManagerServer).GetAllMqttServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MqttServerManager/GetAllMqttServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttServerManagerServer).GetAllMqttServer(ctx, req.(*MqttServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _MqttServerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MqttServerManager",
	HandlerType: (*MqttServerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMqttServer",
			Handler:    _MqttServerManager_AddMqttServer_Handler,
		},
		{
			MethodName: "DelMqttServer",
			Handler:    _MqttServerManager_DelMqttServer_Handler,
		},
		{
			MethodName: "UpdateMqttServer",
			Handler:    _MqttServerManager_UpdateMqttServer_Handler,
		},
		{
			MethodName: "GetAllMqttServer",
			Handler:    _MqttServerManager_GetAllMqttServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mqttServerManager.proto",
}