package repository

import (
	"booking-service/model"

	"github.com/google/uuid"
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

func (repo *BookingRepository) GetAllPending() ([]model.Booking, model.RequestMessage) {
	var bookings []model.Booking
	dbResult := repo.DatabaseConnection.Find(&bookings, "status = ?", 1)

	if dbResult.Error != nil {
		return nil, model.RequestMessage{
			Message: "An error occurred, please try again!",
		}
	}

	return bookings, model.RequestMessage{
		Message: "Success!",
	}
}

func (repo *BookingRepository) FindById(id uuid.UUID) (model.Booking, error) {
	booking := model.Booking{}

	dbResult := repo.DatabaseConnection.First(&booking, "id = ?", id)

	if dbResult != nil {
		return booking, dbResult.Error
	}

	return booking, nil
}

func (repo *BookingRepository) UpdateBooking(booking model.Booking) error {

	sqlStatementUser := `
		UPDATE bookings
		SET status = $2
		WHERE id = $1;`

	dbResult := repo.DatabaseConnection.Exec(sqlStatementUser, booking.ID, booking.Status)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *BookingRepository) GetAllReservations(accomodationID uuid.UUID) ([]model.Booking, error) {
	reservations := []model.Booking{}
	result := repo.DatabaseConnection.Where("accomodation_id = ?", accomodationID).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func (repo *BookingRepository) CanceledReservation(reservationId uuid.UUID) model.RequestMessage {

	sqlStatementBooking := `
		UPDATE booking
		SET status = $2
		WHERE id = $1;`

	dbResult1 := repo.DatabaseConnection.Exec(sqlStatementBooking, reservationId, model.DECLINED)

	if dbResult1.Error != nil {
		return model.RequestMessage{
			Message: "An error occured, please try again!",
		}
	}
	return model.RequestMessage{
		Message: "Success!",
	}
}
