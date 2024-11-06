// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: tech/storezhang/dnsabre/message/domain.proto

package message

import (
	context "context"
	domain "github.com/dnsabre/core/domain"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Domain_Add_FullMethodName    = "/tech.storezhang.dnsabre.message.Domain/Add"
	Domain_Get_FullMethodName    = "/tech.storezhang.dnsabre.message.Domain/Get"
	Domain_Paging_FullMethodName = "/tech.storezhang.dnsabre.message.Domain/Paging"
	Domain_Update_FullMethodName = "/tech.storezhang.dnsabre.message.Domain/Update"
	Domain_Delete_FullMethodName = "/tech.storezhang.dnsabre.message.Domain/Delete"
)

// DomainClient is the client API for Domain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainClient interface {
	// 创建
	Add(ctx context.Context, in *domain.AddReq, opts ...grpc.CallOption) (*domain.AddRsp, error)
	// 获取
	Get(ctx context.Context, in *domain.GetReq, opts ...grpc.CallOption) (*domain.GetRsp, error)
	// 分页
	Paging(ctx context.Context, in *domain.PagingReq, opts ...grpc.CallOption) (*domain.PagingRsp, error)
	// 修改
	Update(ctx context.Context, in *domain.UpdateReq, opts ...grpc.CallOption) (*domain.UpdateRsp, error)
	// 删除
	Delete(ctx context.Context, in *domain.DeleteReq, opts ...grpc.CallOption) (*domain.DeleteRsp, error)
}

type domainClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainClient(cc grpc.ClientConnInterface) DomainClient {
	return &domainClient{cc}
}

func (c *domainClient) Add(ctx context.Context, in *domain.AddReq, opts ...grpc.CallOption) (*domain.AddRsp, error) {
	out := new(domain.AddRsp)
	err := c.cc.Invoke(ctx, Domain_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) Get(ctx context.Context, in *domain.GetReq, opts ...grpc.CallOption) (*domain.GetRsp, error) {
	out := new(domain.GetRsp)
	err := c.cc.Invoke(ctx, Domain_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) Paging(ctx context.Context, in *domain.PagingReq, opts ...grpc.CallOption) (*domain.PagingRsp, error) {
	out := new(domain.PagingRsp)
	err := c.cc.Invoke(ctx, Domain_Paging_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) Update(ctx context.Context, in *domain.UpdateReq, opts ...grpc.CallOption) (*domain.UpdateRsp, error) {
	out := new(domain.UpdateRsp)
	err := c.cc.Invoke(ctx, Domain_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) Delete(ctx context.Context, in *domain.DeleteReq, opts ...grpc.CallOption) (*domain.DeleteRsp, error) {
	out := new(domain.DeleteRsp)
	err := c.cc.Invoke(ctx, Domain_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainServer is the server API for Domain service.
// All implementations must embed UnimplementedDomainServer
// for forward compatibility
type DomainServer interface {
	// 创建
	Add(context.Context, *domain.AddReq) (*domain.AddRsp, error)
	// 获取
	Get(context.Context, *domain.GetReq) (*domain.GetRsp, error)
	// 分页
	Paging(context.Context, *domain.PagingReq) (*domain.PagingRsp, error)
	// 修改
	Update(context.Context, *domain.UpdateReq) (*domain.UpdateRsp, error)
	// 删除
	Delete(context.Context, *domain.DeleteReq) (*domain.DeleteRsp, error)
	mustEmbedUnimplementedDomainServer()
}

// UnimplementedDomainServer must be embedded to have forward compatible implementations.
type UnimplementedDomainServer struct {
}

func (UnimplementedDomainServer) Add(context.Context, *domain.AddReq) (*domain.AddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDomainServer) Get(context.Context, *domain.GetReq) (*domain.GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDomainServer) Paging(context.Context, *domain.PagingReq) (*domain.PagingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paging not implemented")
}
func (UnimplementedDomainServer) Update(context.Context, *domain.UpdateReq) (*domain.UpdateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDomainServer) Delete(context.Context, *domain.DeleteReq) (*domain.DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDomainServer) mustEmbedUnimplementedDomainServer() {}

// UnsafeDomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainServer will
// result in compilation errors.
type UnsafeDomainServer interface {
	mustEmbedUnimplementedDomainServer()
}

func RegisterDomainServer(s grpc.ServiceRegistrar, srv DomainServer) {
	s.RegisterService(&Domain_ServiceDesc, srv)
}

func _Domain_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Domain_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).Add(ctx, req.(*domain.AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Domain_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).Get(ctx, req.(*domain.GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_Paging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).Paging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Domain_Paging_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).Paging(ctx, req.(*domain.PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Domain_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).Update(ctx, req.(*domain.UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Domain_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).Delete(ctx, req.(*domain.DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Domain_ServiceDesc is the grpc.ServiceDesc for Domain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Domain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tech.storezhang.dnsabre.message.Domain",
	HandlerType: (*DomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Domain_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Domain_Get_Handler,
		},
		{
			MethodName: "Paging",
			Handler:    _Domain_Paging_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Domain_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Domain_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tech/storezhang/dnsabre/message/domain.proto",
}
