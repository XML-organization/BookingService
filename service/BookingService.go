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
