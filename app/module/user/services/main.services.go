package user_services

import (
	"context"

	user_database "github.com/phucpham-infinity/go-nextpress/app/module/user/database"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

type UserServices interface {
	RegisterUser(ctx context.Context, data *user_model.UserRegisterStorage) (user_database.CreateUserRow, error)
}

type userServices struct {
	storage *user_database.Queries
}

func NewUserServices(storage *user_database.Queries) *userServices {
	return &userServices{storage: storage}
}
