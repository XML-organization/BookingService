package service

import (
	"booking-service/model"
	"booking-service/repository"
	"time"

	"github.com/google/uuid"
)

type BookingService struct {
	BookingRepo      *repository.BookingRepository
	BookingNeo4jRepo *repository.BookingNeo4jRepository
}

func NewBookingService(repo *repository.BookingRepository, neo4jRepo *repository.BookingNeo4jRepository) *BookingService {
	return &BookingService{
		BookingRepo:      repo,
		BookingNeo4jRepo: neo4jRepo,
	}
}

func (service *BookingService) GuestHasReservationInPast(ids []string, guestId string) string {

	return service.BookingRepo.GuestHasReservationInPast(ids, guestId)
}

func (service *BookingService) CreateBooking(booking model.Booking) (model.RequestMessage, error) {
	service.BookingNeo4jRepo.SaveBooking(booking)

	response := model.RequestMessage{
		Message: service.BookingRepo.Create(booking).Message,
	}

	return response, nil
}

func (service *BookingService) Decline(booking model.Booking) (model.RequestMessage, error) {
	booking, err1 := service.BookingRepo.FindById(booking.ID)
	if err1 != nil {
		return model.RequestMessage{Message: "An error occurred, please try again!"}, err1
	}

	booking.Status = model.DECLINED
	err := service.BookingRepo.UpdateBooking(booking)
	if err != nil {
		return model.RequestMessage{Message: "An error occurred, please try again!"}, err
	}

	return model.RequestMessage{Message: "Success!"}, nil
}

func (service *BookingService) Confirm(booking model.Booking) (model.RequestMessage, error) {
	booking, err1 := service.BookingRepo.FindById(booking.ID)
	if err1 != nil {
		return model.RequestMessage{Message: "An error occurred, please try again!"}, err1
	}

	booking.Status = model.CONFIRMED
	err := service.BookingRepo.UpdateBooking(booking)
	if err != nil {
		return model.RequestMessage{Message: "An error occurred, please try again!"}, err
	}

	return model.RequestMessage{Message: "Success!"}, nil
}

func (service *BookingService) GetAllBookings() ([]model.Booking, model.RequestMessage) {
	bookings, err := service.BookingRepo.GetAll()
	if err.Message != "Success!" {
		return nil, model.RequestMessage{
			Message: "An error occurred, please try again!",
		}
	}
	return bookings, err
}

func (service *BookingService) GetAllPendingBookings() ([]model.Booking, model.RequestMessage) {
	bookings, err := service.BookingRepo.GetAllPending()
	if err.Message != "Success!" {
		return nil, model.RequestMessage{
			Message: "An error occurred, please try again!",
		}
	}
	return bookings, err
}

func (service *BookingService) IfAvailable(params model.AvailabilityParams) (model.IfAvailable, model.RequestMessage) {
	bookings, err := service.BookingRepo.GetAll()
	if err.Message != "Success!" {
		return model.IfAvailable{Message: false}, model.RequestMessage{Message: "An error occurred, please try again!"}
	}

	for _, booking := range bookings {
		if params.StartDate.Before(booking.EndDate) && params.EndDate.After(booking.StartDate) {
			return model.IfAvailable{Message: false}, model.RequestMessage{Message: "An error occurred, please try again!"}
		}
	}

	return model.IfAvailable{Message: true}, model.RequestMessage{Message: "Available!"}
}

func (service *BookingService) GetAllReservations(accomodationID uuid.UUID) ([]model.Booking, error) {
	reservations, err := service.BookingRepo.GetAllReservations(accomodationID)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (service *BookingService) CanceledReservation(reservationId uuid.UUID) (model.RequestMessage, error) {
	reservation, err := service.BookingRepo.FindById(reservationId)
	println("Errorin FindById in booking: ", err)

	if time.Now().AddDate(0, 0, -1).Before(reservation.StartDate) {
		response := model.RequestMessage{
			Message: service.BookingRepo.CanceledReservation(reservationId).Message,
		}
		return response, nil
	}
	response := model.RequestMessage{
		Message: "The time in which you can cancel the reservation has passed!",
	}
	return response, nil
}

func (service *BookingService) GetUserReservations(userID uuid.UUID) ([]model.Booking, error) {
	reservations, err := service.BookingRepo.GetUserReservations(userID)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}
