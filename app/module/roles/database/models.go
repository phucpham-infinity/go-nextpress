// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package roles_database

import (
	"encoding/json"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

type Role struct {
	ID         uuid.UUID       `json:"id"`
	Role       string          `json:"role"`
	Permission json.RawMessage `json:"permission"`
	CreatedAt  null.Time       `json:"created_at"`
	UpdatedAt  null.Time       `json:"updated_at"`
}
