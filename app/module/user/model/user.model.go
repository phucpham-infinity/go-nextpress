package user_model

import (
	"github.com/google/uuid"
)

type User struct {
	UUID          uuid.UUID `json:"uuid"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Status        int       `json:"status"`
	ActivationKey string    `json:"activation_key"`
	RegisteredAt  string    `json:"registered_at"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
	DeletedAt     string    `json:"deleted_at"`
}

func (r *User) TableName() string {
	return "users"
}

type UserRegisterStorage struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
