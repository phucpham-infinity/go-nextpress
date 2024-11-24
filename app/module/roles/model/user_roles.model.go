package roles_model

import (
	"encoding/json"
)

type CreateRolesBody struct {
	Role       string          `json:"role" validate:"required"`
	Permission json.RawMessage `json:"permission" validate:"required"`
}
