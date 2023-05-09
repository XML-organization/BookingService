package handler

import (
	"context"

	"booking-service/service"

	pb "github.com/XML-organization/common/proto/booking_service"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	BookingService *service.BookingService
}

func NewBookingHandler(service *service.BookingService) *BookingHandler {
	return &BookingHandler{
		BookingService: service,
	}
}

func (bookingHandler *BookingHandler) CreateBooking(ctx context.Context, in *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {

	booking := mapBookingFromCreateBookingRequest(in)

	message, err := bookingHandler.BookingService.CreateBooking(booking)

	response := pb.CreateBookingResponse{
		Message: message.Message,
	}

	return &response, err
}
