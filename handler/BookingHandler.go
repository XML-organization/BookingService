package handler

import (
	"context"
	"fmt"

	greet "booking-service/proto/greet"
	"booking-service/service"

	booking "github.com/XML-organization/common/proto/booking_service"
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

func (bookingHandler *BookingHandler) GetAll(ctx context.Context, empty *booking.EmptyRequst) (*pb.BookingResponse, error) {
	bookings, requestMessage := bookingHandler.BookingService.GetAllBookings()

	if requestMessage.Message != "Success!" {
		return nil, fmt.Errorf("an error occurred: %s", requestMessage.Message)
	}

	response := pb.BookingResponse{
		Bookings: []*pb.CreateBookingRequest{},
	}

	for _, booking := range bookings {
		protoBooking := mapBookingToCreateBookingRequest(&booking)
		response.Bookings = append(response.Bookings, protoBooking)
	}

	return &response, nil
}

//Test!

func (h BookingHandler) Greet(ctx context.Context, request *greet.Request) (*greet.Response, error) {
	return &greet.Response{
		Greeting: fmt.Sprintf("Hi %s!", request.Name),
	}, nil
}
