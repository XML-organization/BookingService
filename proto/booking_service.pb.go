// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: booking_service.proto

package booking

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GuestHasReservationInPastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccomodationsId []string `protobuf:"bytes,1,rep,name=accomodationsId,proto3" json:"accomodationsId,omitempty"`
	GuestId         string   `protobuf:"bytes,2,opt,name=guestId,proto3" json:"guestId,omitempty"`
}

func (x *GuestHasReservationInPastRequest) Reset() {
	*x = GuestHasReservationInPastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestHasReservationInPastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestHasReservationInPastRequest) ProtoMessage() {}

func (x *GuestHasReservationInPastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestHasReservationInPastRequest.ProtoReflect.Descriptor instead.
func (*GuestHasReservationInPastRequest) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{0}
}

func (x *GuestHasReservationInPastRequest) GetAccomodationsId() []string {
	if x != nil {
		return x.AccomodationsId
	}
	return nil
}

func (x *GuestHasReservationInPastRequest) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

type GuestHasReservationInPastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GuestHasReservationInPastResponse) Reset() {
	*x = GuestHasReservationInPastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestHasReservationInPastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestHasReservationInPastResponse) ProtoMessage() {}

func (x *GuestHasReservationInPastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestHasReservationInPastResponse.ProtoReflect.Descriptor instead.
func (*GuestHasReservationInPastResponse) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{1}
}

func (x *GuestHasReservationInPastResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllPendingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId string `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *GetAllPendingRequest) Reset() {
	*x = GetAllPendingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPendingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPendingRequest) ProtoMessage() {}

func (x *GetAllPendingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPendingRequest.ProtoReflect.Descriptor instead.
func (*GetAllPendingRequest) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPendingRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

type CreateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccomodationID string `protobuf:"bytes,2,opt,name=accomodationID,proto3" json:"accomodationID,omitempty"`
	UserID         string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	StartDate      string `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate        string `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	GuestNumber    string `protobuf:"bytes,6,opt,name=guestNumber,proto3" json:"guestNumber,omitempty"`
	Status         string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingRequest.ProtoReflect.Descriptor instead.
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBookingRequest) GetAccomodationID() string {
	if x != nil {
		return x.AccomodationID
	}
	return ""
}

func (x *CreateBookingRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateBookingRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CreateBookingRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *CreateBookingRequest) GetGuestNumber() string {
	if x != nil {
		return x.GuestNumber
	}
	return ""
}

func (x *CreateBookingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type EmptyRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequst) Reset() {
	*x = EmptyRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequst) ProtoMessage() {}

func (x *EmptyRequst) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequst.ProtoReflect.Descriptor instead.
func (*EmptyRequst) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{4}
}

type CreateBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateBookingResponse) Reset() {
	*x = CreateBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingResponse) ProtoMessage() {}

func (x *CreateBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingResponse.ProtoReflect.Descriptor instead.
func (*CreateBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*CreateBookingRequest `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{6}
}

func (x *BookingResponse) GetBookings() []*CreateBookingRequest {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type ReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccomodationId string `protobuf:"bytes,1,opt,name=accomodationId,proto3" json:"accomodationId,omitempty"`
}

func (x *ReservationRequest) Reset() {
	*x = ReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationRequest) ProtoMessage() {}

func (x *ReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationRequest.ProtoReflect.Descriptor instead.
func (*ReservationRequest) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{7}
}

func (x *ReservationRequest) GetAccomodationId() string {
	if x != nil {
		return x.AccomodationId
	}
	return ""
}

type UserReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserReservationRequest) Reset() {
	*x = UserReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReservationRequest) ProtoMessage() {}

func (x *UserReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReservationRequest.ProtoReflect.Descriptor instead.
func (*UserReservationRequest) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserReservationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*CreateBookingRequest `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *ReservationResponse) Reset() {
	*x = ReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationResponse) ProtoMessage() {}

func (x *ReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationResponse.ProtoReflect.Descriptor instead.
func (*ReservationResponse) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{9}
}

func (x *ReservationResponse) GetReservations() []*CreateBookingRequest {
	if x != nil {
		return x.Reservations
	}
	return nil
}

type CanceledBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CanceledBookingRequest) Reset() {
	*x = CanceledBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanceledBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanceledBookingRequest) ProtoMessage() {}

func (x *CanceledBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanceledBookingRequest.ProtoReflect.Descriptor instead.
func (*CanceledBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{10}
}

func (x *CanceledBookingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CanceledBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CanceledBookingResponse) Reset() {
	*x = CanceledBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanceledBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanceledBookingResponse) ProtoMessage() {}

func (x *CanceledBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanceledBookingResponse.ProtoReflect.Descriptor instead.
func (*CanceledBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_service_proto_rawDescGZIP(), []int{11}
}

func (x *CanceledBookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_booking_service_proto protoreflect.FileDescriptor

var file_booking_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x20, 0x47, 0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x21, 0x47, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x61, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xd8, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x22,
	0x31, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4c, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x3c, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x30,
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc5, 0x09, 0x0a, 0x0e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x73, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12,
	0x1f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d,
	0x12, 0x65, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f,
	0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x65, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x7e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x12, 0x9c,
	0x01, 0x0a, 0x19, 0x47, 0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x61, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x61, 0x73, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x61, 0x73, 0x74, 0x12, 0x83, 0x01,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x67, 0x65, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x7d, 0x12, 0x7e, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a,
	0x01, 0x2a, 0x1a, 0x1d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x16, 0x5a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_booking_service_proto_rawDescOnce sync.Once
	file_booking_service_proto_rawDescData = file_booking_service_proto_rawDesc
)

func file_booking_service_proto_rawDescGZIP() []byte {
	file_booking_service_proto_rawDescOnce.Do(func() {
		file_booking_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_service_proto_rawDescData)
	})
	return file_booking_service_proto_rawDescData
}

var file_booking_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_booking_service_proto_goTypes = []interface{}{
	(*GuestHasReservationInPastRequest)(nil),  // 0: booking.GuestHasReservationInPastRequest
	(*GuestHasReservationInPastResponse)(nil), // 1: booking.GuestHasReservationInPastResponse
	(*GetAllPendingRequest)(nil),              // 2: booking.GetAllPendingRequest
	(*CreateBookingRequest)(nil),              // 3: booking.CreateBookingRequest
	(*EmptyRequst)(nil),                       // 4: booking.EmptyRequst
	(*CreateBookingResponse)(nil),             // 5: booking.CreateBookingResponse
	(*BookingResponse)(nil),                   // 6: booking.BookingResponse
	(*ReservationRequest)(nil),                // 7: booking.ReservationRequest
	(*UserReservationRequest)(nil),            // 8: booking.UserReservationRequest
	(*ReservationResponse)(nil),               // 9: booking.ReservationResponse
	(*CanceledBookingRequest)(nil),            // 10: booking.CanceledBookingRequest
	(*CanceledBookingResponse)(nil),           // 11: booking.CanceledBookingResponse
}
var file_booking_service_proto_depIdxs = []int32{
	3,  // 0: booking.BookingResponse.bookings:type_name -> booking.CreateBookingRequest
	3,  // 1: booking.ReservationResponse.reservations:type_name -> booking.CreateBookingRequest
	3,  // 2: booking.BookingService.CreateBooking:input_type -> booking.CreateBookingRequest
	4,  // 3: booking.BookingService.GetAll:input_type -> booking.EmptyRequst
	2,  // 4: booking.BookingService.GetAllOnPending:input_type -> booking.GetAllPendingRequest
	3,  // 5: booking.BookingService.Decline:input_type -> booking.CreateBookingRequest
	3,  // 6: booking.BookingService.Confirm:input_type -> booking.CreateBookingRequest
	7,  // 7: booking.BookingService.GetAllReservations:input_type -> booking.ReservationRequest
	0,  // 8: booking.BookingService.GuestHasReservationInPast:input_type -> booking.GuestHasReservationInPastRequest
	8,  // 9: booking.BookingService.GetUserReservations:input_type -> booking.UserReservationRequest
	8,  // 10: booking.BookingService.GetFinishedReservations:input_type -> booking.UserReservationRequest
	10, // 11: booking.BookingService.CanceledBooking:input_type -> booking.CanceledBookingRequest
	5,  // 12: booking.BookingService.CreateBooking:output_type -> booking.CreateBookingResponse
	6,  // 13: booking.BookingService.GetAll:output_type -> booking.BookingResponse
	6,  // 14: booking.BookingService.GetAllOnPending:output_type -> booking.BookingResponse
	5,  // 15: booking.BookingService.Decline:output_type -> booking.CreateBookingResponse
	5,  // 16: booking.BookingService.Confirm:output_type -> booking.CreateBookingResponse
	9,  // 17: booking.BookingService.GetAllReservations:output_type -> booking.ReservationResponse
	1,  // 18: booking.BookingService.GuestHasReservationInPast:output_type -> booking.GuestHasReservationInPastResponse
	9,  // 19: booking.BookingService.GetUserReservations:output_type -> booking.ReservationResponse
	9,  // 20: booking.BookingService.GetFinishedReservations:output_type -> booking.ReservationResponse
	11, // 21: booking.BookingService.CanceledBooking:output_type -> booking.CanceledBookingResponse
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_booking_service_proto_init() }
func file_booking_service_proto_init() {
	if File_booking_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuestHasReservationInPastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuestHasReservationInPastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPendingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequst); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReservationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanceledBookingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_booking_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanceledBookingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_booking_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_service_proto_goTypes,
		DependencyIndexes: file_booking_service_proto_depIdxs,
		MessageInfos:      file_booking_service_proto_msgTypes,
	}.Build()
	File_booking_service_proto = out.File
	file_booking_service_proto_rawDesc = nil
	file_booking_service_proto_goTypes = nil
	file_booking_service_proto_depIdxs = nil
}
