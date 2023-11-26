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
	Role      GetAllCustomersResponseRole
	Profile   GetAllCustomersResponseProfile
	Screens   []GetAllCustomersResponseScreen
}
type GetAllCustomersResponseRole struct {
	RoleKey   string
	RoleCode  string
	RoleValue uint8
}
type GetAllCustomersResponseProfile struct {
	TokenCount         uint32
	BalanceIntegerPart uint32
	BalanceDecimalPart uint32
}
type GetAllCustomersResponseScreen struct {
	Key         string
	Value       string
	Description string
}
