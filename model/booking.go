package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID             uuid.UUID `json:"id"`
	AccomodationID uuid.UUID `json:"accomodationID"`
	StartDate      time.Time `json:"startDate" gorm:"not null;type:string"`
	EndDate        time.Time `json:"endDate" gorm:"not null;type:string"`
	GuestNumber    int       `json:"guestNumber" gorm:"not null"`
	Status         Status    `json:"status"`
}

/*func (booking *Booking) BeforeCreate(scope *gorm.DB) error {
	booking.ID = uuid.New()
	return nil
}*/

type RequestMessage struct {
	Message string `json:"message"`
}

type Empty struct {
}
