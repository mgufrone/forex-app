// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rate_service

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

// RateServiceClient is the client API for RateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateServiceClient interface {
	// read
	GetAll(ctx context.Context, in *RateFilter, opts ...grpc.CallOption) (*RateData, error)
	Count(ctx context.Context, in *RateFilter, opts ...grpc.CallOption) (*RateCountResult, error)
	GetAndCount(ctx context.Context, in *RateFilter, opts ...grpc.CallOption) (*RateCount, error)
	Latest(ctx context.Context, in *DateFilter, opts ...grpc.CallOption) (*RateData, error)
	History(ctx context.Context, in *SpanFilter, opts ...grpc.CallOption) (*RateData, error)
	// write
	Create(ctx context.Context, in *Rate, opts ...grpc.CallOption) (*Rate, error)
	Update(ctx context.Context, in *Rate, opts ...grpc.CallOption) (*Rate, error)
	Delete(ctx context.Context, in *Rate, opts ...grpc.CallOption) (*RateResult, error)
}

type rateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateServiceClient(cc grpc.ClientConnInterface) RateServiceClient {
	return &rateServiceClient{cc}
}

func (c *rateServiceClient) GetAll(ctx context.Context, in *RateFilter, opts ...grpc.CallOption) (*RateData, error) {
	out := new(RateData)
	err := c.cc.Invoke(ctx, "/service.RateService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) Count(ctx context.Context, in *RateFilter, opts ...grpc.CallOption) (*RateCountResult, error) {
	out := new(RateCountResult)
	err := c.cc.Invoke(ctx, "/service.RateService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) GetAndCount(ctx context.Context, in *RateFilter, opts ...grpc.CallOption) (*RateCount, error) {
	out := new(RateCount)
	err := c.cc.Invoke(ctx, "/service.RateService/GetAndCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) Latest(ctx context.Context, in *DateFilter, opts ...grpc.CallOption) (*RateData, error) {
	out := new(RateData)
	err := c.cc.Invoke(ctx, "/service.RateService/Latest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) History(ctx context.Context, in *SpanFilter, opts ...grpc.CallOption) (*RateData, error) {
	out := new(RateData)
	err := c.cc.Invoke(ctx, "/service.RateService/History", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) Create(ctx context.Context, in *Rate, opts ...grpc.CallOption) (*Rate, error) {
	out := new(Rate)
	err := c.cc.Invoke(ctx, "/service.RateService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) Update(ctx context.Context, in *Rate, opts ...grpc.CallOption) (*Rate, error) {
	out := new(Rate)
	err := c.cc.Invoke(ctx, "/service.RateService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) Delete(ctx context.Context, in *Rate, opts ...grpc.CallOption) (*RateResult, error) {
	out := new(RateResult)
	err := c.cc.Invoke(ctx, "/service.RateService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServiceServer is the server API for RateService service.
// All implementations must embed UnimplementedRateServiceServer
// for forward compatibility
type RateServiceServer interface {
	// read
	GetAll(context.Context, *RateFilter) (*RateData, error)
	Count(context.Context, *RateFilter) (*RateCountResult, error)
	GetAndCount(context.Context, *RateFilter) (*RateCount, error)
	Latest(context.Context, *DateFilter) (*RateData, error)
	History(context.Context, *SpanFilter) (*RateData, error)
	// write
	Create(context.Context, *Rate) (*Rate, error)
	Update(context.Context, *Rate) (*Rate, error)
	Delete(context.Context, *Rate) (*RateResult, error)
	mustEmbedUnimplementedRateServiceServer()
}

// UnimplementedRateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRateServiceServer struct {
}

func (UnimplementedRateServiceServer) GetAll(context.Context, *RateFilter) (*RateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRateServiceServer) Count(context.Context, *RateFilter) (*RateCountResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedRateServiceServer) GetAndCount(context.Context, *RateFilter) (*RateCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAndCount not implemented")
}
func (UnimplementedRateServiceServer) Latest(context.Context, *DateFilter) (*RateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Latest not implemented")
}
func (UnimplementedRateServiceServer) History(context.Context, *SpanFilter) (*RateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedRateServiceServer) Create(context.Context, *Rate) (*Rate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRateServiceServer) Update(context.Context, *Rate) (*Rate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRateServiceServer) Delete(context.Context, *Rate) (*RateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRateServiceServer) mustEmbedUnimplementedRateServiceServer() {}

// UnsafeRateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateServiceServer will
// result in compilation errors.
type UnsafeRateServiceServer interface {
	mustEmbedUnimplementedRateServiceServer()
}

func RegisterRateServiceServer(s grpc.ServiceRegistrar, srv RateServiceServer) {
	s.RegisterService(&RateService_ServiceDesc, srv)
}

func _RateService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetAll(ctx, req.(*RateFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).Count(ctx, req.(*RateFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_GetAndCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetAndCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/GetAndCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetAndCount(ctx, req.(*RateFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_Latest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).Latest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/Latest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).Latest(ctx, req.(*DateFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/History",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).History(ctx, req.(*SpanFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).Create(ctx, req.(*Rate))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).Update(ctx, req.(*Rate))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RateService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).Delete(ctx, req.(*Rate))
	}
	return interceptor(ctx, in, info, handler)
}

// RateService_ServiceDesc is the grpc.ServiceDesc for RateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.RateService",
	HandlerType: (*RateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _RateService_GetAll_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _RateService_Count_Handler,
		},
		{
			MethodName: "GetAndCount",
			Handler:    _RateService_GetAndCount_Handler,
		},
		{
			MethodName: "Latest",
			Handler:    _RateService_Latest_Handler,
		},
		{
			MethodName: "History",
			Handler:    _RateService_History_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RateService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RateService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RateService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
