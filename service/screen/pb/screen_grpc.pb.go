// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: pb/screen.proto

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
	Screen_InsertScreenProject_FullMethodName     = "/screen.Screen/InsertScreenProject"
	Screen_UpdateScreenProject_FullMethodName     = "/screen.Screen/UpdateScreenProject"
	Screen_DeleteScreenProject_FullMethodName     = "/screen.Screen/DeleteScreenProject"
	Screen_SelectScreenProjectById_FullMethodName = "/screen.Screen/SelectScreenProjectById"
	Screen_SelectScreenProjectList_FullMethodName = "/screen.Screen/SelectScreenProjectList"
)

// ScreenClient is the client API for Screen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScreenClient interface {
	// 创建大屏信息
	InsertScreenProject(ctx context.Context, in *AddScreenProjectReq, opts ...grpc.CallOption) (*AddScreenProjectResp, error)
	// 更新大屏信息
	UpdateScreenProject(ctx context.Context, in *UpdateScreenProjectReq, opts ...grpc.CallOption) (*UpdateScreenProjectResp, error)
	// 根据大屏信息ID删除
	DeleteScreenProject(ctx context.Context, in *DelScreenProjectReq, opts ...grpc.CallOption) (*DelScreenProjectResp, error)
	// 根据大屏信息ID获取详情
	SelectScreenProjectById(ctx context.Context, in *SelectScreenProjectByIdReq, opts ...grpc.CallOption) (*SelectScreenProjectByIdResp, error)
	// 大屏信息列表
	SelectScreenProjectList(ctx context.Context, in *SelectScreenProjectListReq, opts ...grpc.CallOption) (*SelectScreenProjectListResp, error)
}

type screenClient struct {
	cc grpc.ClientConnInterface
}

func NewScreenClient(cc grpc.ClientConnInterface) ScreenClient {
	return &screenClient{cc}
}

func (c *screenClient) InsertScreenProject(ctx context.Context, in *AddScreenProjectReq, opts ...grpc.CallOption) (*AddScreenProjectResp, error) {
	out := new(AddScreenProjectResp)
	err := c.cc.Invoke(ctx, Screen_InsertScreenProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screenClient) UpdateScreenProject(ctx context.Context, in *UpdateScreenProjectReq, opts ...grpc.CallOption) (*UpdateScreenProjectResp, error) {
	out := new(UpdateScreenProjectResp)
	err := c.cc.Invoke(ctx, Screen_UpdateScreenProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screenClient) DeleteScreenProject(ctx context.Context, in *DelScreenProjectReq, opts ...grpc.CallOption) (*DelScreenProjectResp, error) {
	out := new(DelScreenProjectResp)
	err := c.cc.Invoke(ctx, Screen_DeleteScreenProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screenClient) SelectScreenProjectById(ctx context.Context, in *SelectScreenProjectByIdReq, opts ...grpc.CallOption) (*SelectScreenProjectByIdResp, error) {
	out := new(SelectScreenProjectByIdResp)
	err := c.cc.Invoke(ctx, Screen_SelectScreenProjectById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screenClient) SelectScreenProjectList(ctx context.Context, in *SelectScreenProjectListReq, opts ...grpc.CallOption) (*SelectScreenProjectListResp, error) {
	out := new(SelectScreenProjectListResp)
	err := c.cc.Invoke(ctx, Screen_SelectScreenProjectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScreenServer is the server API for Screen service.
// All implementations must embed UnimplementedScreenServer
// for forward compatibility
type ScreenServer interface {
	// 创建大屏信息
	InsertScreenProject(context.Context, *AddScreenProjectReq) (*AddScreenProjectResp, error)
	// 更新大屏信息
	UpdateScreenProject(context.Context, *UpdateScreenProjectReq) (*UpdateScreenProjectResp, error)
	// 根据大屏信息ID删除
	DeleteScreenProject(context.Context, *DelScreenProjectReq) (*DelScreenProjectResp, error)
	// 根据大屏信息ID获取详情
	SelectScreenProjectById(context.Context, *SelectScreenProjectByIdReq) (*SelectScreenProjectByIdResp, error)
	// 大屏信息列表
	SelectScreenProjectList(context.Context, *SelectScreenProjectListReq) (*SelectScreenProjectListResp, error)
	mustEmbedUnimplementedScreenServer()
}

// UnimplementedScreenServer must be embedded to have forward compatible implementations.
type UnimplementedScreenServer struct {
}

func (UnimplementedScreenServer) InsertScreenProject(context.Context, *AddScreenProjectReq) (*AddScreenProjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertScreenProject not implemented")
}
func (UnimplementedScreenServer) UpdateScreenProject(context.Context, *UpdateScreenProjectReq) (*UpdateScreenProjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScreenProject not implemented")
}
func (UnimplementedScreenServer) DeleteScreenProject(context.Context, *DelScreenProjectReq) (*DelScreenProjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScreenProject not implemented")
}
func (UnimplementedScreenServer) SelectScreenProjectById(context.Context, *SelectScreenProjectByIdReq) (*SelectScreenProjectByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectScreenProjectById not implemented")
}
func (UnimplementedScreenServer) SelectScreenProjectList(context.Context, *SelectScreenProjectListReq) (*SelectScreenProjectListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectScreenProjectList not implemented")
}
func (UnimplementedScreenServer) mustEmbedUnimplementedScreenServer() {}

// UnsafeScreenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScreenServer will
// result in compilation errors.
type UnsafeScreenServer interface {
	mustEmbedUnimplementedScreenServer()
}

func RegisterScreenServer(s grpc.ServiceRegistrar, srv ScreenServer) {
	s.RegisterService(&Screen_ServiceDesc, srv)
}

func _Screen_InsertScreenProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddScreenProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServer).InsertScreenProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Screen_InsertScreenProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServer).InsertScreenProject(ctx, req.(*AddScreenProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Screen_UpdateScreenProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScreenProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServer).UpdateScreenProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Screen_UpdateScreenProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServer).UpdateScreenProject(ctx, req.(*UpdateScreenProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Screen_DeleteScreenProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelScreenProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServer).DeleteScreenProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Screen_DeleteScreenProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServer).DeleteScreenProject(ctx, req.(*DelScreenProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Screen_SelectScreenProjectById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectScreenProjectByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServer).SelectScreenProjectById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Screen_SelectScreenProjectById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServer).SelectScreenProjectById(ctx, req.(*SelectScreenProjectByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Screen_SelectScreenProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectScreenProjectListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenServer).SelectScreenProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Screen_SelectScreenProjectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenServer).SelectScreenProjectList(ctx, req.(*SelectScreenProjectListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Screen_ServiceDesc is the grpc.ServiceDesc for Screen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Screen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "screen.Screen",
	HandlerType: (*ScreenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertScreenProject",
			Handler:    _Screen_InsertScreenProject_Handler,
		},
		{
			MethodName: "UpdateScreenProject",
			Handler:    _Screen_UpdateScreenProject_Handler,
		},
		{
			MethodName: "DeleteScreenProject",
			Handler:    _Screen_DeleteScreenProject_Handler,
		},
		{
			MethodName: "SelectScreenProjectById",
			Handler:    _Screen_SelectScreenProjectById_Handler,
		},
		{
			MethodName: "SelectScreenProjectList",
			Handler:    _Screen_SelectScreenProjectList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/screen.proto",
}