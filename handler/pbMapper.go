package handler

import (
	"booking-service/model"
	"time"

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

	userID, err := uuid.Parse(booking.UserID)
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(booking.GuestNumber)
	if err != nil {
		panic(err)
	}

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, booking.StartDate)
	if err != nil {
		panic(err)
	}

	endDate, err := time.Parse(layout, booking.EndDate)
	if err != nil {
		panic(err)
	}

	statusInt, err := strconv.Atoi(booking.Status)

	return model.Booking{
		ID:             bookingID,
		AccomodationID: accomodationID,
		UserID:         userID,
		StartDate:      startDate,
		EndDate:        endDate,
		GuestNumber:    num,
		Status:         model.Status(statusInt),
	}
}

func mapBookingToCreateBookingRequest(booking *model.Booking) *pb.CreateBookingRequest {
	bookingID := booking.ID.String()
	accomodationID := booking.AccomodationID.String()
	userID := booking.UserID.String()
	guestNumber := strconv.Itoa(booking.GuestNumber)
	startDate := booking.StartDate.Format("2006-01-02")
	endDate := booking.EndDate.Format("2006-01-02")

	var statusString string
	switch booking.Status {
	case model.CONFIRMED:
		statusString = "CONFIRMED"
	case model.PENDING:
		statusString = "PENDING"
	case model.DECLINED:
		statusString = "DECLINED"
	}

	return &pb.CreateBookingRequest{
		Id:             bookingID,
		AccomodationID: accomodationID,
		UserID:         userID,
		StartDate:      startDate,
		EndDate:        endDate,
		GuestNumber:    guestNumber,
		Status:         statusString,
	}
}
