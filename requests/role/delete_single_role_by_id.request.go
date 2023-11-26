package requests

import "github.com/google/uuid"

type DeleteSingleRoleByIdRequest struct {
	Id uuid.UUID
}
