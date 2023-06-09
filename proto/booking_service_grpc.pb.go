// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: booking_service.proto

package booking

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
	BookingService_CreateBooking_FullMethodName             = "/booking.BookingService/CreateBooking"
	BookingService_GetAll_FullMethodName                    = "/booking.BookingService/GetAll"
	BookingService_GetAllOnPending_FullMethodName           = "/booking.BookingService/GetAllOnPending"
	BookingService_Decline_FullMethodName                   = "/booking.BookingService/Decline"
	BookingService_Confirm_FullMethodName                   = "/booking.BookingService/Confirm"
	BookingService_GetAllReservations_FullMethodName        = "/booking.BookingService/GetAllReservations"
	BookingService_GuestHasReservationInPast_FullMethodName = "/booking.BookingService/GuestHasReservationInPast"
	BookingService_GetUserReservations_FullMethodName       = "/booking.BookingService/GetUserReservations"
	BookingService_GetFinishedReservations_FullMethodName   = "/booking.BookingService/GetFinishedReservations"
	BookingService_CanceledBooking_FullMethodName           = "/booking.BookingService/CanceledBooking"
)

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetAll(ctx context.Context, in *EmptyRequst, opts ...grpc.CallOption) (*BookingResponse, error)
	GetAllOnPending(ctx context.Context, in *GetAllPendingRequest, opts ...grpc.CallOption) (*BookingResponse, error)
	Decline(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	Confirm(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetAllReservations(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	GuestHasReservationInPast(ctx context.Context, in *GuestHasReservationInPastRequest, opts ...grpc.CallOption) (*GuestHasReservationInPastResponse, error)
	GetUserReservations(ctx context.Context, in *UserReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	GetFinishedReservations(ctx context.Context, in *UserReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	CanceledBooking(ctx context.Context, in *CanceledBookingRequest, opts ...grpc.CallOption) (*CanceledBookingResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_CreateBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAll(ctx context.Context, in *EmptyRequst, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, BookingService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllOnPending(ctx context.Context, in *GetAllPendingRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, BookingService_GetAllOnPending_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) Decline(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_Decline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) Confirm(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_Confirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllReservations(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, BookingService_GetAllReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GuestHasReservationInPast(ctx context.Context, in *GuestHasReservationInPastRequest, opts ...grpc.CallOption) (*GuestHasReservationInPastResponse, error) {
	out := new(GuestHasReservationInPastResponse)
	err := c.cc.Invoke(ctx, BookingService_GuestHasReservationInPast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetUserReservations(ctx context.Context, in *UserReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, BookingService_GetUserReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetFinishedReservations(ctx context.Context, in *UserReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, BookingService_GetFinishedReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CanceledBooking(ctx context.Context, in *CanceledBookingRequest, opts ...grpc.CallOption) (*CanceledBookingResponse, error) {
	out := new(CanceledBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_CanceledBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetAll(context.Context, *EmptyRequst) (*BookingResponse, error)
	GetAllOnPending(context.Context, *GetAllPendingRequest) (*BookingResponse, error)
	Decline(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	Confirm(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetAllReservations(context.Context, *ReservationRequest) (*ReservationResponse, error)
	GuestHasReservationInPast(context.Context, *GuestHasReservationInPastRequest) (*GuestHasReservationInPastResponse, error)
	GetUserReservations(context.Context, *UserReservationRequest) (*ReservationResponse, error)
	GetFinishedReservations(context.Context, *UserReservationRequest) (*ReservationResponse, error)
	CanceledBooking(context.Context, *CanceledBookingRequest) (*CanceledBookingResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetAll(context.Context, *EmptyRequst) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBookingServiceServer) GetAllOnPending(context.Context, *GetAllPendingRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOnPending not implemented")
}
func (UnimplementedBookingServiceServer) Decline(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decline not implemented")
}
func (UnimplementedBookingServiceServer) Confirm(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedBookingServiceServer) GetAllReservations(context.Context, *ReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservations not implemented")
}
func (UnimplementedBookingServiceServer) GuestHasReservationInPast(context.Context, *GuestHasReservationInPastRequest) (*GuestHasReservationInPastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestHasReservationInPast not implemented")
}
func (UnimplementedBookingServiceServer) GetUserReservations(context.Context, *UserReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserReservations not implemented")
}
func (UnimplementedBookingServiceServer) GetFinishedReservations(context.Context, *UserReservationRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinishedReservations not implemented")
}
func (UnimplementedBookingServiceServer) CanceledBooking(context.Context, *CanceledBookingRequest) (*CanceledBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanceledBooking not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAll(ctx, req.(*EmptyRequst))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllOnPending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllOnPending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetAllOnPending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllOnPending(ctx, req.(*GetAllPendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_Decline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).Decline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_Decline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).Decline(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_Confirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).Confirm(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetAllReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllReservations(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GuestHasReservationInPast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestHasReservationInPastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GuestHasReservationInPast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GuestHasReservationInPast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GuestHasReservationInPast(ctx, req.(*GuestHasReservationInPastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetUserReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetUserReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetUserReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetUserReservations(ctx, req.(*UserReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetFinishedReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetFinishedReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetFinishedReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetFinishedReservations(ctx, req.(*UserReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CanceledBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanceledBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CanceledBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CanceledBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CanceledBooking(ctx, req.(*CanceledBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BookingService_GetAll_Handler,
		},
		{
			MethodName: "GetAllOnPending",
			Handler:    _BookingService_GetAllOnPending_Handler,
		},
		{
			MethodName: "Decline",
			Handler:    _BookingService_Decline_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _BookingService_Confirm_Handler,
		},
		{
			MethodName: "GetAllReservations",
			Handler:    _BookingService_GetAllReservations_Handler,
		},
		{
			MethodName: "GuestHasReservationInPast",
			Handler:    _BookingService_GuestHasReservationInPast_Handler,
		},
		{
			MethodName: "GetUserReservations",
			Handler:    _BookingService_GetUserReservations_Handler,
		},
		{
			MethodName: "GetFinishedReservations",
			Handler:    _BookingService_GetFinishedReservations_Handler,
		},
		{
			MethodName: "CanceledBooking",
			Handler:    _BookingService_CanceledBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking_service.proto",
}
