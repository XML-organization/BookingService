package model

import (
	"time"

	"github.com/google/uuid"
)

type AvailabilityParams struct {
	AccomodationID uuid.UUID `json:"accomodationID"`
	StartDate      time.Time `json:"startDate" gorm:"not null;type:string"`
	EndDate        time.Time `json:"endDate" gorm:"not null;type:string"`
}

type IfAvailable struct {
	Message bool `json:"message"`
}
