package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID             uuid.UUID `json:"id"`
	AccomodationID uuid.UUID `json:"accomodationID"`
	StartDate      string    `json:"startDate" gorm:"not null;type:string"`
	EndDate        string    `json:"endDate" gorm:"not null;type:string"`
	GuestNumber    int       `json:"guestNumber" gorm:"not null"`
	Status         Status    `json:"status"`
}

func (booking *Booking) BeforeCreate(scope *gorm.DB) error {
	booking.ID = uuid.New()
	return nil
}

type RequestMessage struct {
	Message string `json:"message"`
}
