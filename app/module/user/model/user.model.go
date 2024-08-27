package user_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"type:char(36);index;not null"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Username string    `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string    `json:"password" gorm:"type:varchar(255);not null"`
	Status   int       `json:"status" gorm:"type:int;default:1"`
}

func (r *User) TableName() string {
	return "users"
}

type UserCreateStorage struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}
