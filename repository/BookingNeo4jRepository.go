package repository

import (
	"booking-service/model"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type BookingNeo4jRepository struct {
	Session neo4j.Session
}

func NewBookingNeo4jRepository(driver neo4j.Driver) *BookingNeo4jRepository {
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil
	}

	return &BookingNeo4jRepository{
		Session: session,
	}
}

func (repo *BookingNeo4jRepository) Close() {
	repo.Session.Close()
}

func (repo *BookingNeo4jRepository) SaveBooking(booking model.Booking) error {
	session := repo.Session
	println("User id: " + booking.UserID.String())
	println("Accomodation id: " + booking.AccomodationID.String())
	// Izvrši upit za kreiranje veze između korisnika i smještaja
	res, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			`
			MATCH (u:User {idInPostgre: $userId})
			MATCH (a:Accommodation {idInPostgre: $accommodationId})
			CREATE (u)-[:Reservation {
				id: $bookingId,
				startDate: $startDate,
				endDate: $endDate,
				status: $status,
				guestNumber: $guestNumber
			}]->(a)
			`,
			map[string]interface{}{
				"userId":          booking.UserID.String(),
				"accommodationId": booking.AccomodationID.String(),
				"bookingId":       booking.ID.String(),
				"startDate":       booking.StartDate,
				"endDate":         booking.EndDate,
				"status":          booking.Status,
				"guestNumber":     booking.GuestNumber,
			},
		)
		if err != nil {
			return nil, err
		}
		println("ALOOO USPIO SAM")
		return result, nil
	})

	if err != nil {
		println("USPIO SI AL GOVNO")
		return err
	}

	println(res)

	return nil
}
