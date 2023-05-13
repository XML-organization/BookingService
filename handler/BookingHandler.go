package handler

import (
	"context"
	"fmt"

	"log"

	"booking-service/model"
	"booking-service/service"

	accServicepb "github.com/XML-organization/common/proto/accomodation_service"
	booking "github.com/XML-organization/common/proto/booking_service"
	pb "github.com/XML-organization/common/proto/booking_service"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.Dial("accomodation-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	accomodationService := accServicepb.NewAccommodationServiceClient(conn)

	autoApprovalResponse, err := accomodationService.GetAutoApprovalForAccommodation(context.TODO(), &accServicepb.AutoApprovalRequest{AccommodationId: booking.AccomodationID.String()})
	if err != nil {
		println(err.Error())
		return nil, err
	}

	println(autoApprovalResponse.AutoApproval)

	if autoApprovalResponse.AutoApproval == true {
		booking.Status = model.CONFIRMED
		println(booking.Status)
		println("usao u TRUE")
	} else {
		booking.Status = model.PENDING
		println(booking.Status)
		println("usao u FALSE")
	}

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

func (bookingHandler *BookingHandler) GetAllOnPending(ctx context.Context, request *booking.GetAllPendingRequest) (*pb.BookingResponse, error) {
	hostID, err := uuid.Parse(request.HostId)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial("accomodation-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	accomodationService := accServicepb.NewAccommodationServiceClient(conn)

	accommodations, err := accomodationService.GetAllAccomodations(context.TODO(), &accServicepb.GetAllAccomodationsRequest{HostId: hostID.String()})
	if err != nil {
		println(err.Error())
		return nil, err
	}

	bookings, requestMessage := bookingHandler.BookingService.GetAllPendingBookings()

	if requestMessage.Message != "Success!" {
		return nil, fmt.Errorf("an error occurred: %s", requestMessage.Message)
	}

	println("Smjestaji od usera sa ovim id-om", hostID.String())
	for index, accommodation := range accommodations.Accomodations {
		fmt.Printf("Index: %d, Name: %s\n", index, accommodation.Name)
	}
	println("kraj")

	accommodationsMap := make(map[string]bool)
	for _, accommodation := range accommodations.Accomodations {
		accommodationsMap[accommodation.Id] = true
	}

	// Filtriranje bookings na osnovu accommodations
	filteredBookings := []model.Booking{}
	for _, booking := range bookings {
		if accommodationsMap[booking.AccomodationID.String()] {
			filteredBookings = append(filteredBookings, booking)
		}
	}

	response := pb.BookingResponse{
		Bookings: []*pb.CreateBookingRequest{},
	}

	for _, booking := range filteredBookings {
		protoBooking := mapBookingToCreateBookingRequest(&booking)
		response.Bookings = append(response.Bookings, protoBooking)
	}

	return &response, nil
}

func (bookingHandler *BookingHandler) Decline(ctx context.Context, in *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	message, err := bookingHandler.BookingService.Decline(mapBookingFromCreateBookingRequest(in))

	return &pb.CreateBookingResponse{
		Message: message.Message,
	}, err
}

func (bookingHandler *BookingHandler) Confirm(ctx context.Context, in *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	message, err := bookingHandler.BookingService.Confirm(mapBookingFromCreateBookingRequest(in))

	return &pb.CreateBookingResponse{
		Message: message.Message,
	}, err
}
