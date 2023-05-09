package handler

import (
	"booking-service/model"

	pb "github.com/XML-organization/common/proto/booking_service"

	"github.com/google/uuid"

	"strconv"
)

func mapBookingFromCreateBookingRequest(booking *pb.CreateBookingRequest) model.Booking {

	bookingID, err := uuid.Parse(booking.Id)
	if err != nil {
		panic(err)
	}

	accomodationID, err := uuid.Parse(booking.AccomodationID)
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(booking.GuestNumber)
	if err != nil {
		panic(err)
	}

	return model.Booking{
		ID:             bookingID,
		AccomodationID: accomodationID,
		StartDate:      booking.StartDate,
		EndDate:        booking.EndDate,
		GuestNumber:    num,
		Status:         model.Status(booking.Status.Number()),
	}
}
