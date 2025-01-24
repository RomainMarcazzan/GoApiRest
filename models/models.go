package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type Notif struct {
	ID          uuid.UUID  `json:"id"`
	Message     string     `json:"message"`
	NotifSongId string     `json:"notifSongId"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	IsView      bool       `json:"isView"`
	ReceiverId  uuid.UUID  `json:"receiverId"`
	Avatar      *string    `json:"avatar"`
}

type ProPosition struct {
	ID        uuid.UUID  `json:"id"`
	ProId     uuid.UUID  `json:"proId"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}
