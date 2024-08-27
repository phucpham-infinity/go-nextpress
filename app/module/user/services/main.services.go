package user_services

import (
	"context"

	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
	user_storage "github.com/phucpham-infinity/go-nextpress/app/module/user/storage"
)

type UserServices interface {
	CreateUser(ctx context.Context, data *user_model.UserCreateStorage) error
}

type userServices struct {
	storage user_storage.UserStorage
}

func NewUserServices(storage user_storage.UserStorage) *userServices {
	return &userServices{storage: storage}
}
