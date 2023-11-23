package responses

import (
	"time"

	"github.com/google/uuid"
)

type GetAllCustomersResponse struct {
	Id        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint8
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
