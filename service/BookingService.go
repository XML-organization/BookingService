package service

import (
	"booking-service/model"
	"booking-service/repository"
)

type BookingService struct {
	BookingRepo *repository.BookingRepository
}

func NewBookingService(repo *repository.BookingRepository) *BookingService {
	return &BookingService{
		BookingRepo: repo,
	}
}

func (service *BookingService) CreateBooking(booking model.Booking) (model.RequestMessage, error) {

	response := model.RequestMessage{
		Message: service.BookingRepo.Create(booking).Message,
	}

	return response, nil
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

