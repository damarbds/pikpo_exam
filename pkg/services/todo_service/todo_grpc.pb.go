// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: pkg/services/todo_service/todo.proto

package todo_service

import (
	context "context"
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
	ToDoService_Create_FullMethodName   = "/todo_service.ToDoService/Create"
	ToDoService_Update_FullMethodName   = "/todo_service.ToDoService/Update"
	ToDoService_ReadById_FullMethodName = "/todo_service.ToDoService/ReadById"
	ToDoService_ReadAll_FullMethodName  = "/todo_service.ToDoService/ReadAll"
)

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToDoServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ReadByIdResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ReadByIdResponse, error)
	ReadById(ctx context.Context, in *ReadByIdRequest, opts ...grpc.CallOption) (*ReadByIdResponse, error)
	ReadAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ReadByIdResponse, error) {
	out := new(ReadByIdResponse)
	err := c.cc.Invoke(ctx, ToDoService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ReadByIdResponse, error) {
	out := new(ReadByIdResponse)
	err := c.cc.Invoke(ctx, ToDoService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadById(ctx context.Context, in *ReadByIdRequest, opts ...grpc.CallOption) (*ReadByIdResponse, error) {
	out := new(ReadByIdResponse)
	err := c.cc.Invoke(ctx, ToDoService_ReadById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, ToDoService_ReadAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
// All implementations must embed UnimplementedToDoServiceServer
// for forward compatibility
type ToDoServiceServer interface {
	Create(context.Context, *CreateRequest) (*ReadByIdResponse, error)
	Update(context.Context, *UpdateRequest) (*ReadByIdResponse, error)
	ReadById(context.Context, *ReadByIdRequest) (*ReadByIdResponse, error)
	ReadAll(context.Context, *emptypb.Empty) (*ReadAllResponse, error)
	mustEmbedUnimplementedToDoServiceServer()
}

// UnimplementedToDoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (UnimplementedToDoServiceServer) Create(context.Context, *CreateRequest) (*ReadByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedToDoServiceServer) Update(context.Context, *UpdateRequest) (*ReadByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedToDoServiceServer) ReadById(context.Context, *ReadByIdRequest) (*ReadByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadById not implemented")
}
func (UnimplementedToDoServiceServer) ReadAll(context.Context, *emptypb.Empty) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}
func (UnimplementedToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {}

// UnsafeToDoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToDoServiceServer will
// result in compilation errors.
type UnsafeToDoServiceServer interface {
	mustEmbedUnimplementedToDoServiceServer()
}

func RegisterToDoServiceServer(s grpc.ServiceRegistrar, srv ToDoServiceServer) {
	s.RegisterService(&ToDoService_ServiceDesc, srv)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ReadById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_ReadById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadById(ctx, req.(*ReadByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ToDoService_ReadAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ToDoService_ServiceDesc is the grpc.ServiceDesc for ToDoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToDoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo_service.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "ReadById",
			Handler:    _ToDoService_ReadById_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/todo_service/todo.proto",
}
