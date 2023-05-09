package repository

import (
	"booking-service/model"

	"gorm.io/gorm"
)

type BookingRepository struct {
	DatabaseConnection *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	err := db.AutoMigrate(&model.Booking{})
	if err != nil {
		return nil
	}

	return &BookingRepository{
		DatabaseConnection: db,
	}
}

func (repo *BookingRepository) Create(booking model.Booking) model.RequestMessage {
	dbResult := repo.DatabaseConnection.Save(booking)

	if dbResult.Error != nil {
		return model.RequestMessage{
			Message: "An error occured, please try again!",
		}
	}

	return model.RequestMessage{
		Message: "Success!",
	}
}

func (repo *BookingRepository) GetAll() ([]model.Booking, model.RequestMessage) {
	var bookings []model.Booking
	dbResult := repo.DatabaseConnection.Find(&bookings)

	if dbResult.Error != nil {
		return nil, model.RequestMessage{
			Message: "An error occurred, please try again!",
		}
	}

	return bookings, model.RequestMessage{
		Message: "Success!",
	}
}
