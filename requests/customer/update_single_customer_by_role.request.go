package requests

import "github.com/google/uuid"

type UpdateSingleCustomerByRole struct {
	CustomerId uuid.UUID
	RoleId     uuid.UUID
}
