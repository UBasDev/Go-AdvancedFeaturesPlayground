package requests

import "github.com/google/uuid"

type CreateSingleCustomerRequest struct {
	Firstname          string
	Lastname           string
	Email              string
	Age                uint8
	Gender             string
	RoleId             uuid.UUID
	ScreenIds          []uuid.UUID
	TokenCount         uint32
	BalanceIntegerPart uint32
	BalanceDecimalPart uint32
}
