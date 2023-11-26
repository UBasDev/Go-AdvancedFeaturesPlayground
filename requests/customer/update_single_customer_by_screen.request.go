package requests

import "github.com/google/uuid"

type UpdateSingleCustomerByScreen struct {
	CustomerId uuid.UUID
	ScreenIds  []uuid.UUID
}
