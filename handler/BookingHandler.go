package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"booking-service/model"
	"booking-service/service"

	accServicepb "github.com/XML-organization/common/proto/accomodation_service"
	booking "github.com/XML-organization/common/proto/booking_service"
	pb "github.com/XML-organization/common/proto/booking_service"
	userServicepb "github.com/XML-organization/common/proto/user_service"
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

func (handler *BookingHandler) GuestHasReservationInPast(ctx context.Context, in *pb.GuestHasReservationInPastRequest) (*pb.GuestHasReservationInPastResponse, error) {
	println("aaaa")
	return &pb.GuestHasReservationInPastResponse{
		Message: handler.BookingService.GuestHasReservationInPast(in.AccomodationsId, in.GuestId),
	}, nil
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
		log.Println(err)
		return nil, err
	}
	accommodation, err := accomodationService.GetOneAccomodation(context.TODO(), &accServicepb.GetOneAccomodationRequest{AccomodationId: booking.AccomodationID.String()})

	println(autoApprovalResponse.AutoApproval)

	if autoApprovalResponse.AutoApproval == true {
		booking.Status = model.CONFIRMED
		println(booking.Status)
		println("usao u TRUE")
	} else {
		booking.Status = model.PENDING
		println(booking.Status)
		println("usao u FALSE")
		//slanje notifikacije

		conn, err := grpc.Dial("user_service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		userService := userServicepb.NewUserServiceClient(conn)

		saveResponse, err := userService.SaveNotification(context.TODO(), &userServicepb.SaveRequest{Id: uuid.NewString(), NotificationTime: time.Now().Format("2006-01-02 15:04:05"), Text: "Imate novi zahtjev za rezervaciju u " + accommodation.Accomodation.Name + " !", UserID: accommodation.Accomodation.IdHost, Status: "0", Category: "RequestCreated"})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		println(saveResponse.Message)

	}

	message, err := bookingHandler.BookingService.CreateBooking(booking)
	if err != nil {
		log.Println(err)
	}
	response := pb.CreateBookingResponse{
		Message: message.Message,
	}

	return &response, err
}

func (bookingHandler *BookingHandler) GetAll(ctx context.Context, empty *booking.EmptyRequst) (*pb.BookingResponse, error) {
	bookings, requestMessage := bookingHandler.BookingService.GetAllBookings()

	if requestMessage.Message != "Success!" {
		log.Println(requestMessage)
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
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}

	bookings, requestMessage := bookingHandler.BookingService.GetAllPendingBookings()

	if requestMessage.Message != "Success!" {
		log.Println(requestMessage)
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
	booking := mapBookingFromCreateBookingRequest(in)
	message, err := bookingHandler.BookingService.Decline(booking)
	//slanje notifikacije

	conn, err := grpc.Dial("user_service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userService := userServicepb.NewUserServiceClient(conn)

	saveResponse, err := userService.SaveNotification(context.TODO(), &userServicepb.SaveRequest{Id: uuid.NewString(), NotificationTime: time.Now().Format("2006-01-02 15:04:05"), Text: "Vas zahtjev za rezervaciju je odbijen !", UserID: booking.UserID.String(), Status: "0", Category: "ReservationReply"})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println(saveResponse.Message)

	return &pb.CreateBookingResponse{
		Message: message.Message,
	}, err
}

func (bookingHandler *BookingHandler) Confirm(ctx context.Context, in *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	booking := mapBookingFromCreateBookingRequest(in)
	message, err := bookingHandler.BookingService.Confirm(booking)
	if err != nil {
		log.Println(err)
	}

	//slanje notifikacije

	conn, err := grpc.Dial("user_service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userService := userServicepb.NewUserServiceClient(conn)

	saveResponse, err := userService.SaveNotification(context.TODO(), &userServicepb.SaveRequest{Id: uuid.NewString(), NotificationTime: time.Now().Format("2006-01-02 15:04:05"), Text: "Vas zahtjev za rezervaciju je odobren !", UserID: booking.UserID.String(), Status: "0", Category: "ReservationReply"})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println(saveResponse.Message)

	return &pb.CreateBookingResponse{
		Message: message.Message,
	}, err
}

func (handler *BookingHandler) GetAllReservations(ctx context.Context, request *pb.ReservationRequest) (*pb.ReservationResponse, error) {
	accomodationID, err := uuid.Parse(request.AccomodationId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	reservations, err := handler.BookingService.GetAllReservations(accomodationID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := &pb.ReservationResponse{
		Reservations: []*pb.CreateBookingRequest{},
	}
	for _, availability := range reservations {
		current := mapBookingToCreateBookingRequest(&availability)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (bookingHandler *BookingHandler) CanceledBooking(ctx context.Context, in *pb.CanceledBookingRequest) (*pb.CanceledBookingResponse, error) {

	bookingId := mapBookingFromCanceledBookingRequest(in)

	booking, _ := bookingHandler.BookingService.BookingRepo.FindById(bookingId)

	message, err := bookingHandler.BookingService.CanceledReservation(bookingId)
	if err != nil {
		log.Println(err)
	}

	//slanje notifikacije

	conn, err := grpc.Dial("accomodation-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	accomodationService := accServicepb.NewAccommodationServiceClient(conn)
	accommodation, err := accomodationService.GetOneAccomodation(context.TODO(), &accServicepb.GetOneAccomodationRequest{AccomodationId: booking.AccomodationID.String()})

	conn1, err1 := grpc.Dial("user_service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userService := userServicepb.NewUserServiceClient(conn1)

	saveResponse, err := userService.SaveNotification(context.TODO(), &userServicepb.SaveRequest{Id: uuid.NewString(), NotificationTime: time.Now().Format("2006-01-02 15:04:05"), Text: "Korisnik je otkazao smjestaj u  " + accommodation.Accomodation.Name + " od " + booking.StartDate.Format("2006-01-02 15:04:05") + " do " + booking.EndDate.Format("2006-01-02 15:04:05") + "!", UserID: accommodation.Accomodation.IdHost, Status: "0", Category: "ReservationCanceled"})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println(saveResponse.Message)

	response := pb.CanceledBookingResponse{
		Message: message.Message,
	}

	return &response, err
}

func (handler *BookingHandler) GetUserReservations(ctx context.Context, request *pb.UserReservationRequest) (*pb.ReservationResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	reservations, err := handler.BookingService.GetUserReservations(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := &pb.ReservationResponse{
		Reservations: []*pb.CreateBookingRequest{},
	}
	for _, availability := range reservations {
		current := mapBookingToCreateBookingRequest(&availability)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *BookingHandler) IsExceptional(ctx context.Context, in *pb.IsExceptionalRequest) (*pb.IsExceptionalResponse, error) {
	println("Usao sam u Booking Service---------")
	println("HostID u Booking Servicu: ", in.UserId)

	conn, err := grpc.Dial("accomodation-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	accomodationService := accServicepb.NewAccommodationServiceClient(conn)

	accommodations, err := accomodationService.GetAllAccomodations(context.TODO(), &accServicepb.GetAllAccomodationsRequest{HostId: in.UserId})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bookings, _ := handler.BookingService.GetAllBookings()

	var cancelCnt float64 = 0
	var confirmCnt float64 = 0
	var days = 0

	println("Bookings:")
	for j, tmp := range bookings {
		println(j, ". Accomodation ID", tmp.AccomodationID.String())
	}

	println("Accomodations:")
	for j, tmp := range accommodations.Accomodations {
		println(j, ". Accomodation ID", tmp.IdHost)
	}

	for _, book := range bookings {
		println("Usao u prvi for.")
		for i, accommodation := range accommodations.Accomodations {
			println("------------------- ", i)
			println("accommodation.Id:", accommodation.Id, ".")
			println("book.ID:", book.AccomodationID.String(), ".")
			_, err := uuid.Parse(accommodation.Id)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			if accommodation.Id == book.AccomodationID.String() {
				println("PROSAO")
				println("book.status: ", book.Status)
				println("model.CONFIRMED: ", model.CONFIRMED)
				if book.Status == model.CONFIRMED {
					println("Usao u CONFIRMED")
					confirmCnt = confirmCnt + 1
					days = days + (int(book.EndDate.Sub(book.StartDate).Hours() / 24))
					println("Trenutno potvrdjenih rezervacija:", confirmCnt)
					println("Trenutno ukupno dana:", days)
				} else if book.Status == model.CANCELED {
					println("Usao u CANCELED")
					cancelCnt = cancelCnt + 1
					println("Trenutno odbijenih rezervacija:", cancelCnt)
				}
			}
		}
	}

	var cancelPer float64 = cancelCnt / (cancelCnt + confirmCnt)

	println("cancelCnt: ", cancelCnt)
	println("confirmCnt: ", confirmCnt)
	println("days: ", days)

	var threshold = 0.05
	if confirmCnt >= 5 && days > 50 && cancelPer < threshold {
		response := pb.IsExceptionalResponse{
			IsExceptional: true,
		}
		return &response, err
	} else {
		response := pb.IsExceptionalResponse{
			IsExceptional: false,
		}
		return &response, err
	}

}

func (handler *BookingHandler) GetFinishedReservations(ctx context.Context, request *pb.UserReservationRequest) (*pb.ReservationResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	reservations, err := handler.BookingService.GetFinishedReservations(userID)
	if err != nil {
		return nil, err
	}

	response := &pb.ReservationResponse{
		Reservations: []*pb.CreateBookingRequest{},
	}
	for _, availability := range reservations {
		current := mapBookingToCreateBookingRequest(&availability)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
