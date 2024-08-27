package user_storage

import (
	"context"

	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
	"gorm.io/gorm"
)

type UserStorage interface {
	CreateUser(ctx context.Context, data *user_model.UserCreateStorage) error
}

type userStore struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *userStore {
	return &userStore{db: db}
}
